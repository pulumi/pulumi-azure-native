// Copyright 2016-2020, Pulumi Corporation.

package openapi

import (
	"fmt"
	"net/url"

	"github.com/go-openapi/jsonreference"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
)

// ReferenceContext contains a pointer to a swagger schema and can navigate references from that schema.
// A swagger specification may consist of multiple files with definitions pointing between those files. In order to
// resolve those definitions, we need to keep track of the context where a given reference was defined.
type ReferenceContext struct {
	ReferenceName string
	swagger       *spec.Swagger
	url           *url.URL
}

// Spec is a swagger specification with reference context.
type Spec struct {
	*ReferenceContext
	spec.Swagger
}

// NewSpec load swagger specification from a given location.
func NewSpec(swaggerLocation string) (*Spec, error) {
	base, err := url.Parse(swaggerLocation)
	if err != nil {
		return nil, err
	}

	swagger, err := loadSwaggerSpec(swaggerLocation)
	if err != nil {
		return nil, err
	}

	ctx := &ReferenceContext{swagger: swagger, url: base}
	return &Spec{ctx, *swagger}, nil
}

// Parameter contains a fully resolved swagger parameter and can navigate references from its schema source.
type Parameter struct {
	*ReferenceContext
	*spec.Parameter
}

// Response contains a fully resolved swagger response and can navigate references from its schema source.
type Response struct {
	*ReferenceContext
	*spec.Response
}

// Response contains a fully resolved swagger response and can navigate references from its schema source.
type Schema struct {
	*ReferenceContext
	*spec.Schema
}

// ResolveParameter resolves a given swagger parameter. If needed, it navigates to the source of the parameter reference
// and returns the referenced parameter and its context.
func (ctx *ReferenceContext) ResolveParameter(param spec.Parameter) (*Parameter, error) {
	ptr, ok, err := ctx.resolveReference(param.Ref)
	if err != nil {
		return nil, err
	}
	if !ok {
		return &Parameter{ctx, &param}, nil
	}

	parameter := ptr.value.(spec.Parameter)
	return &Parameter{ptr.ReferenceContext, &parameter}, nil
}

// ResolveResponse resolves a given swagger response. If needed, it navigates to the source of the response reference
// and returns the referenced response and its context.
func (ctx *ReferenceContext) ResolveResponse(resp spec.Response) (*Response, error) {
	ptr, ok, err := ctx.resolveReference(resp.Ref)
	if err != nil {
		return nil, err
	}
	if !ok {
		return &Response{ctx, &resp}, nil
	}

	response := ptr.value.(spec.Response)
	return &Response{ptr.ReferenceContext, &response}, nil
}

// ResolveSchema resolves a given swagger schema. If needed, it navigates to the source of the schema reference
// and returns the referenced schema and its context.
func (ctx *ReferenceContext) ResolveSchema(s *spec.Schema) (*Schema, error) {
	ptr, ok, err := ctx.resolveReference(s.Ref)
	if err != nil {
		return nil, err
	}
	if !ok {
		return &Schema{ctx, s}, nil
	}

	resolvedSchema := ptr.value.(spec.Schema)

	// JSON Reference spec demands that all attributes sibling to a $ref are ignored:
	// > Any members other than "$ref" in a JSON Reference object SHALL be ignored.
	// https://tools.ietf.org/html/draft-pbryan-zyp-json-ref-03#section-3
	// ---------
	// Open API Spec v3 demands the same:
	// > Any sibling elements of a $ref are ignored. This is because $ref works by replacing itself and
	// > everything on its level with the definition it is pointing at.
	// https://swagger.io/docs/specification/using-ref/#sibling
	// ---------
	// However, Open API Spec v2 doesn't seem to define this rule. Azure API specs use a lot
	// of overrides with sibling attributes, so we should take those into account.

	// This group contains the attributes that we know are commonly overridden:
	// Default, ReadOnly, Required, and Extensions; and in one case, AllOf.
	// If the source spec has a value, we merge that value in the resulting specs.
	if s.Default != nil {
		resolvedSchema.Default = s.Default
	}
	resolvedSchema.ReadOnly = resolvedSchema.ReadOnly || s.ReadOnly
	if len(s.Required) > 0 {
		resolvedSchema.Required = s.Required
	}
	if len(s.Extensions) > 0 {
		if resolvedSchema.Extensions == nil {
			resolvedSchema.Extensions = s.Extensions
		} else {
			for k, v := range s.Extensions {
				resolvedSchema.Extensions[k] = v
			}
		}
	}

	if len(s.AllOf) > 0 {
		resolvedSchema.AllOf = s.AllOf
	}

	// All the other properties aren't currently overridden. We add an assertion, so that
	// if a new specification does override a value, we can catch this and decide what to do further.
	if s.Maximum != nil {
		return nil, errors.New("'Maximum' defined as a sibling to a $ref")
	}
	if s.Minimum != nil {
		return nil, errors.New("'Minimum' defined as a sibling to a $ref")
	}
	if s.MaxLength != nil {
		return nil, errors.New("'MaxLength' defined as a sibling to a $ref")
	}
	if s.MinLength != nil {
		return nil, errors.New("'MinLength' defined as a sibling to a $ref")
	}
	if len(s.Pattern) > 0 {
		return nil, errors.New("'Pattern' defined as a sibling to a $ref")
	}
	if len(s.Discriminator) > 0 {
		return nil, errors.New("'Discriminator' defined as a sibling to a $ref")
	}
	if len(s.Enum) > 0 {
		return nil, errors.New("'Enum' defined as a sibling to a $ref")
	}
	if s.Items != nil {
		return nil, errors.New("'Items' defined as a sibling to a $ref")
	}
	if s.Properties != nil {
		return nil, errors.New("'Properties' defined as a sibling to a $ref")
	}
	if s.AdditionalProperties != nil {
		return nil, errors.New("'AdditionalProperties' defined as a sibling to a $ref")
	}

	// Note that many other Open API schema properties aren't validated above
	// because those aren't used in our code generation, or in Azure specs in general.

	return &Schema{ptr.ReferenceContext, &resolvedSchema}, nil
}

