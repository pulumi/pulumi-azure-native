


package v20200101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAATPDataConnector(ctx *pulumi.Context, args *LookupAATPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAATPDataConnectorResult, error) {
	var rv LookupAATPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20200101:getAATPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAATPDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAATPDataConnectorResult struct {
	DataTypes *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag      *string                                `pulumi:"etag"`
	Id        string                                 `pulumi:"id"`
	Kind      string                                 `pulumi:"kind"`
	Name      string                                 `pulumi:"name"`
	TenantId  *string                                `pulumi:"tenantId"`
	Type      string                                 `pulumi:"type"`
}
