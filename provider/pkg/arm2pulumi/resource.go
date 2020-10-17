package arm2pulumi

import (
	"fmt"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/pcl"
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"strings"
)

func newResource(name string, rawBody map[string]interface{}, resourceToken string) (*resource, error) {
	return &resource{
		resourceName:   name,
		rawBody:        rawBody,
		resourceToken:  resourceToken,
		resourceParams: map[string]interface{}{},
	}, nil
}

type resource struct {
	resourceName   string
	rawBody        map[string]interface{}
	nameParam      string
	resourceToken  string // initial guess at pulumi resource token
	resourceParams map[string]interface{}
	dependsOn      model.BodyItem
	exclude        bool
}

func (r *resource) Name() string {
	return r.resourceName
}

func (r *resource) Args() interface{} {
	return r.rawBody
}

func (r *resource) SetArgs(body interface{}) {
	r.rawBody = body.(map[string]interface{})
}

func (r *resource) IntrospectElement(ctx *pclRenderContext, implicits implicitVariables) error {
	diag := Diagnostic{
		SourceElement: r.resourceName,
		Severity:      SevHigh,
	}
	templateName, ok := r.rawBody["name"]
	if !ok {
		r.exclude = true
		diag.Description = "missing required 'name' attribute"
		return &diag
	}

	var apiVersion, typ string
	var location interface{}
	additionalParams := map[string]interface{}{}
	for k, v := range r.rawBody {
		switch k {
		case "apiVersion":
			var ok bool
			apiVersion, ok = v.(string)
			if !ok {
				diag.Description = "expect apiVersion to be specified as a literal string"
				return &diag
			}
		case "type":
			var ok bool
			typ, ok = v.(string)
			if !ok {
				r.exclude = true
				diag.Description = "expect type to be specified as a literal string"
				return &diag
			}
		case "properties":
			r.resourceParams = map[string]interface{}{"properties": r.rawBody[k].(map[string]interface{})}
		case "dependsOn":
			asSlice, ok := v.([]interface{})
			if !ok {
				r.exclude = true
				diag.Description = "expect dependsOn to be a list"
				return &diag
			}

			var refs []model.Expression
			for _, dependency := range asSlice {
				switch dependency.(type) {
				case string:
					resourceId := dependency.(string)
					if variable, ok := ctx.resourceIdMap[resourceId]; ok {
						refs = append(refs, variable)
					}
				default:
					if expr, ok := dependency.(model.Expression); ok {
						refs = append(refs, expr)
					}
				}
			}

			if len(refs) > 0 {
				r.dependsOn = &model.Block{
					Type: "options",
					Body: &model.Body{
						Items: []model.BodyItem{
							&model.Attribute{
								Name: "dependsOn",
								Value: &model.TupleConsExpression{
									Expressions: refs,
								},
							},
						},
					},
				}
			}
		case "location":
			location = v
		case "tags", "sku", "plan", "kind":
			// These are additional/optional parameters which should be injected into body parameters of
			// the API requests if not specified explicitly.
			// See https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/template-syntax#resources
			// These could be a string or a variable reference
			additionalParams[k] = v
		}
	}

	if typ == "" || apiVersion == "" {
		r.exclude = true
		diag.Description = "expect type and apiVersion fields to be specified"
		return &diag
	}

	// Can't trust the casing in the template.
	// Need to do a soft-match against resources supported
	token := ctx.resourceTokenConverter.convert(typ, apiVersion)
	if token == "" {
		r.exclude = true
		diag.Description = fmt.Sprintf("no metadata found for resource type '%s' and version '%s'", typ, apiVersion)
		return &diag
	}
	diag.SourceToken = token
	r.resourceToken = token
	diag.SourceToken = r.resourceToken

	if gen.ShouldExclude(r.resourceToken) {
		r.exclude = true
		diag.Description = "resource doesn't support automatic conversion at this time"
		return &diag
	}

	// TODO: simplify lookupResource
	_, res, _ := lookupResource(ctx, token)
	// Set the name
	nameParam, err := extractResourceNameParameter(res)
	if err != nil {
		return err
	}
	r.nameParam = nameParam.Name
	r.resourceParams, err = r.transformRequestBody(ctx, res, r.resourceParams, templateName)
	if err != nil {
		return fmt.Errorf("missing resource name attribute for resource: %s", r.resourceName)
	}

	supported := codegen.NewStringSet()
	required := codegen.NewStringSet()
	for _, param := range res.PutParameters {
		// We fold body params in so we don't look for the body field directly
		if param.IsRequired && param.Body == nil {
			required.Add(param.Name)
			supported.Add(param.Name)
		}
		// Add the required properties in the body as well
		if param.Body != nil {
			for _, req := range param.Body.RequiredProperties {
				required.Add(req)
			}
			for prop := range param.Body.Properties {
				supported.Add(prop)
			}
		}
	}

	rgParam := extractResourceGroupNameParameter(res)
	if rgParam != nil {
		if required.Has(rgParam.Name) {
			if _, ok := r.resourceParams[rgParam.Name]; !ok {
				variable, dep, err := implicits.defaultResourceGroupName()
				if err != nil {
					return err
				}
				r.resourceParams[rgParam.Name] = model.VariableReference(variable)
				ctx.dep.RefersTo(dep)
			}
		}
	}

	// If location not specified, try to use the default resource group's location instead
	if _, ok := r.resourceParams["location"]; !ok {
		if supported.Has("location") {
			if location != nil {
				r.resourceParams["location"] = location
			} else {
				variable, dep, err := implicits.defaultResourceGroup()
				if err != nil {
					return err
				}
				r.resourceParams["location"] = pcl.RelativeTraversal(model.VariableReference(variable), "location")
				ctx.dep.RefersTo(dep)
			}
		}
	}

	for paramName, paramVal := range additionalParams {
		if _, ok := r.resourceParams[paramName]; !ok {
			if supported.Has(paramName) {
				r.resourceParams[paramName] = paramVal
			}
		}
	}

	// TODO: Need to find the missing required link back to the parent resource
	return nil
}

