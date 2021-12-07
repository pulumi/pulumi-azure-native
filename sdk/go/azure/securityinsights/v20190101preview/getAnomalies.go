


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAnomalies(ctx *pulumi.Context, args *LookupAnomaliesArgs, opts ...pulumi.InvokeOption) (*LookupAnomaliesResult, error) {
	var rv LookupAnomaliesResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getAnomalies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnomaliesArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	SettingsName                        string `pulumi:"settingsName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupAnomaliesResult struct {
	Etag      *string `pulumi:"etag"`
	Id        string  `pulumi:"id"`
	IsEnabled bool    `pulumi:"isEnabled"`
	Kind      string  `pulumi:"kind"`
	Name      string  `pulumi:"name"`
	Type      string  `pulumi:"type"`
}
