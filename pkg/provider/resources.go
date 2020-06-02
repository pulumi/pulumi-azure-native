package provider

import (
	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"
)

// Resource is an ARM resource supporting CRUD
type Resource struct {
	swagggerSpecLocation string
	path                 string
	apiVersion           string // TODO: Get this from spec
}

// LoadSpec loads the Swagger spec for the resource
func (r *Resource) LoadSpec() (*spec.Swagger, error) {
	return loadSwaggerSpec(r.swagggerSpecLocation)
}

func loadSwaggerSpec(path string) (*spec.Swagger, error) {
	byts, err := swag.LoadFromFileOrHTTP(path)
	if err != nil {
		return nil, err
	}
	swagger := spec.Swagger{}
	err = swagger.UnmarshalJSON(byts)
	if err != nil {
		return nil, err
	}
	return &swagger, nil
}

// ResourceMap gets the full map of ARM resource supported by this provider
func ResourceMap() map[string]Resource {
	// TODO:  Resources here are currently manually mapped to REST APIs.  Ideally we would automatically pick these up.
	// There are two sources:
	//
	// * The REST API specs in https://github.om/Azure/azure-rest-api-specs used below
	//
	// * The provider and resource type list returned from
	// https://management.azure.com/subscriptions/{subscriptionId}/providers which provides the nouns that ARM maps to
	// resource types in ARM templates.  (the same results as `az provider list`)
	//
	// The latter more closely describes what's available from ARM templates, but there doesn't apepar to be any clear
	// way to get from those resource types to the underlying REST API endpoints.  Since we use the REST API endponits
	// directly to do CRUD, we need to somehow construct this mapping.
	//
	// It may be that we want to use both sources, slurp up all of the REST API specs, and then look up resources from
	// the provider/resorucetype list in the ersulting API specs.
	//
	// Or it could be tht just finding all of resources in the  REST API specs that support `put`+`get`+`delete` is
	// sufficient, and that we don't need to use the provider list at all.  We would need to figure out how to construct
	// the name - roughly it's the following two transformations - which seem close to being automatable:
	//
	// * `Microsoft.ContainerInstance` => `containerinstance` (from resource namespace)
	//
	// * `containerGroups` => `ContainerGroup`(from the second to last part of the resource path, also known as
	// "resource type" in the `/providers` API)
	return map[string]Resource{
		"azurerm:cache:Redis": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/redis/resource-manager/Microsoft.Cache/stable/2018-03-01/redis.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}",
			apiVersion:           "2018-03-01",
		},
		"azurerm:compute:VirtualMachine": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/compute/resource-manager/Microsoft.Compute/stable/2018-10-01/compute.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}",
			apiVersion:           "2018-10-01",
		},
		"azurerm:containerinstance:ContainerGroup": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2018-10-01/containerInstance.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}",
			apiVersion:           "2018-10-01",
		},
		"azurerm:core:ResourceGroup": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/resources/resource-manager/Microsoft.Resources/stable/2018-05-01/resources.json",
			path:                 "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}",
			apiVersion:           "2018-05-01",
		},
		"azurerm:network:NetworkInterface": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/Microsoft.Network/stable/2018-12-01/networkInterface.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}",
			apiVersion:           "2018-12-01",
		},
		"azurerm:network:NetworkSecurityGroup": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/Microsoft.Network/stable/2020-04-01/networkSecurityGroup.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}",
			apiVersion:           "2020-04-01",
		},
		"azurerm:network:PublicIPAddress": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/Microsoft.Network/stable/2020-04-01/publicIpAddress.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}",
			apiVersion:           "2020-04-01",
		},
		"azurerm:network:VirtualNetwork": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/Microsoft.Network/stable/2018-12-01/virtualNetwork.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}",
			apiVersion:           "2018-12-01",
		},
		"azurerm:storage:Account": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/storage.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}",
			apiVersion:           "2019-06-01",
		},
		"azurerm:web:AppService": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/web/resource-manager/Microsoft.Web/stable/2019-08-01/WebApps.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}",
			apiVersion:           "2019-08-01",
		},
		"azurerm:web:AppServicePlan": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/web/resource-manager/Microsoft.Web/stable/2019-08-01/AppServicePlans.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}",
			apiVersion:           "2019-08-01",
		},
		"azurerm:web:StaticSite": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/619f12bbb9d1092c9c126f4bec54f1f0c718ca58/specification/web/resource-manager/Microsoft.Web/stable/2019-08-01/StaticSites.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/staticSites/{name}",
			apiVersion:           "2019-08-01",
		},
	}
}
