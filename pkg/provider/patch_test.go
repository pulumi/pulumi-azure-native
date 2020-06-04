package provider

import (
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
	"net/url"
	"sort"
	"strings"
	"testing"
)

// TestPatchVsPut is not really a test - it prints the differences between the shapes of PATCH and PUT operation for
// different known resources.
func TestPatchVsPut(t *testing.T) {
	resourceMap := ResourceMap()

	keys := make([]string, 0, len(resourceMap))
	for k := range resourceMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		r := resourceMap[name]
		spec, err := loadSwaggerSpec(r.swagggerSpecLocation)
		assert.NoError(t, err)

		path := spec.Paths.Paths[r.path]

		base, err := url.Parse(r.swagggerSpecLocation)
		assert.NoError(t, err)

		resolver := referenceResolver{baseURL: base}
		putParams := resolver.listProperties(t, path.Put, spec)
		fmt.Println(name)
		fmt.Printf("  PUT  : %s\n", strings.Join(putParams, ", "))
		patchParams := resolver.listProperties(t, path.Patch, spec)
		fmt.Printf("  PATCH: %s\n", strings.Join(patchParams, ", "))
		fmt.Println("======")
	}
}

func (r *referenceResolver) listProperties(t *testing.T, operation *spec.Operation, spec *spec.Swagger) []string {
	parameters := make([]string, 0)
	for _, param := range operation.Parameters {
		param, paramSpec, err := r.resolveParameter(param, spec)
		assert.NoError(t, err)

		if param.In == "body" {
			assert.NotNil(t, param.Schema)

			properties, err := r.resolvePropertyMap(*param.Schema, paramSpec, 3)
			assert.NoError(t, err)

			for k, v := range properties {
				if len(v) == 0 {
					parameters = append(parameters, k)
				} else {
					sort.Strings(v)
					p := fmt.Sprintf("%s (%s)", k, strings.Join(v, ","))
					parameters = append(parameters, p)
				}
			}
		}
	}
	sort.Strings(parameters)
	return parameters
}

func (r *referenceResolver) resolvePropertyMap(schema spec.Schema, swagger *spec.Swagger, rec int) (map[string][]string, error) {
	ptr, swagger, ok, err := r.resolveReference(schema.Ref, swagger)
	if err != nil {
		return nil, err
	}
	if ok {
		schema = ptr.(spec.Schema)
	}

	properties := map[string][]string{}

	for k, v := range schema.Properties {
		if v.ReadOnly {
			continue
		}
		subs := make([]string, 0)
		if rec > 0 {
			subProperties, _ := r.resolvePropertyMap(v, swagger, rec-1)
			for sk, _ := range subProperties {
				subs = append(subs, sk)
			}
		}
		properties[k] = subs
	}

	if rec > 0 {
		for _, s := range schema.AllOf {
			ps, err := r.resolvePropertyMap(s, swagger, rec)
			if err != nil {
				return nil, err
			}

			for k, v := range ps {
				properties[k] = v
			}
		}
	}

	return properties, nil
}
