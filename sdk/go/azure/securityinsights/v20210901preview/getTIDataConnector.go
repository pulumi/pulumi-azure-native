


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTIDataConnector(ctx *pulumi.Context, args *LookupTIDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupTIDataConnectorResult, error) {
	var rv LookupTIDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getTIDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTIDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupTIDataConnectorResult struct {
	DataTypes         TIDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag              *string                          `pulumi:"etag"`
	Id                string                           `pulumi:"id"`
	Kind              string                           `pulumi:"kind"`
	Name              string                           `pulumi:"name"`
	SystemData        SystemDataResponse               `pulumi:"systemData"`
	TenantId          string                           `pulumi:"tenantId"`
	TipLookbackPeriod *string                          `pulumi:"tipLookbackPeriod"`
	Type              string                           `pulumi:"type"`
}
