


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMTPDataConnector(ctx *pulumi.Context, args *LookupMTPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMTPDataConnectorResult, error) {
	var rv LookupMTPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getMTPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMTPDataConnectorArgs struct {
	DataConnectorId                     string `pulumi:"dataConnectorId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupMTPDataConnectorResult struct {
	DataTypes  MTPDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                           `pulumi:"etag"`
	Id         string                            `pulumi:"id"`
	Kind       string                            `pulumi:"kind"`
	Name       string                            `pulumi:"name"`
	SystemData SystemDataResponse                `pulumi:"systemData"`
	TenantId   string                            `pulumi:"tenantId"`
	Type       string                            `pulumi:"type"`
}
