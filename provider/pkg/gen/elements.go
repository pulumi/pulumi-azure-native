package gen

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/gedex/inflector"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/pcl"
	"github.com/pulumi/pulumi-azurerm/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

type TemplateElement interface {
	Name() string
	Args() interface{}
	SetArgs(interface{})
	PCLExpression(ctx *pclRenderContext) ([]model.BodyItem, error)
}

type pclRenderContext struct {
	vars     map[string]*model.Variable
	metadata *provider.AzureApiMetadata
}

func newDependencyTracking(e TemplateElement) *dependencyTracking {
	return &dependencyTracking{
		TemplateElement: e,
		dependencies:    map[*dependencyTracking]bool{},
	}
}

type dependencyTracking struct {
	TemplateElement
	dependencies map[*dependencyTracking]bool // things this item references (dependencies, i.e. must precede in topo sort order)
}

func (d *dependencyTracking) RefersTo(ref *dependencyTracking) {
	d.dependencies[ref] = true
}

func (d *dependencyTracking) Dependencies() []*dependencyTracking {
	var result []*dependencyTracking
	for k := range d.dependencies {
		result = append(result, k)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name() < result[j].Name()
	})

	return result
}

func newParameter(name string, rawBody map[string]interface{}) *parameter {
	return &parameter{
		paramName: name,
		rawBody:   rawBody,
	}
}

type parameter struct {
	paramName string
	rawBody   map[string]interface{}
}

func (p *parameter) Name() string {
	return p.paramName
}

func (p *parameter) Args() interface{} {
	return p.rawBody
}

func (p *parameter) SetArgs(body interface{}) {
	p.rawBody = body.(map[string]interface{})
}

func (p *parameter) PCLExpression(_ *pclRenderContext) ([]model.BodyItem, error) {
	typ, err := extractType(p.rawBody)
	if err != nil {
		return nil, fmt.Errorf("invalid parameter %s: %w", p.paramName, err)
	}

	typeExpr := typ // TODO handle secure variants

	var defaultValue model.Expression
	if def, ok := p.rawBody["defaultValue"]; ok {
		v, err := pcl.RenderValue(def)
		if err != nil {
			return nil, err
		}
		defaultValue = v
	}

	configDef := &model.Block{
		Type:   "config",
		Labels: []string{p.paramName, typeExpr},
		Body:   &model.Body{},
	}

	if defaultValue != nil {
		configDef.Body.Items = append(configDef.Body.Items, &model.Attribute{
			Name:  "default",
			Value: defaultValue,
		})
	}
	return []model.BodyItem{configDef}, nil
}

func newVariable(name string, rawBody interface{}) *variable {
	return &variable{
		varName: name,
		rawBody: rawBody,
	}
}

type variable struct {
	varName string
	rawBody interface{}
}

func (v *variable) Name() string {
	return v.varName
}

func (v *variable) Args() interface{} {
	return v.rawBody
}

func (v *variable) SetArgs(body interface{}) {
	v.rawBody = body.(map[string]interface{})
}

func (v *variable) PCLExpression(_ *pclRenderContext) ([]model.BodyItem, error) {
	fObj := v.rawBody
	m, err := pcl.RenderValue(fObj)
	if err != nil {
		return nil, err
	}

	return []model.BodyItem{&model.Attribute{
		Name:  v.varName,
		Value: m,
	}}, nil
}

func newOutput(name string, rawBody map[string]interface{}) *output {
	return &output{
		outputName: name,
		rawBody:    rawBody,
	}
}

type output struct {
	outputName string
	rawBody    map[string]interface{}
}

func (o *output) Name() string {
	return o.outputName
}

func (o *output) Args() interface{} {
	return o.rawBody
}

func (o *output) SetArgs(body interface{}) {
	o.rawBody = body.(map[string]interface{})
}

