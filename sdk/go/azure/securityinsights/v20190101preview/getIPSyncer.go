


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIPSyncer(ctx *pulumi.Context, args *LookupIPSyncerArgs, opts ...pulumi.InvokeOption) (*LookupIPSyncerResult, error) {
	var rv LookupIPSyncerResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getIPSyncer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIPSyncerArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	SettingsName                        string `pulumi:"settingsName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupIPSyncerResult struct {
	Etag      *string `pulumi:"etag"`
	Id        string  `pulumi:"id"`
	IsEnabled bool    `pulumi:"isEnabled"`
	Kind      string  `pulumi:"kind"`
	Name      string  `pulumi:"name"`
	Type      string  `pulumi:"type"`
}
