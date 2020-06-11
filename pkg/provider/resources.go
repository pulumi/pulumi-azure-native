package provider

import (
	"fmt"
	"github.com/gedex/inflector"
	"github.com/pulumi/pulumi-azurerm/pkg/openapi"
	"strings"
)

// Resource is an ARM resource supporting CRUD
type Resource struct {
	swagggerSpecLocation string
	path                 string
	apiVersion           string // TODO: Get this from spec
}

func ResourceMap() (map[string]Resource, error) {
	root := "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/"
	locations := []string {
		"applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/components_API.json",
		"cdn/resource-manager/Microsoft.Cdn/stable/2019-12-31/cdn.json",
		"compute/resource-manager/Microsoft.Compute/stable/2018-10-01/compute.json",
		"containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2018-10-01/containerInstance.json",
		"network/resource-manager/Microsoft.Network/stable/2018-12-01/networkInterface.json",
		"network/resource-manager/Microsoft.Network/stable/2020-04-01/networkSecurityGroup.json",
		"network/resource-manager/Microsoft.Network/stable/2020-04-01/publicIpAddress.json",
		"network/resource-manager/Microsoft.Network/stable/2018-12-01/virtualNetwork.json",
		"redis/resource-manager/Microsoft.Cache/stable/2018-03-01/redis.json",
		"resources/resource-manager/Microsoft.Resources/stable/2018-05-01/resources.json",
		"sql/resource-manager/Microsoft.Sql/stable/2014-04-01/databases.json",
		"sql/resource-manager/Microsoft.Sql/stable/2014-04-01/servers.json",
		"storage/resource-manager/Microsoft.Storage/stable/2019-06-01/blob.json",
		"storage/resource-manager/Microsoft.Storage/stable/2019-06-01/storage.json",
		"web/resource-manager/Microsoft.Web/stable/2019-08-01/AppServicePlans.json",
		"web/resource-manager/Microsoft.Web/stable/2019-08-01/StaticSites.json",
		"web/resource-manager/Microsoft.Web/stable/2019-08-01/WebApps.json",
	}
	result := make(map[string]Resource)

	for _, location := range locations {
		swagggerSpecLocation := root + location
		spec, err := openapi.NewSpec(swagggerSpecLocation)
		if err != nil {
			return nil, err
		}

		for key, path := range spec.Paths.Paths {
			if path.Put == nil || path.Get == nil || path.Delete == nil {
				continue
			}

			typeName := resourceTypeName(key)
			if typeName == "" {
				continue
			}

			providerTypeName := fmt.Sprintf("azurerm:%s", typeName)

			// TODO: store the spec in the resource instead of the location to avoid double-loading.
			result[providerTypeName] = Resource{
				swagggerSpecLocation: swagggerSpecLocation,
				path:                 key,
				apiVersion:           spec.Info.Version,
			}
		}
	}

	return result, nil
}

func resourceTypeName(path string) string {
	if path == "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}" {
		return "core:ResourceGroup"
	}

	parts := strings.Split(path, "/")

	// The path is /subscriptions/id/resoursegroups/id/providers/Microsoft.SomeProvider/foos/id/bars/id/...
	if len(parts) < 8 || len(parts) % 2 != 1 {
		return ""
	}

	// We build a name like someprovider:FooBar.
	// TODO: handle non-Microsoft providers.
	provider := strings.ToLower(parts[6][10:])
	resource := ""
	for i := 7; i < len(parts); i += 2 {
		// TODO: generalize this case to a map of well-known aliases.
		switch strings.ToLower(parts[i]) {
		case "redis":
			resource += "Redis"
		case "sites":
			resource += "AppService"
		case "serverfarms":
			resource += "AppServicePlan"
		default:
			// TODO: we may get better singular names from some metadata.
			resource += strings.Title(inflector.Singularize(parts[i]))
		}
	}
	return fmt.Sprintf("%s:%s", provider, resource)
}
