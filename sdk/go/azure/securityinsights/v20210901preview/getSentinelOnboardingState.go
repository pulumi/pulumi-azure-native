


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSentinelOnboardingState(ctx *pulumi.Context, args *LookupSentinelOnboardingStateArgs, opts ...pulumi.InvokeOption) (*LookupSentinelOnboardingStateResult, error) {
	var rv LookupSentinelOnboardingStateResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getSentinelOnboardingState", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSentinelOnboardingStateArgs struct {
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	SentinelOnboardingStateName string `pulumi:"sentinelOnboardingStateName"`
	WorkspaceName               string `pulumi:"workspaceName"`
}


type LookupSentinelOnboardingStateResult struct {
	CustomerManagedKey *bool              `pulumi:"customerManagedKey"`
	Etag               *string            `pulumi:"etag"`
	Id                 string             `pulumi:"id"`
	Name               string             `pulumi:"name"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
}
