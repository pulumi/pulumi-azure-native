package provider

import (
	"fmt"
	"github.com/gedex/inflector"
	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azurerm/pkg/openapi"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func TestResourceMap(t *testing.T) {
	swagggerSpecLocations := []string {
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/redis/resource-manager/Microsoft.Cache/stable/2018-03-01/redis.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/compute/resource-manager/Microsoft.Compute/stable/2018-10-01/compute.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2018-10-01/containerInstance.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/resources/resource-manager/Microsoft.Resources/stable/2018-05-01/resources.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/Microsoft.Network/stable/2018-12-01/networkInterface.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/Microsoft.Network/stable/2020-04-01/networkSecurityGroup.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/Microsoft.Network/stable/2020-04-01/publicIpAddress.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/Microsoft.Network/stable/2018-12-01/virtualNetwork.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/Microsoft.Network/stable/2018-12-01/virtualNetwork.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/storage.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/web/resource-manager/Microsoft.Web/stable/2019-08-01/WebApps.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/web/resource-manager/Microsoft.Web/stable/2019-08-01/AppServicePlans.json",
		"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/web/resource-manager/Microsoft.Web/stable/2019-08-01/StaticSites.json",
	}
	result := make(map[string]Resource)

	for _, swagggerSpecLocation := range swagggerSpecLocations {
		spec, err := openapi.NewSpec(swagggerSpecLocation)
		assert.NoError(t, err)

		for key, path := range spec.Paths.Paths {
			if path.Put == nil || path.Get == nil || path.Delete == nil {
				continue
			}

			typeName := resourceTypeName(key)
			if typeName == "" {
				continue
			}

			providerTypeName := fmt.Sprintf("azurerm:%s", typeName)

			result[providerTypeName] = Resource{
				swagggerSpecLocation: swagggerSpecLocation,
				path:                 key,
				apiVersion:           spec.Info.Version,
			}
		}
	}

	var keys []string
	for key, _ := range result {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s\n", key)
	}

}

func resourceTypeName(path string) string {
	if path == "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}" {
		return "core:ResourceGroup"
	}

	parts := strings.Split(path, "/")
	if len(parts) < 8 || len(parts) % 2 != 1 {
		return ""
	}

	provider := strings.ToLower(parts[6][10:])
	resource := ""
	for i := 7; i < len(parts); i += 2 {
		switch strings.ToLower(parts[i]) {
		case "redis":
			resource += "Redis"
		case "sites":
			resource += "AppService"
		case "serverfarms":
			resource += "AppServicePlan"
		default:
			resource += strings.Title(inflector.Singularize(parts[i]))
		}
	}
	return fmt.Sprintf("%s:%s", provider, resource)
}

// TestDiscover is not really a test - it prints the candidates for Azure resources that it finds in specs.
// It also compares the features for each candidate: whether GET and DELETE are supported and whether a resource is
// marked as `x-ms-azure-resource`.
func TestDiscover(t *testing.T) {
	resourceMap := ResourceMap()

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
