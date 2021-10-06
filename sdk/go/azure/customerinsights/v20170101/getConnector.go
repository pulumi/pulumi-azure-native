


package v20170101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	var rv LookupConnectorResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	ConnectorName     string `pulumi:"connectorName"`
	HubName           string `pulumi:"hubName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectorResult struct {
	ConnectorId         int                    `pulumi:"connectorId"`
	ConnectorName       *string                `pulumi:"connectorName"`
	ConnectorProperties map[string]interface{} `pulumi:"connectorProperties"`
	ConnectorType       string                 `pulumi:"connectorType"`
	Created             string                 `pulumi:"created"`
	Description         *string                `pulumi:"description"`
	DisplayName         *string                `pulumi:"displayName"`
	Id                  string                 `pulumi:"id"`
	IsInternal          *bool                  `pulumi:"isInternal"`
	LastModified        string                 `pulumi:"lastModified"`
	Name                string                 `pulumi:"name"`
	State               string                 `pulumi:"state"`
	TenantId            string                 `pulumi:"tenantId"`
	Type                string                 `pulumi:"type"`
}
