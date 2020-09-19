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
	}
}

type resource struct {
	resourceName  string
	rawBody       map[string]interface{}
	resourceToken string // initial guess at pulumi resource token
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

func (r *resource) RequiresManualConversion(ctx *pclRenderContext) (bool, *Diagnostic) {
	actualToken, _, found := ctx.lookupResource(r.resourceToken)
	diag := Diagnostic{
		SourceElement: r.resourceName,
		SourceToken:   r.resourceToken,
		Severity:      SevHigh,
		Description:   "this resource can't be automatically converted at this time",
	}
	if !found || gen.ShouldExclude(actualToken) {
		r.exclude = true
		return true, &diag
	}

	return false, nil
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
	var resourceParams map[string]interface{}
	var location interface{}
	var apiVersion, typ string
	var dependsOn []model.BodyItem

	templateResourceName, ok := r.rawBody["name"]
	if !ok {
		return nil, fmt.Errorf("missing required 'name' attribute for resource: %s", r.resourceName)
	}

	for k, v := range r.rawBody {
		switch k {
		case "apiVersion":
			var ok bool
			apiVersion, ok = v.(string)
			if !ok {
				return nil, fmt.Errorf("expect apiVersion to be specified as a string, received %T", v)
			}
		case "type":
			var ok bool
			typ, ok = v.(string)
			if !ok {
				return nil, fmt.Errorf("expect type to be specified as a string, received %T", v)
			}
		case "location":
			// location could be a string or a variable reference. Grab it so we can inject
			// into properties of the resource if missing
			location = v
		case "properties":
			resourceParams = v.(map[string]interface{})
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
		return nil, fmt.Errorf("expect type and apiVersion fields to be specified for resource: %s", r.resourceName)
	}

	// Can't trust the casing in the template.
	// Need to do a soft-match against resources supported
	token := toResourceToken(typ, apiVersion)
	actual, res, ok := ctx.lookupResource(token)
	if !ok {
		return nil, fmt.Errorf("no metadata found for resource %s", token)
	}
	token = actual // Use the actual token going forward

	metadataResParams := map[string]provider.AzureAPIParameter{}
	for _, param := range res.PutParameters {
		metadataResParams[param.Name] = param
	}
	_, foundProperties := metadataResParams["properties"]

	// block specified under properties seem to be converted to body payload
	// so determine the body compliant field against metadata if a direct match
	// against "properties" is not found
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
	nameParamFromMetadata, err := extractResourceNameParameter(res)
	if err != nil {
		return nil, fmt.Errorf("missing resource name attribute for resource: %s", r.resourceName)
	}

	resourceParams[nameParamFromMetadata.Name] = templateResourceName

	flattened, err := gen.FlattenInput(resourceParams, metadataResParams, ctx.metadata.Types)
	if err != nil {
		return nil, fmt.Errorf("failed to transform parameters for resource: %s: %w", r.resourceName, err)
	}

	if _, found := flattened["location"]; !found {
		flattened["location"] = location
	}

	keys := codegen.SortedKeys(flattened)
	for _, k := range keys {
		v := flattened[k]
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
		Labels: []string{r.resourceName, token},
	}
	return []model.BodyItem{&block}, nil
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

