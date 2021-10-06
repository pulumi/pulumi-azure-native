


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOfficeATPDataConnector(ctx *pulumi.Context, args *LookupOfficeATPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupOfficeATPDataConnectorResult, error) {
	var rv LookupOfficeATPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getOfficeATPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOfficeATPDataConnectorArgs struct {
	DataConnectorId                     string `pulumi:"dataConnectorId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupOfficeATPDataConnectorResult struct {
	DataTypes *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag      *string                                `pulumi:"etag"`
	Id        string                                 `pulumi:"id"`
	Kind      string                                 `pulumi:"kind"`
	Name      string                                 `pulumi:"name"`
	TenantId  string                                 `pulumi:"tenantId"`
	Type      string                                 `pulumi:"type"`
}
