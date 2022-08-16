


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledTrigger(ctx *pulumi.Context, args *LookupScheduledTriggerArgs, opts ...pulumi.InvokeOption) (*LookupScheduledTriggerResult, error) {
	var rv LookupScheduledTriggerResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getScheduledTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledTriggerArgs struct {
	AccountName           string `pulumi:"accountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
	TriggerName           string `pulumi:"triggerName"`
}


type LookupScheduledTriggerResult struct {
	CreatedAt           string  `pulumi:"createdAt"`
	Id                  string  `pulumi:"id"`
	Kind                string  `pulumi:"kind"`
	Name                string  `pulumi:"name"`
	ProvisioningState   string  `pulumi:"provisioningState"`
	RecurrenceInterval  string  `pulumi:"recurrenceInterval"`
	SynchronizationMode *string `pulumi:"synchronizationMode"`
	SynchronizationTime string  `pulumi:"synchronizationTime"`
	TriggerStatus       string  `pulumi:"triggerStatus"`
	Type                string  `pulumi:"type"`
	UserName            string  `pulumi:"userName"`
}

func LookupScheduledTriggerOutput(ctx *pulumi.Context, args LookupScheduledTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledTriggerResult, error) {
			args := v.(LookupScheduledTriggerArgs)
			r, err := LookupScheduledTrigger(ctx, &args, opts...)
			var s LookupScheduledTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduledTriggerResultOutput)
}

type LookupScheduledTriggerOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
	TriggerName           pulumi.StringInput `pulumi:"triggerName"`
}

func (LookupScheduledTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledTriggerArgs)(nil)).Elem()
}


type LookupScheduledTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledTriggerResult)(nil)).Elem()
}

func (o LookupScheduledTriggerResultOutput) ToLookupScheduledTriggerResultOutput() LookupScheduledTriggerResultOutput {
	return o
}

func (o LookupScheduledTriggerResultOutput) ToLookupScheduledTriggerResultOutputWithContext(ctx context.Context) LookupScheduledTriggerResultOutput {
	return o
}

func (o LookupScheduledTriggerResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupScheduledTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScheduledTriggerResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupScheduledTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScheduledTriggerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupScheduledTriggerResultOutput) RecurrenceInterval() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.RecurrenceInterval }).(pulumi.StringOutput)
}

func (o LookupScheduledTriggerResultOutput) SynchronizationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) *string { return v.SynchronizationMode }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledTriggerResultOutput) SynchronizationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.SynchronizationTime }).(pulumi.StringOutput)
}

func (o LookupScheduledTriggerResultOutput) TriggerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.TriggerStatus }).(pulumi.StringOutput)
}

func (o LookupScheduledTriggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupScheduledTriggerResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledTriggerResultOutput{})
}
