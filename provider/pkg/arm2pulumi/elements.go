package arm2pulumi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/gen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"reflect"
	"sort"
	"strings"

	"github.com/gedex/inflector"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/pcl"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)

// TemplateElement is a top-level item in the template, e.g. parameters, variables,
// resources and outputs.
type TemplateElement interface {
	Name() string
	Args() interface{}
	SetArgs(interface{})
	// RequiresManualConversion is called after all elements are loaded but before
	// any rendering of PCL has occurred. If an element has fundamental issues being
	// converted, it should return true and provide an optional diagnostic.
	RequiresManualConversion(ctx *pclRenderContext) (bool, *Diagnostic)
	PCLExpression(ctx *pclRenderContext) ([]model.BodyItem, error)
}

type pclRenderContext struct {
	metadata *provider.AzureAPIMetadata
	pkgSpec  *schema.PackageSpec
}

// lookupResources looks up the specified resource token and
// returns a corresponding resource if found.
// We don't have every resource in metadata so this does a slower
// lookup for aliases in the package spec if the metadata misses.
// Also, since the templates seem to be case insensitive, we allow
// a slow case insensitive lookup.
func (p *pclRenderContext) lookupResource(resourceToken string) (string, *provider.AzureAPIResource, bool) {
	var actualResourceToken string
	res, ok := p.metadata.Resources[resourceToken]
	if !ok {
		resourceToken = strings.ToLower(resourceToken)
		// first search for case insensitive hit on metadata
		for k := range p.metadata.Resources {
			if strings.ToLower(k) == resourceToken {
				actualResourceToken = k
				break
			}
		}
		// next search for aliases in pkgSpec
		for name, r := range p.pkgSpec.Resources {
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
			res, ok = p.metadata.Resources[actualResourceToken]
		}
	} else {
		actualResourceToken = resourceToken
	}
	return actualResourceToken, &res, ok
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

func (p *parameter) RequiresManualConversion(ctx *pclRenderContext) (bool, *Diagnostic) {
	return false, nil
}

func (p *parameter) PCLExpression(_ *pclRenderContext) ([]model.BodyItem, error) {
	typ, err := extractType(p.rawBody)
	if err != nil {
		return nil, fmt.Errorf("invalid parameter %s: %w", p.paramName, err)
	}

	typeExpr := typ // TODO handle secure variants or return diagnostics for unsupported.

	var defaultValue model.Expression
	if def, ok := p.rawBody["defaultValue"]; ok {
		v, err := pcl.RenderValue(def)
		if err != nil {
			return nil, err
		}
		defaultValue = v
	}

	comment := extractDescription(p.rawBody)
	configDef := &model.Block{
		Type:   "config",
		Labels: []string{p.paramName, typeExpr},
		Body:   &model.Body{},
		Tokens: getLeadingTriviaTokens(comment),
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
	v.rawBody = body
}

func (v *variable) RequiresManualConversion(ctx *pclRenderContext) (bool, *Diagnostic) {
	return false, nil
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

func (o *output) RequiresManualConversion(ctx *pclRenderContext) (bool, *Diagnostic) {
	return false, nil
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

	comment := extractDescription(o.rawBody)
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
		Tokens: getLeadingTriviaTokens(comment),
	}}, nil
}

func NewTemplateElements() *TemplateElements {
	return &TemplateElements{
		elements:               map[string]*dependencyTracking{},
		nameCaseInsensitiveMap: map[string]string{},
		diagnostics:            map[string][]Diagnostic{},
	}
}

// TemplateElements keeps an internal state of template elements to be rendered to PCL,
// including maintaining dependency ordering of the elements
type TemplateElements struct {
	elements               map[string]*dependencyTracking
	nameCaseInsensitiveMap map[string]string
	diagnostics            map[string][]Diagnostic
}

func (t *TemplateElements) lookup(name string) *dependencyTracking {
	if el, ok := t.elements[name]; ok {
		return el
	}
	name, ok := t.nameCaseInsensitiveMap[strings.ToLower(name)]
	if ok {
		return t.elements[name]
	}
	return nil
}

func (t *TemplateElements) recordCaseInsensitiveName(name string) {
	ciName := strings.ToLower(name)
	t.nameCaseInsensitiveMap[ciName] = name
}

// EvaluateExpressions evaluates template expressions. By default, it tries
// to evaluate all templates if preliminaryPhase is true, it will skip resolving
// evaluating templates with certain functions. This is required for instance,
// with functions like reference. Since we may add additional elements due to
// template evaluation and the order of evaluation is random, we wait to perform
// referential evaluations till after all other evaluations are complete.
func (t *TemplateElements) EvaluateExpressions(preliminaryPhase bool) error {
	for name, el := range t.elements {
		args := el.TemplateElement.Args()
		var err error
		if preliminaryPhase {
			args, err = t.evalTemplateExpressions(el, args, "", "reference")
		} else {
			args, err = t.evalTemplateExpressions(el, args, "")
		}
		if err != nil {
			return fmt.Errorf("failed to evaluate template expression for %s: %w", name, err)
		}
		el.SetArgs(args)
	}
	return nil
}

func (t *TemplateElements) evalTemplateExpressions(d *dependencyTracking, args interface{}, crumbs string, excludeFuncs ...string) (interface{}, error) {
	val := reflect.ValueOf(args)
	exclusions := codegen.NewStringSet()
	for _, excl := range excludeFuncs {
		exclusions.Add(excl)
	}
	typ := reflect.TypeOf(args)
	if typ == nil {
		return args, nil
	}
	switch typ.Kind() {
	case reflect.Map:
		ret := map[string]interface{}{}
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			curr := fmt.Sprintf(".%s", k.String())
			updated, err := t.evalTemplateExpressions(d, v.Interface(), curr, excludeFuncs...)
			if err != nil {
				return nil, err
			}
			ret[k.String()] = updated
		}
		return ret, nil
	case reflect.Slice, reflect.Array:
		var ret []interface{}
		for i := 0; i < val.Len(); i++ {
			curr := fmt.Sprintf("%s[%d]", crumbs, i)
			updated, err := t.evalTemplateExpressions(d, val.Index(i).Interface(), curr, excludeFuncs...)
			if err != nil {
				return nil, err
			}
			ret = append(ret, updated)
		}
		return ret, nil
	case reflect.String:
		// json.Number is typed as a string so we must disambiguate.
		if num, ok := args.(json.Number); ok {
			fl, err := num.Float64()
			if err != nil {
				return nil, fmt.Errorf("invalid numeric value at %s: %w", crumbs, err)
			}
			return fl, nil
		}
		s := val.String()
		evaled, err := t.eval(d, s, exclusions)
		if err != nil {
			return nil, err
		}
		return evaled, nil
		//}
	}

	return args, nil
}