func (o *output) PCLExpression(_ *pclRenderContext) ([]model.BodyItem, error) {
	value := o.rawBody

	if cond, ok := value["condition"]; ok {
		if keep, ok := cond.(bool); !ok {
			return nil, fmt.Errorf("expect condition field in output %s to be boolean, received: %T", o.outputName, cond)
		} else if !keep {
			return nil, nil
		}
	}

	// TODO: Need to add support for copy
	val, ok := value["value"]
	if !ok {
		return nil, fmt.Errorf("output '%v' has no \"value\" attribute", o.outputName)
	}

	x, err := pcl.RenderValue(val)
	if err != nil {
		return nil, err
	}

	// TODO: Add support for secure variants as type

	return []model.BodyItem{&model.Block{
		Type:   "output",
		Labels: []string{o.outputName},
		Body: &model.Body{
			Items: []model.BodyItem{
				&model.Attribute{
					Name:  "value",
					Value: x,
				},
			},
		},
	}}, nil
}

func newResource(name string, rawBody map[string]interface{}) *resource {
	return &resource{
		resourceName: name,
		rawBody:      rawBody,
	}
}

type resource struct {
	resourceName string
	rawBody      map[string]interface{}
}

func (r *resource) Name() string {
	return r.resourceName
}

func (r *resource) PCLExpression(ctx *pclRenderContext) ([]model.BodyItem, error) {
	var items []model.BodyItem

	var resourceParams map[string]interface{}
	var location interface{}
	var apiVersion, typ string
	var dependsOn []model.BodyItem
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

	if resourceParams == nil {
		return nil, fmt.Errorf("expect parameters for resource %s", r.resourceName)
	}
	if typ == "" || apiVersion == "" {
		return nil, fmt.Errorf("expect type and apiVersion fields to be specified for resource: %s", r.resourceName)
	}

	token := toResourceToken(typ, apiVersion)
	res, ok := ctx.metadata.Resources[token]
	if !ok {
		return nil, fmt.Errorf("no metadata found for resource %s", token)
	}

	metadataResParams := map[string]provider.AzureApiParameter{}
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

	flattened, err := flattenInput(resourceParams, metadataResParams, ctx.metadata.Types)
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
	block := model.Block{
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
	res := inflector.Singularize((strings.Title(toLowerCamel(resource))))
	return fmt.Sprintf("azurerm:%s/%s:%s", prov, apiVersion, res)
}

func NewTemplateElements() *TemplateElements {
	return &TemplateElements{
		elements:                 map[string]*dependencyTracking{},
		knownTemplateExpressions: map[string]*model.Variable{},
	}
}

// TemplateElements keeps an internal state of template elements to be rendered to PCL,
// including maintaining dependency ordering of the elements
type TemplateElements struct {
	elements                 map[string]*dependencyTracking
	knownTemplateExpressions map[string]*model.Variable
}

func (t *TemplateElements) extractTemplateExpressionsAsVars(args interface{}, crumbs []string, excludeKeys ...string) (interface{}, map[string]model.Expression, error) {
	variables := map[string]model.Expression{}
	val := reflect.ValueOf(args)
	exclusions := codegen.NewStringSet()
	for _, excl := range excludeKeys {
		exclusions.Add(excl)
	}
	switch typ := reflect.TypeOf(args); typ.Kind() {
	case reflect.Map:
		ret := map[string]interface{}{}
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			if exclusions.Has(k.String()) {
				ret[k.String()] = v.Interface()
				continue
			}
			nextCrumbs := append(crumbs, k.String())
			updated, vars, err := t.extractTemplateExpressionsAsVars(v.Interface(), nextCrumbs)
			if err != nil {
				return nil, nil, err
			}
			ret[k.String()] = updated
			for k, v := range vars {
				variables[k] = v
			}
		}
		return ret, variables, nil
	case reflect.Slice, reflect.Array:
		var ret []interface{}
		for i := 0; i < val.Len(); i++ {
			updated, vars, err := t.extractTemplateExpressionsAsVars(val.Index(i).Interface(), crumbs)
			if err != nil {
				return nil, nil, err
			}
			ret = append(ret, updated)
			for k, v := range vars {
				variables[k] = v
			}
		}
		return ret, variables, nil
	case reflect.String:
		s := val.String()
		if isTemplateExpression(s) {
			// Check if we already have an existing variable tracking
			// the template expression
			if exprVar, seen := t.knownTemplateExpressions[s]; seen {
				return model.VariableReference(exprVar),
					map[string]model.Expression{
						exprVar.Name: nil, // No need for tracking params since var already exists
					}, nil
			}
			varName := t.assignName(crumbs)
			v := &model.Variable{
				Name:         varName,
				VariableType: model.DynamicType,
			}
			val, err := pcl.RenderValue(s)
			if err != nil {
				return nil, nil, err
			}
			// Record template to variable mapping so we avoid
			// emitting duplicates
			t.knownTemplateExpressions[s] = v
			return model.VariableReference(v),
				map[string]model.Expression{
					varName: val,
				}, nil
		}
	}

	return args, nil, nil
}

