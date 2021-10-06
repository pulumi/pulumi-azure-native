


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListShareSubscriptionSourceShareSynchronizationSettings(ctx *pulumi.Context, args *ListShareSubscriptionSourceShareSynchronizationSettingsArgs, opts ...pulumi.InvokeOption) (*ListShareSubscriptionSourceShareSynchronizationSettingsResult, error) {
	var rv ListShareSubscriptionSourceShareSynchronizationSettingsResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:listShareSubscriptionSourceShareSynchronizationSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListShareSubscriptionSourceShareSynchronizationSettingsArgs struct {
	AccountName           string  `pulumi:"accountName"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	SkipToken             *string `pulumi:"skipToken"`
}


type ListShareSubscriptionSourceShareSynchronizationSettingsResult struct {
	NextLink *string                                         `pulumi:"nextLink"`
	Value    []ScheduledSourceSynchronizationSettingResponse `pulumi:"value"`
}
