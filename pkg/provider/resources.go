package provider

import (
	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"
)

// Resource is an ARM resource supporting CRUD
type Resource struct {
	swagggerSpecLocation string
	path                 string
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
	return map[string]Resource{
		"azurerm:containerinstance:ContainerGroup": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2018-10-01/containerInstance.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}",
		},
		"azurerm:compute:VirtualMachine": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/compute/resource-manager/Microsoft.Compute/stable/2018-10-01/compute.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}",
		},
		"azurerm:network:NetworkInterface": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/Microsoft.Network/stable/2018-12-01/networkInterface.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}",
		},
		"azurerm:network:VirtualNetwork": Resource{
			swagggerSpecLocation: "https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/Microsoft.Network/stable/2018-12-01/virtualNetwork.json",
			path:                 "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}",
		},
	}
}
