


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMCASDataConnector(ctx *pulumi.Context, args *LookupMCASDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMCASDataConnectorResult, error) {
	var rv LookupMCASDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getMCASDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMCASDataConnectorArgs struct {
	DataConnectorId                     string `pulumi:"dataConnectorId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupMCASDataConnectorResult struct {
	DataTypes MCASDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag      *string                            `pulumi:"etag"`
	Id        string                             `pulumi:"id"`
	Kind      string                             `pulumi:"kind"`
	Name      string                             `pulumi:"name"`
	TenantId  string                             `pulumi:"tenantId"`
	Type      string                             `pulumi:"type"`
}
