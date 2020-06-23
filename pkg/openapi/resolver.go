package openapi

import (
	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
	"net/url"
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

	schema := ptr.value.(spec.Schema)
	return &Schema{ptr.ReferenceContext, &schema}, nil
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

var cache map[string]*spec.Swagger

func init() {
	cache = map[string]*spec.Swagger{}
}

func loadSwaggerSpec(path string) (*spec.Swagger, error) {
	if cached, ok := cache[path]; ok {
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

	cache[path] = &swagger
	return &swagger, nil
}