// assignName tries to create a unique name elements based on the
// nesting information in crumbs
func (t *TemplateElements) assignName(crumbs []string) string {
	varName := ""
	for _, c := range crumbs {
		varName += strings.Title(c)
	}
	varName = gen.ToLowerCamel(gen.MakeLegalIdentifier(varName))
	if el := t.lookup(varName); el == nil {
		return varName
	}

	base := varName
	i := 0
	for {
		varName = fmt.Sprintf("%s%d", base, i)
		if el := t.lookup(varName); el == nil {
			break
		}
		i++
	}
	return varName
}

func (t *TemplateElements) AddParameter(name string, args map[string]interface{}, addSuffix bool) error {
	if addSuffix {
		if !strings.HasSuffix(strings.ToLower(name), "param") {
			name = name + "Param"
		}
	}

	if el := t.lookup(name); el != nil {
		// Don't expect name collision with params
		return fmt.Errorf("another item with the same name as parameter %s already defined", name)
	}
	p := newParameter(name, args)
	dep := newDependencyTracking(p)

	t.recordCaseInsensitiveName(name)
	t.elements[name] = dep
	return nil
}

func (t *TemplateElements) AddVariable(name string, args interface{}, addSuffix bool) (*dependencyTracking, error) {
	if addSuffix {
		if !strings.HasSuffix(strings.ToLower(name), "var") {
			name = name + "Var"
		}
	}

	if el := t.lookup(name); el != nil {
		return nil, fmt.Errorf("another item with the same name as variable %s already defined", name)
	}

	v := newVariable(name, args)
	dep := newDependencyTracking(v)

	t.recordCaseInsensitiveName(name)
	t.elements[name] = dep

	return dep, nil
}