// FindSubtypes returns a slice of schemas, each schema is a reference to
// a type definition that is a subtype of a given type.
// The following rules apply:
// - All subtypes reside in the same Open API specification.
// - A subtype is defines as `allOf` of the base type.
// - The search is applied recursively and all types along `allOf` hiereachy are returned.
func (ctx *ReferenceContext) FindSubtypes() ([]*spec.Schema, error) {
	var schemas []*spec.Schema
	for _, name := range codegen.SortedKeys(ctx.swagger.Definitions) {
		def := ctx.swagger.Definitions[name]
		subTypes, err := ctx.recursiveAllOf(&def)
		if err != nil {
			return nil, err
		}
		for _, schema := range subTypes {
			if resolved, ok, _ := ctx.resolveReference(schema.Ref); ok {
				if resolved.ReferenceName == ctx.ReferenceName {
					ref := spec.Schema{
						SchemaProps: spec.SchemaProps{
							Ref: spec.Ref{
								Ref: jsonreference.MustCreateRef("#/definitions/" + name),
							},
						},
					}
					schemas = append(schemas, &ref)
				}
			}
		}
	}
	return schemas, nil
}

func (ctx *ReferenceContext) recursiveAllOf(def *spec.Schema) ([]spec.Schema, error) {
	var result []spec.Schema
	for _, ref := range def.AllOf {
		result = append(result, ref)
		schema, err := ctx.ResolveSchema(&ref)
		if err != nil {
			return nil, err
		}
		children, err := schema.ReferenceContext.recursiveAllOf(schema.Schema)
		if err != nil {
			return nil, err
		}
		result = append(result, children...)
	}
	return result, nil
}

// MergeParameters combines the Path Item parameters with Operation parameters.
func (ctx *ReferenceContext) MergeParameters(operation []spec.Parameter, pathItem []spec.Parameter) []spec.Parameter {
	// Open API spec for operations:
	// > If a parameter is already defined at the Path Item, the new definition will override it.
	// > A unique parameter is defined by a combination of a name and location.
	var result []spec.Parameter
	seen := map[string]bool{}
	for _, p := range operation {
		schema, err := ctx.ResolveParameter(p)
		if err != nil {
			panic(err)
		}
		key := fmt.Sprintf("%s@%s", schema.Name, schema.In)
		seen[key] = true
		result = append(result, p)
	}
	for _, p := range pathItem {
		schema, err := ctx.ResolveParameter(p)
		if err != nil {
			panic(err)
		}
		key := fmt.Sprintf("%s@%s", schema.Name, schema.In)
		if _, ok := seen[key]; !ok {
			result = append(result, p)
		}
	}
	return result
}

// ResolveReference resolves a relative reference relative to current swagger spec URL
func (ctx *ReferenceContext) ResolveReference(ref string) (string, error) {
	relativeURL, err := url.Parse(ref)
	if err != nil {
		return "", err
	}
	return ctx.url.ResolveReference(relativeURL).String(), nil
}

type reference struct {
	*ReferenceContext
	value interface{}
}

func (ctx *ReferenceContext) resolveReference(ref spec.Ref) (*reference, bool, error) {
	ptr := ref.GetPointer()
	if ptr == nil || ptr.IsEmpty() {
		return nil, false, nil
	}

	referenceName := ptr.DecodedTokens()[len(ptr.DecodedTokens())-1]
	relative, err := url.Parse(ref.RemoteURI())
	if err == nil && ref.RemoteURI() != "" {
		finalURL := ctx.url.ResolveReference(relative)
		swagger, err := loadSwaggerSpec(finalURL.String())
		if err != nil {
			return nil, false, errors.Wrapf(err, "load Swagger spec")
		}

		ctx = &ReferenceContext{swagger: swagger, ReferenceName: referenceName, url: finalURL}
	}

	value, _, err := ptr.Get(ctx.swagger)
	if err != nil {
		return nil, false, errors.Wrapf(err, "get pointer")
	}

	newCtx := &ReferenceContext{swagger: ctx.swagger, ReferenceName: referenceName, url: ctx.url}
	return &reference{newCtx, value}, true, nil
}

// Cache of parsed Swagger specifications for a location.
var specCache = map[string]*spec.Swagger{}

func loadSwaggerSpec(path string) (*spec.Swagger, error) {
	if cached, ok := specCache[path]; ok {
		return cached, nil
	}

	byts, err := swag.LoadFromFileOrHTTP(path)
	if err != nil {
		return nil, err
	}
	swagger := spec.Swagger{}
	err = swagger.UnmarshalJSON(byts)
	if err != nil {
		return nil, err
	}

	specCache[path] = &swagger
	return &swagger, nil
}
