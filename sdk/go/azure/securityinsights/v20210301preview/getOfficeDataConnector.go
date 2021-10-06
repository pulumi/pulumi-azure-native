


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOfficeDataConnector(ctx *pulumi.Context, args *LookupOfficeDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupOfficeDataConnectorResult, error) {
	var rv LookupOfficeDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getOfficeDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOfficeDataConnectorArgs struct {
	DataConnectorId                     string `pulumi:"dataConnectorId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupOfficeDataConnectorResult struct {
	DataTypes  OfficeDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                              `pulumi:"etag"`
	Id         string                               `pulumi:"id"`
	Kind       string                               `pulumi:"kind"`
	Name       string                               `pulumi:"name"`
	SystemData SystemDataResponse                   `pulumi:"systemData"`
	TenantId   string                               `pulumi:"tenantId"`
	Type       string                               `pulumi:"type"`
}
