package provider

import (
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azurerm/pkg/openapi"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

// TestPatchVsPut is not really a test - it prints the differences between the shapes of PATCH and PUT operation for
// different known resources.
func TestPatchVsPut(t *testing.T) {
	resourceMap, err := ResourceMap()
	assert.NoError(t, err)

	keys := make([]string, 0, len(resourceMap))
	for k := range resourceMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		r := resourceMap[name]
		spec, err := openapi.NewSpec(r.swagggerSpecLocation)
		assert.NoError(t, err)

		path := spec.Swagger.Paths.Paths[r.path]

		putParams := listProperties(t, spec, path.Put)
		fmt.Println(name)
		fmt.Printf("  PUT  : %s\n", strings.Join(putParams, ", "))
		patchParams := listProperties(t, spec, path.Patch)
		fmt.Printf("  PATCH: %s\n", strings.Join(patchParams, ", "))
		fmt.Println("======")
	}
}

func listProperties(t *testing.T, spec *openapi.Spec, operation *spec.Operation) []string {
	parameters := make([]string, 0)
	for _, param := range operation.Parameters {
		param, err := spec.ResolveParameter(param)
		assert.NoError(t, err)

		if param.In == "body" {
			assert.NotNil(t, param.Schema)

			properties, err := resolvePropertyMap(spec.ReferenceContext, *param.Schema, 3)
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

func resolvePropertyMap(spec *openapi.ReferenceContext, s spec.Schema, rec int) (map[string][]string, error) {
	schema, err := spec.ResolveSchema(s)
	if err != nil {
		return nil, err
	}

	properties := map[string][]string{}

	for k, v := range schema.Properties {
		if v.ReadOnly {
			continue
		}
		subs := make([]string, 0)
		if rec > 0 {
			subProperties, _ := resolvePropertyMap(schema.ReferenceContext, v, rec-1)
			for sk, _ := range subProperties {
				subs = append(subs, sk)
			}
		}
		properties[k] = subs
	}

	if rec > 0 {
		for _, s := range schema.AllOf {
			ps, err := resolvePropertyMap(schema.ReferenceContext, s, rec-1)
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
