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
	byts, err := swag.LoadFromFileOrHTTP(r.swagggerSpecLocation)
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
	}
}
