package arm2pulumi

import (
	"fmt"
	"github.com/gedex/inflector"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/gen"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/pcl"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"strings"
)

func newResource(name string, rawBody map[string]interface{}, resourceToken string) *resource {
	return &resource{
		resourceName:  name,
		rawBody:       rawBody,
		resourceToken: resourceToken,
		resourceParams: map[string]interface{}{},
	}
}

type resource struct {
	resourceName  string
	rawBody       map[string]interface{}
	resourceToken string // initial guess at pulumi resource token
	resourceParams map[string]interface{}
	exclude       bool
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

func (r *resource) FinishInitializing(ctx *pclRenderContext, implicits implicitVariables) error {
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

	var location interface{}
	var apiVersion, typ string
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
		case "location":
			// location could be a string or a variable reference. Grab it so we can inject
			// into properties of the resource if missing
			location = v
		case "properties":
			r.resourceParams = map[string]interface{} { "properties": r.rawBody[k].(map[string]interface{})}
		case "dependsOn":
			// TODO
			continue
			// var arr []string
			// switch val := reflect.ValueOf(v); val.Kind() {
			// case reflect.Slice, reflect.Array:
			// 	for i := 0; i < val.Len(); i++ {
			// 		arr = append(arr, val.Index(i).String())
			// 	}
			// default:
			// 	return nil, fmt.Errorf("the \"dependsOn\" attribute for resource '%v' must be a string or list of strings", r.resourceName)
			// }

			// var refs []model.Expression
			// for _, v := range arr {
			// 	resourceName := v
			// 	resourceVar, ok := ctx.vars[resourceName]
			// 	if !ok {
			// 		return nil, fmt.Errorf("unknown resource '%v'", resourceName)
			// 	}
			// 	refs = append(refs, model.VariableReference(resourceVar))
			// }
			// dependsOn = append(dependsOn, &model.Block{
			// 	Type: "options",
			// 	Body: &model.Body{
			// 		Items: []model.BodyItem{
			// 			&model.Attribute{
			// 				Name: "dependsOn",
			// 				Value: &model.TupleConsExpression{
			// 					Expressions: refs,
			// 				},
			// 			},
			// 		},
			// 	},
			// })
		}
	}

	if typ == "" || apiVersion == "" {
		r.exclude = true
		diag.Description = "expect type and apiVersion fields to be specified"
		return &diag
	}

	// Can't trust the casing in the template.
	// Need to do a soft-match against resources supported
	token := toResourceToken(typ, apiVersion)
	actual, res, ok := ctx.lookupResource(token)
	if !ok {
		r.exclude = true
		diag.Description = "no metadata found for resource"
		return &diag
	}
	r.resourceToken = actual
	diag.SourceToken = r.resourceToken

	if gen.ShouldExclude(r.resourceToken) {
		r.exclude = true
		diag.Description = "resource doesn't support automatic conversion at this time"
		return &diag
	}

	var err error
	r.resourceParams, err = r.transformRequestBody(ctx, res, r.resourceParams, templateName)
	if err != nil {
		return err
	}

	required := codegen.NewStringSet()
	for _, param := range res.PutParameters {
		if param.IsRequired {
			required.Add(param.Name)
		}
		if param.Body != nil {
			for _, req := range param.Body.RequiredProperties {
				required.Add(req)
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

	if _, ok := r.resourceParams["location"]; !ok {
		if required.Has("location"){
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
	var dependsOn []model.BodyItem

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

	items = append(items, dependsOn...)

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

	// Set the name
	nameParamFromMetadata, err := extractResourceNameParameter(res)
	if err != nil {
		return nil, fmt.Errorf("missing resource name attribute for resource: %s", r.resourceName)
	}
	resourceParams[nameParamFromMetadata.Name] = templateResourceName

	flattened, err := gen.FlattenParams(resourceParams, metadataResParams, ctx.metadata.Types)
	if err != nil {
		return nil, fmt.Errorf("failed to transform parameters for resource: %s: %w", r.resourceName, err)
	}
	return flattened, nil
}

func toResourceToken(resourceType, apiVersion string) string {
	apiVersion = "v" + strings.ReplaceAll(apiVersion, "-", "")
	provAndResource := strings.Split(resourceType, "/")
	prov, resource := provAndResource[0], provAndResource[len(provAndResource)-1]

	if strings.HasPrefix(prov, "Microsoft.") {
		prov = strings.TrimPrefix(strings.ToLower(prov), "microsoft.")
	}
	res := inflector.Singularize(strings.Title(gen.ToLowerCamel(resource)))
	return fmt.Sprintf("azure-nextgen:%s/%s:%s", prov, apiVersion, res)
}

