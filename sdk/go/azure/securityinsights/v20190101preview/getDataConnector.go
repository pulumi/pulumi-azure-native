


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDataConnector(ctx *pulumi.Context, args *LookupDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupDataConnectorResult, error) {
	var rv LookupDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataConnectorArgs struct {
	DataConnectorId                     string `pulumi:"dataConnectorId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupDataConnectorResult struct {
	Etag *string `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Kind string  `pulumi:"kind"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}
