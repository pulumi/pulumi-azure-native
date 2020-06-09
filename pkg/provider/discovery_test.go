package provider

import (
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azurerm/pkg/openapi"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// TestDiscover is not really a test - it prints the candidates for Azure resources that it finds in specs.
// It also compares the features for each candidate: whether GET and DELETE are supported and whether a resource is
// marked as `x-ms-azure-resource`.
func TestDiscover(t *testing.T) {
	resourceMap, err := ResourceMap()
	assert.NoError(t, err)

	keys := make([]string, 0, len(resourceMap))
	for k := range resourceMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	resourceEndpoints := make(map[string]string)
	resourceAttributes := make(map[string]string)

	for _, name := range keys {
		r := resourceMap[name]
		spec, err := openapi.NewSpec(r.swagggerSpecLocation)
		assert.NoError(t, err)

		for key, path := range spec.Paths.Paths {
			if path.Put == nil {
				continue
			}

			switch {
			case path.Get != nil && path.Delete != nil:
				resourceEndpoints[key] = "B"
			case path.Get != nil:
				resourceEndpoints[key] = "G"
			case path.Delete != nil:
				resourceEndpoints[key] = "D"
			default:
				resourceEndpoints[key] = " "
			}

			for code, r := range path.Put.Responses.StatusCodeResponses {

				if code >= 300 || r.Schema == nil {
					continue
				}

				response, err := spec.ResolveResponse(r)
				assert.NoError(t, err)

				if response.Schema == nil {
					continue
				}

				isResource, err := isResource(response.ReferenceContext, *response.Schema)
				assert.NoError(t, err)

				if isResource {
					resourceAttributes[key] = "+"
					break
				}
			}
		}
	}

	for key, value := range resourceEndpoints {
		endpoint := " "
		if v, ok := resourceAttributes[key]; ok {
			endpoint = v
		}
		fmt.Printf("%s%s: %s\n", value, endpoint, key)
	}
}

func isResource(response *openapi.ReferenceContext, s spec.Schema) (bool, error) {
	schema, err := response.ResolveSchema(s)
	if err != nil {
		return false, err
	}

	if v, ok := schema.Extensions["x-ms-azure-resource"]; ok {
		if b, ok := v.(bool); ok && b {
			return true, nil
		}
	}

	for _, sub := range schema.AllOf {
		b, err := isResource(schema.ReferenceContext, sub)
		if err != nil {
			return false, err
		}

		if b {
			return b, nil
		}
	}

	return false, nil
}
