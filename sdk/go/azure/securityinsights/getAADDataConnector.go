


package securityinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAADDataConnector(ctx *pulumi.Context, args *LookupAADDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAADDataConnectorResult, error) {
	var rv LookupAADDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights:getAADDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAADDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAADDataConnectorResult struct {
	DataTypes *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag      *string                                `pulumi:"etag"`
	Id        string                                 `pulumi:"id"`
	Kind      string                                 `pulumi:"kind"`
	Name      string                                 `pulumi:"name"`
	TenantId  *string                                `pulumi:"tenantId"`
	Type      string                                 `pulumi:"type"`
}
