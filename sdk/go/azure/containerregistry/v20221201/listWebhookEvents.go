


package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebhookEvents(ctx *pulumi.Context, args *ListWebhookEventsArgs, opts ...pulumi.InvokeOption) (*ListWebhookEventsResult, error) {
	var rv ListWebhookEventsResult
	err := ctx.Invoke("azure-native:containerregistry/v20221201:listWebhookEvents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebhookEventsArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WebhookName       string `pulumi:"webhookName"`
}


type ListWebhookEventsResult struct {
	NextLink *string         `pulumi:"nextLink"`
	Value    []EventResponse `pulumi:"value"`
}

func ListWebhookEventsOutput(ctx *pulumi.Context, args ListWebhookEventsOutputArgs, opts ...pulumi.InvokeOption) ListWebhookEventsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebhookEventsResult, error) {
			args := v.(ListWebhookEventsArgs)
			r, err := ListWebhookEvents(ctx, &args, opts...)
			var s ListWebhookEventsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebhookEventsResultOutput)
}

type ListWebhookEventsOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WebhookName       pulumi.StringInput `pulumi:"webhookName"`
}

func (ListWebhookEventsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebhookEventsArgs)(nil)).Elem()
}


type ListWebhookEventsResultOutput struct{ *pulumi.OutputState }

func (ListWebhookEventsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebhookEventsResult)(nil)).Elem()
}

func (o ListWebhookEventsResultOutput) ToListWebhookEventsResultOutput() ListWebhookEventsResultOutput {
	return o
}

func (o ListWebhookEventsResultOutput) ToListWebhookEventsResultOutputWithContext(ctx context.Context) ListWebhookEventsResultOutput {
	return o
}

func (o ListWebhookEventsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebhookEventsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListWebhookEventsResultOutput) Value() EventResponseArrayOutput {
	return o.ApplyT(func(v ListWebhookEventsResult) []EventResponse { return v.Value }).(EventResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebhookEventsResultOutput{})
}
