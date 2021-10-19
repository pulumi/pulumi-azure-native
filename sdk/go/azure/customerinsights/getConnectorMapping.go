


package customerinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectorMapping(ctx *pulumi.Context, args *LookupConnectorMappingArgs, opts ...pulumi.InvokeOption) (*LookupConnectorMappingResult, error) {
	var rv LookupConnectorMappingResult
	err := ctx.Invoke("azure-native:customerinsights:getConnectorMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorMappingArgs struct {
	ConnectorName     string `pulumi:"connectorName"`
	HubName           string `pulumi:"hubName"`
	MappingName       string `pulumi:"mappingName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectorMappingResult struct {
	ConnectorMappingName string                             `pulumi:"connectorMappingName"`
	ConnectorName        string                             `pulumi:"connectorName"`
	ConnectorType        *string                            `pulumi:"connectorType"`
	Created              string                             `pulumi:"created"`
	DataFormatId         string                             `pulumi:"dataFormatId"`
	Description          *string                            `pulumi:"description"`
	DisplayName          *string                            `pulumi:"displayName"`
	EntityType           string                             `pulumi:"entityType"`
	EntityTypeName       string                             `pulumi:"entityTypeName"`
	Id                   string                             `pulumi:"id"`
	LastModified         string                             `pulumi:"lastModified"`
	MappingProperties    ConnectorMappingPropertiesResponse `pulumi:"mappingProperties"`
	Name                 string                             `pulumi:"name"`
	NextRunTime          string                             `pulumi:"nextRunTime"`
	RunId                string                             `pulumi:"runId"`
	State                string                             `pulumi:"state"`
	TenantId             string                             `pulumi:"tenantId"`
	Type                 string                             `pulumi:"type"`
}
