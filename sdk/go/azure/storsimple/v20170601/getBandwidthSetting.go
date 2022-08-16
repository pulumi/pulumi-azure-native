


package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBandwidthSetting(ctx *pulumi.Context, args *LookupBandwidthSettingArgs, opts ...pulumi.InvokeOption) (*LookupBandwidthSettingResult, error) {
	var rv LookupBandwidthSettingResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:getBandwidthSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBandwidthSettingArgs struct {
	BandwidthSettingName string `pulumi:"bandwidthSettingName"`
	ManagerName          string `pulumi:"managerName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupBandwidthSettingResult struct {
	Id          string                      `pulumi:"id"`
	Kind        *string                     `pulumi:"kind"`
	Name        string                      `pulumi:"name"`
	Schedules   []BandwidthScheduleResponse `pulumi:"schedules"`
	Type        string                      `pulumi:"type"`
	VolumeCount int                         `pulumi:"volumeCount"`
}

func LookupBandwidthSettingOutput(ctx *pulumi.Context, args LookupBandwidthSettingOutputArgs, opts ...pulumi.InvokeOption) LookupBandwidthSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBandwidthSettingResult, error) {
			args := v.(LookupBandwidthSettingArgs)
			r, err := LookupBandwidthSetting(ctx, &args, opts...)
			var s LookupBandwidthSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBandwidthSettingResultOutput)
}

type LookupBandwidthSettingOutputArgs struct {
	BandwidthSettingName pulumi.StringInput `pulumi:"bandwidthSettingName"`
	ManagerName          pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBandwidthSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBandwidthSettingArgs)(nil)).Elem()
}


type LookupBandwidthSettingResultOutput struct{ *pulumi.OutputState }

func (LookupBandwidthSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBandwidthSettingResult)(nil)).Elem()
}

func (o LookupBandwidthSettingResultOutput) ToLookupBandwidthSettingResultOutput() LookupBandwidthSettingResultOutput {
	return o
}

func (o LookupBandwidthSettingResultOutput) ToLookupBandwidthSettingResultOutputWithContext(ctx context.Context) LookupBandwidthSettingResultOutput {
	return o
}

func (o LookupBandwidthSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBandwidthSettingResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupBandwidthSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBandwidthSettingResultOutput) Schedules() BandwidthScheduleResponseArrayOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) []BandwidthScheduleResponse { return v.Schedules }).(BandwidthScheduleResponseArrayOutput)
}

func (o LookupBandwidthSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBandwidthSettingResultOutput) VolumeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) int { return v.VolumeCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBandwidthSettingResultOutput{})
}