func (t *TemplateElements) assignName(crumbs []string) string {
	varName := ""
	for _, c := range crumbs {
		varName += strings.Title(c)
	}
	varName = toLowerCamel(makeLegalIdentifier(varName))
	if _, has := t.elements[varName]; !has {
		return varName
	}

	base := varName
	i := 0
	for {
		varName = fmt.Sprintf("%s%d", base, i)
		if _, has := t.elements[varName]; !has {
			break
		}
		i++
	}
	return varName
}

func (t *TemplateElements) detectTemplateExpressions(srcName string, in *dependencyTracking, args interface{}, exclusionList ...string) (interface{}, error) {
	typ := reflect.TypeOf(args)
	kind := typ.Kind()

	var exprs map[string]model.Expression
	var err error
	switch kind {
	case reflect.Map, reflect.Slice, reflect.String:
		args, exprs, err = t.extractTemplateExpressionsAsVars(
			args,
			[]string{srcName},
			exclusionList...)
		if err != nil {
			return nil, err
		}
	}

	for name, expr := range exprs {
		dep, ok := t.elements[name]
		if !ok {
			v := newVariable(name, expr)
			dep = newDependencyTracking(v)
			t.elements[name] = dep
		}
		in.RefersTo(dep)
	}

	return args, nil
}

func (t *TemplateElements) AddParameter(name string, args map[string]interface{}) error {
	if _, exists := t.elements[name]; exists {
		// Don't expect name collision with params
		return fmt.Errorf("another item with the same name as parameter %s already defined", name)
	}

	p := newParameter(name, args)
	dep := newDependencyTracking(p)

	if def, ok := args["defaultValue"]; ok {
		var err error
		args["defaultValue"], err = t.detectTemplateExpressions(name+"Default", dep, def)
		if err != nil {
			return fmt.Errorf("failed to extract templates in parameter %s: %w", name, err)
		}
	}
	p.rawBody = args

	t.elements[name] = dep
	return nil
}

func (t *TemplateElements) AddVariable(name string, args interface{}) error {
	if _, exists := t.elements[name]; exists {
		return fmt.Errorf("another item with the same name as variable %s already defined", name)
	}
	v := newVariable(name, args)
	dep := newDependencyTracking(v)
	var err error
	v.rawBody, err = t.detectTemplateExpressions(name, dep, args)
	if err != nil {
		return fmt.Errorf("failed to extract templates in variable %s: %w", name, err)
	}

	t.elements[name] = dep
	return nil
}

func (t *TemplateElements) AddOutput(name string, args map[string]interface{}) error {
	if _, exists := t.elements[name]; exists {
		// Don't expect name collision with outputs
		return fmt.Errorf("another item with the same name as output %s already defined", name)
	}
	o := newOutput(name, args)

	dep := newDependencyTracking(o)
	var err error
	val, ok := args["value"]
	if ok {
		args["value"], err = t.detectTemplateExpressions(name+"Value", dep, val)
		if err != nil {
			return fmt.Errorf("failed to extract templates in output %s: %w", name, err)
		}
	}
	o.rawBody = args

	t.elements[name] = dep
	return nil
}

