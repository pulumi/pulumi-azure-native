


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupASCDataConnector(ctx *pulumi.Context, args *LookupASCDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupASCDataConnectorResult, error) {
	var rv LookupASCDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getASCDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupASCDataConnectorArgs struct {
	DataConnectorId                     string `pulumi:"dataConnectorId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupASCDataConnectorResult struct {
	DataTypes      *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag           *string                                `pulumi:"etag"`
	Id             string                                 `pulumi:"id"`
	Kind           string                                 `pulumi:"kind"`
	Name           string                                 `pulumi:"name"`
	SubscriptionId *string                                `pulumi:"subscriptionId"`
	Type           string                                 `pulumi:"type"`
}
