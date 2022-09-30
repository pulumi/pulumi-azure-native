


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListShareSubscriptionSourceShareSynchronizationSettings(ctx *pulumi.Context, args *ListShareSubscriptionSourceShareSynchronizationSettingsArgs, opts ...pulumi.InvokeOption) (*ListShareSubscriptionSourceShareSynchronizationSettingsResult, error) {
	var rv ListShareSubscriptionSourceShareSynchronizationSettingsResult
	err := ctx.Invoke("azure-native:datashare/v20200901:listShareSubscriptionSourceShareSynchronizationSettings", args, &rv, opts...)
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

func ListShareSubscriptionSourceShareSynchronizationSettingsOutput(ctx *pulumi.Context, args ListShareSubscriptionSourceShareSynchronizationSettingsOutputArgs, opts ...pulumi.InvokeOption) ListShareSubscriptionSourceShareSynchronizationSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListShareSubscriptionSourceShareSynchronizationSettingsResult, error) {
			args := v.(ListShareSubscriptionSourceShareSynchronizationSettingsArgs)
			r, err := ListShareSubscriptionSourceShareSynchronizationSettings(ctx, &args, opts...)
			var s ListShareSubscriptionSourceShareSynchronizationSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListShareSubscriptionSourceShareSynchronizationSettingsResultOutput)
}

type ListShareSubscriptionSourceShareSynchronizationSettingsOutputArgs struct {
	AccountName           pulumi.StringInput    `pulumi:"accountName"`
	ResourceGroupName     pulumi.StringInput    `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput    `pulumi:"shareSubscriptionName"`
	SkipToken             pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListShareSubscriptionSourceShareSynchronizationSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSubscriptionSourceShareSynchronizationSettingsArgs)(nil)).Elem()
}


type ListShareSubscriptionSourceShareSynchronizationSettingsResultOutput struct{ *pulumi.OutputState }

func (ListShareSubscriptionSourceShareSynchronizationSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListShareSubscriptionSourceShareSynchronizationSettingsResult)(nil)).Elem()
}

func (o ListShareSubscriptionSourceShareSynchronizationSettingsResultOutput) ToListShareSubscriptionSourceShareSynchronizationSettingsResultOutput() ListShareSubscriptionSourceShareSynchronizationSettingsResultOutput {
	return o
}

func (o ListShareSubscriptionSourceShareSynchronizationSettingsResultOutput) ToListShareSubscriptionSourceShareSynchronizationSettingsResultOutputWithContext(ctx context.Context) ListShareSubscriptionSourceShareSynchronizationSettingsResultOutput {
	return o
}

func (o ListShareSubscriptionSourceShareSynchronizationSettingsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListShareSubscriptionSourceShareSynchronizationSettingsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListShareSubscriptionSourceShareSynchronizationSettingsResultOutput) Value() ScheduledSourceSynchronizationSettingResponseArrayOutput {
	return o.ApplyT(func(v ListShareSubscriptionSourceShareSynchronizationSettingsResult) []ScheduledSourceSynchronizationSettingResponse {
		return v.Value
	}).(ScheduledSourceSynchronizationSettingResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListShareSubscriptionSourceShareSynchronizationSettingsResultOutput{})
}