func (r *resource) LinkElements(ctx *pclRenderContext) error {
	return nil
}

func (r *resource) PCLExpression(ctx *pclRenderContext) ([]model.BodyItem, error) {
	if r.exclude {
		needsManualConversion := &model.Variable{
			Name:         "null",
			VariableType: model.NoneType,
		}

		return []model.BodyItem{&model.Attribute{
			Name:  r.Name(),
			Value: model.VariableReference(needsManualConversion),
		}}, nil
	}

	var items []model.BodyItem
	keys := codegen.SortedKeys(r.resourceParams)
	for _, k := range keys {
		v := r.resourceParams[k]
		val, err := pcl.RenderValue(v)
		if err != nil {
			return nil, err
		}
		items = append(items, &model.Attribute{
			Name:  k,
			Value: val,
		})
	}

	if r.dependsOn != nil {
		items = append(items, r.dependsOn)
	}
	comment := extractDescription(r.rawBody)
	block := model.Block{
		Tokens: getLeadingTriviaTokens(comment),
		Type:   "resource",
		Body:   &model.Body{Items: items},
		Labels: []string{r.resourceName, r.resourceToken},
	}
	return []model.BodyItem{&block}, nil
}

func (r *resource) transformRequestBody(ctx *pclRenderContext,
	res *provider.AzureAPIResource,
	resourceParams map[string]interface{},
	templateResourceName interface{},
) (map[string]interface{}, error) {
	metadataResParams := map[string]provider.AzureAPIParameter{}
	for _, param := range res.PutParameters {
		metadataResParams[param.Name] = param
	}

	// TODO: Stop guessing. Just always use the body field
	_, foundProperties := metadataResParams["properties"]

	// block specified as properties seems to be converted to body payload
	// so determine the body compliant field against metadata if a direct match
	// against "properties" is not found and wrap it
	if !foundProperties {
		var bodyField string
		for k, v := range metadataResParams {
			if v.Body != nil {
				bodyField = k
				break
			}
		}
		resourceParams = map[string]interface{}{
			bodyField: resourceParams,
		}
	}

	for k := range metadataResParams {
		_, found := resourceParams[k]
		if v, ok := r.rawBody[k]; !found && ok {
			resourceParams[k] = v
		}
	}

	resourceParams[r.nameParam] = templateResourceName

	flattened, err := gen.FlattenParams(resourceParams, metadataResParams, ctx.metadata.Types)
	if err != nil {
		return nil, fmt.Errorf("failed to transform parameters for resource: %s: %w", r.resourceName, err)
	}
	return flattened, nil
}

// lookupResources looks up the specified resource token and
// returns a corresponding resource if found.
// We don't have every resource in metadata so this does a slower
// lookup for aliases in the package spec if the metadata misses.
// Also, since the templates seem to be case insensitive, we allow
// a slow case insensitive lookup.
func lookupResource(ctx *pclRenderContext, resourceToken string) (string, *provider.AzureAPIResource, bool) {
	var actualResourceToken string

	res, ok := ctx.metadata.Resources[resourceToken]
	if !ok {
		resourceToken = strings.ToLower(resourceToken)
		// first search for case insensitive hit on metadata
		for k := range ctx.metadata.Resources {
			if strings.ToLower(k) == resourceToken {
				actualResourceToken = k
				break
			}
		}
		// next search for aliases in pkgSpec
		for name, r := range ctx.pkgSpec.Resources {
			if actualResourceToken != "" {
				break // detect inner loop break
			}
			if strings.ToLower(actualResourceToken) == resourceToken {
				actualResourceToken = name
				break
			}
			for _, a := range r.Aliases {
				if a.Type != nil && strings.ToLower(*a.Type) == resourceToken {
					actualResourceToken = name
					break
				}
			}
		}
		if actualResourceToken != "" {
			res, ok = ctx.metadata.Resources[actualResourceToken]
		}
	} else {
		actualResourceToken = resourceToken
	}
	return actualResourceToken, &res, ok
}
