


package securityinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEntityAnalytics(ctx *pulumi.Context, args *LookupEntityAnalyticsArgs, opts ...pulumi.InvokeOption) (*LookupEntityAnalyticsResult, error) {
	var rv LookupEntityAnalyticsResult
	err := ctx.Invoke("azure-native:securityinsights:getEntityAnalytics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEntityAnalyticsArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	SettingsName                        string `pulumi:"settingsName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupEntityAnalyticsResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	IsEnabled  bool               `pulumi:"isEnabled"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}