func (t *TemplateElements) AddOutput(name string, args map[string]interface{}, addSuffix bool) error {
	if addSuffix {
		if !strings.HasSuffix(strings.ToLower(name), "out") {
			name = name + "Out"
		}
	}

	if el := t.lookup(name); el != nil {
		// Don't expect name collision with outputs
		return fmt.Errorf("another item with the same name as output %s already defined", name)
	}

	o := newOutput(name, args)
	dep := newDependencyTracking(o)

	t.recordCaseInsensitiveName(name)
	t.elements[name] = dep
	return nil
}

func (t *TemplateElements) AddResource(args map[string]interface{}) error {
	name, ok := args["name"]
	if !ok {
		return errors.New("missing required name field in resource")
	}

	var prefix []string

	var typestr string
	typ, ok := args["type"]
	if !ok {
		return errors.New("missing required type field in resource")
	}
	typestr = typ.(string)

	// Microsoft.ContainerService/managedClusters -> managedClusterResource
	if !isTemplateExpression(typestr) {
		resource := strings.Split(typestr, "/")
		res := resource[len(resource)-1]
		res = inflector.Singularize(strings.Title(gen.ToLowerCamel(res)))
		prefix = append(prefix, res)
	}
	prefix = append(prefix, "Resource")

	// we often have names as template expressions. We don't want to rely on evaluating the name
	// as the resource variable name so we create a reasonable variable name to house the resource
	// instead. In this case we use the type to help define the variable name, unless that is also
	// a template expression - in which case we just add a 'Resource' suffix.
	// Otherwise just use the resource name field.
	namestr := name.(string)
	if !isTemplateExpression(namestr) {
		prefix = []string{namestr}
	}

	varName := t.assignName(prefix)
	if el := t.lookup(varName); el != nil {
		// Don't expect name collision with resources
		return fmt.Errorf("another resource with name: %s already defined", name)
	}

	var resourceToken string
	apiVersion, ok := args["apiVersion"]
	if !ok {
		return fmt.Errorf("required field 'apiVersion' missing for %s", name)
	}
	apiVersionStr, ok := apiVersion.(string)
	if ok {
		if !isTemplateExpression(apiVersionStr) && !isTemplateExpression(typestr) {
			resourceToken = toResourceToken(typestr, apiVersionStr)
		}
	}

	r := newResource(varName, args, resourceToken)
	dep := newDependencyTracking(r)
	t.elements[varName] = dep
	t.recordCaseInsensitiveName(varName)
	return nil
}

func (t *TemplateElements) handleExclusions(ctx *pclRenderContext) bool {
	haveExclusions := false
	for _, el := range t.elements {
		if exclude, diag := el.RequiresManualConversion(ctx); exclude {
			t.addDiagnostic(*diag)
			haveExclusions = true
		}
	}
	return haveExclusions
}

func (t *TemplateElements) GetDiagnostics() map[string][]Diagnostic {
	return t.diagnostics
}

func (t *TemplateElements) RenderPCL(metadata *provider.AzureAPIMetadata, pkgSpec *schema.PackageSpec) (*model.Body, error) {
	ctx := pclRenderContext{
		metadata: metadata,
		pkgSpec:  pkgSpec,
	}

	t.handleExclusions(&ctx)

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

func extractDescription(info map[string]interface{}) string {
	if metadata, ok := info["metadata"]; ok {
		if metamap, ok := metadata.(map[string]interface{}); ok {
			if desc, ok := metamap["description"]; ok {
				return fmt.Sprintf("%s", desc)
			}
		}
	}
	return ""
}

// TODO: Looks like trivia is never serialized by HCL when rendered... :(
func getLeadingTriviaTokens(comment string) *syntax.BlockTokens {
	if comment == "" {
		return nil
	}
	return &syntax.BlockTokens{
		Type: syntax.Token{
			LeadingTrivia: syntax.TriviaList{&syntax.Comment{Lines: []string{comment}}},
		},
	}
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

	typeStr = strings.ToLower(typeStr)

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