func (t *TemplateElements) AddResource(args map[string]interface{}) error {
	name, ok := args["name"]
	if !ok {
		return errors.New("missing required name field in resource")
	}

	namestr := name.(string)
	prefix := []string{}

	if !isTemplateExpression(namestr) {
		prefix = []string{namestr}
	} else {
		typ, ok := args["type"]
		if !ok {
			return errors.New("missing required type field in resource")
		}
		typestr := typ.(string)

		if !isTemplateExpression(typestr) {
			resource := strings.Split(typestr, "/")
			res := resource[len(resource)-1]
			res = inflector.Singularize((strings.Title(toLowerCamel(res))))
			prefix = append(prefix, res)
		}
		prefix = append(prefix, "resource")
	}

	varName := t.assignName(prefix)

	if _, exists := t.elements[varName]; exists {
		// Don't expect name collision with resources
		return fmt.Errorf("another resource with name: %s already defined", name)
	}
	r := newResource(varName, args)

	dep := newDependencyTracking(r)
	// Delay processing of dependsOn
	replaced, err := t.detectTemplateExpressions(varName, dep, args, "dependsOn")
	if err != nil {
		return fmt.Errorf("failed to extract templates in resource %s: %w", name, err)
	}

	r.rawBody = replaced.(map[string]interface{})

	t.elements[varName] = dep
	return nil
}

func (t *TemplateElements) RenderPCL(metadata *provider.AzureApiMetadata) (*model.Body, error) {
	ctx := pclRenderContext{
		vars:     map[string]*model.Variable{},
		metadata: metadata,
	}
	for k := range t.elements {
		ctx.vars[k] = &model.Variable{
			Name:         k,
			VariableType: model.DynamicType,
		}
	}

	elements, err := t.topologicalOrder()
	if err != nil {
		return nil, err
	}

	var items []model.BodyItem
	for _, el := range elements {
		bodyItems, err := el.PCLExpression(&ctx)
		if err != nil {
			return nil, err
		}
		items = append(items, bodyItems...)
	}

	return &model.Body{
		Items: items,
	}, nil
}

func (t *TemplateElements) topologicalOrder() ([]*dependencyTracking, error) {
	visited := map[*dependencyTracking]bool{}
	var orderedElements []*dependencyTracking
	keys := codegen.SortedKeys(t.elements)
	for _, k := range keys {
		v := t.elements[k]
		var err error
		var ordered []*dependencyTracking
		ordered, err = topoSort(v, visited, nil)
		if err != nil {
			return nil, err
		}

		for _, o := range ordered {
			orderedElements = append(orderedElements, o)
		}

		var names []string
		for _, o := range ordered {
			names = append(names, o.Name())
		}
	}
	return orderedElements, nil
}

func topoSort(e *dependencyTracking, visited map[*dependencyTracking]bool, order []*dependencyTracking) ([]*dependencyTracking, error) {
	committed, seen := visited[e]
	if committed {
		return order, nil
	}
	if seen {
		return nil, fmt.Errorf("detected cycle at element %s", e.Name())
	}
	visited[e] = false // Mark as visited (but not committed)
	var sorted []*dependencyTracking
	for _, el := range e.Dependencies() {
		s, err := topoSort(el, visited, order)
		if err != nil {
			return nil, err
		}
		sorted = append(sorted, s...)
	}
	sorted = append(sorted, e)
	visited[e] = true // Mark as committed
	return sorted, nil
}

func extractType(info map[string]interface{}) (string, error) {
	typ, ok := info["type"]
	if !ok {
		return "", errors.New("missing required field 'type'")
	}

	typeStr, ok := typ.(string)
	if !ok {
		return "", errors.New("expect type to be a string")
	}

	switch typeStr {
	case "int":
		return "number", nil
	case "bool", "string", "object", "array":
		return typeStr, nil
	case "securestring":
		return "string", nil // TODO: Fix this
	case "secureobject":
		return "object", nil // TODO: Fix this
	default:
		return "", fmt.Errorf("unexpected type: %s", typeStr)
	}
}
