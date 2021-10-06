


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMSTIDataConnector(ctx *pulumi.Context, args *LookupMSTIDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMSTIDataConnectorResult, error) {
	var rv LookupMSTIDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getMSTIDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMSTIDataConnectorArgs struct {
	DataConnectorId                     string `pulumi:"dataConnectorId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupMSTIDataConnectorResult struct {
	DataTypes MSTIDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag      *string                            `pulumi:"etag"`
	Id        string                             `pulumi:"id"`
	Kind      string                             `pulumi:"kind"`
	Name      string                             `pulumi:"name"`
	TenantId  string                             `pulumi:"tenantId"`
	Type      string                             `pulumi:"type"`
}
