


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledSynchronizationSetting(ctx *pulumi.Context, args *LookupScheduledSynchronizationSettingArgs, opts ...pulumi.InvokeOption) (*LookupScheduledSynchronizationSettingResult, error) {
	var rv LookupScheduledSynchronizationSettingResult
	err := ctx.Invoke("azure-native:datashare/v20181101preview:getScheduledSynchronizationSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledSynchronizationSettingArgs struct {
	AccountName                string `pulumi:"accountName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	ShareName                  string `pulumi:"shareName"`
	SynchronizationSettingName string `pulumi:"synchronizationSettingName"`
}


type LookupScheduledSynchronizationSettingResult struct {
	CreatedAt           string `pulumi:"createdAt"`
	Id                  string `pulumi:"id"`
	Kind                string `pulumi:"kind"`
	Name                string `pulumi:"name"`
	ProvisioningState   string `pulumi:"provisioningState"`
	RecurrenceInterval  string `pulumi:"recurrenceInterval"`
	SynchronizationTime string `pulumi:"synchronizationTime"`
	Type                string `pulumi:"type"`
	UserName            string `pulumi:"userName"`
}

func LookupScheduledSynchronizationSettingOutput(ctx *pulumi.Context, args LookupScheduledSynchronizationSettingOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledSynchronizationSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledSynchronizationSettingResult, error) {
			args := v.(LookupScheduledSynchronizationSettingArgs)
			r, err := LookupScheduledSynchronizationSetting(ctx, &args, opts...)
			var s LookupScheduledSynchronizationSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduledSynchronizationSettingResultOutput)
}

type LookupScheduledSynchronizationSettingOutputArgs struct {
	AccountName                pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName                  pulumi.StringInput `pulumi:"shareName"`
	SynchronizationSettingName pulumi.StringInput `pulumi:"synchronizationSettingName"`
}

func (LookupScheduledSynchronizationSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledSynchronizationSettingArgs)(nil)).Elem()
}


type LookupScheduledSynchronizationSettingResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledSynchronizationSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledSynchronizationSettingResult)(nil)).Elem()
}

func (o LookupScheduledSynchronizationSettingResultOutput) ToLookupScheduledSynchronizationSettingResultOutput() LookupScheduledSynchronizationSettingResultOutput {
	return o
}

func (o LookupScheduledSynchronizationSettingResultOutput) ToLookupScheduledSynchronizationSettingResultOutputWithContext(ctx context.Context) LookupScheduledSynchronizationSettingResultOutput {
	return o
}

func (o LookupScheduledSynchronizationSettingResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledSynchronizationSettingResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupScheduledSynchronizationSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledSynchronizationSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScheduledSynchronizationSettingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledSynchronizationSettingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupScheduledSynchronizationSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledSynchronizationSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScheduledSynchronizationSettingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledSynchronizationSettingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupScheduledSynchronizationSettingResultOutput) RecurrenceInterval() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledSynchronizationSettingResult) string { return v.RecurrenceInterval }).(pulumi.StringOutput)
}

func (o LookupScheduledSynchronizationSettingResultOutput) SynchronizationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledSynchronizationSettingResult) string { return v.SynchronizationTime }).(pulumi.StringOutput)
}

func (o LookupScheduledSynchronizationSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledSynchronizationSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupScheduledSynchronizationSettingResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledSynchronizationSettingResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledSynchronizationSettingResultOutput{})
}
