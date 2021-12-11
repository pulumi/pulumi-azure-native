


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMDATPDataConnector(ctx *pulumi.Context, args *LookupMDATPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMDATPDataConnectorResult, error) {
	var rv LookupMDATPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getMDATPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMDATPDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMDATPDataConnectorResult struct {
	DataTypes  *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag       *string                                `pulumi:"etag"`
	Id         string                                 `pulumi:"id"`
	Kind       string                                 `pulumi:"kind"`
	Name       string                                 `pulumi:"name"`
	SystemData SystemDataResponse                     `pulumi:"systemData"`
	TenantId   string                                 `pulumi:"tenantId"`
	Type       string                                 `pulumi:"type"`
}
