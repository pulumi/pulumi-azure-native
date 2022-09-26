


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSynchronizationSetting(ctx *pulumi.Context, args *LookupSynchronizationSettingArgs, opts ...pulumi.InvokeOption) (*LookupSynchronizationSettingResult, error) {
	var rv LookupSynchronizationSettingResult
	err := ctx.Invoke("azure-native:datashare/v20181101preview:getSynchronizationSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSynchronizationSettingArgs struct {
	AccountName                string `pulumi:"accountName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	ShareName                  string `pulumi:"shareName"`
	SynchronizationSettingName string `pulumi:"synchronizationSettingName"`
}


type LookupSynchronizationSettingResult struct {
	Id   string `pulumi:"id"`
	Kind string `pulumi:"kind"`
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

func LookupSynchronizationSettingOutput(ctx *pulumi.Context, args LookupSynchronizationSettingOutputArgs, opts ...pulumi.InvokeOption) LookupSynchronizationSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSynchronizationSettingResult, error) {
			args := v.(LookupSynchronizationSettingArgs)
			r, err := LookupSynchronizationSetting(ctx, &args, opts...)
			var s LookupSynchronizationSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSynchronizationSettingResultOutput)
}

type LookupSynchronizationSettingOutputArgs struct {
	AccountName                pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName                  pulumi.StringInput `pulumi:"shareName"`
	SynchronizationSettingName pulumi.StringInput `pulumi:"synchronizationSettingName"`
}

func (LookupSynchronizationSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSynchronizationSettingArgs)(nil)).Elem()
}


type LookupSynchronizationSettingResultOutput struct{ *pulumi.OutputState }

func (LookupSynchronizationSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSynchronizationSettingResult)(nil)).Elem()
}

func (o LookupSynchronizationSettingResultOutput) ToLookupSynchronizationSettingResultOutput() LookupSynchronizationSettingResultOutput {
	return o
}

func (o LookupSynchronizationSettingResultOutput) ToLookupSynchronizationSettingResultOutputWithContext(ctx context.Context) LookupSynchronizationSettingResultOutput {
	return o
}

func (o LookupSynchronizationSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynchronizationSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSynchronizationSettingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynchronizationSettingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupSynchronizationSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynchronizationSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSynchronizationSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynchronizationSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSynchronizationSettingResultOutput{})
}
