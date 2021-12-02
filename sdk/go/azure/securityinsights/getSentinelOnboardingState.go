


package securityinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSentinelOnboardingState(ctx *pulumi.Context, args *LookupSentinelOnboardingStateArgs, opts ...pulumi.InvokeOption) (*LookupSentinelOnboardingStateResult, error) {
	var rv LookupSentinelOnboardingStateResult
	err := ctx.Invoke("azure-native:securityinsights:getSentinelOnboardingState", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSentinelOnboardingStateArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	SentinelOnboardingStateName         string `pulumi:"sentinelOnboardingStateName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupSentinelOnboardingStateResult struct {
	CustomerManagedKey *bool              `pulumi:"customerManagedKey"`
	Etag               *string            `pulumi:"etag"`
	Id                 string             `pulumi:"id"`
	Name               string             `pulumi:"name"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
}
