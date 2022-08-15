


package v20140401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupAutoscaleSetting(ctx *pulumi.Context, args *LookupAutoscaleSettingArgs, opts ...pulumi.InvokeOption) (*LookupAutoscaleSettingResult, error) {
	var rv LookupAutoscaleSettingResult
	err := ctx.Invoke("azure-native:insights/v20140401:getAutoscaleSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAutoscaleSettingArgs struct {
	AutoscaleSettingName string `pulumi:"autoscaleSettingName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupAutoscaleSettingResult struct {
	Enabled                *bool                           `pulumi:"enabled"`
	Id                     string                          `pulumi:"id"`
	Location               string                          `pulumi:"location"`
	Name                   string                          `pulumi:"name"`
	Notifications          []AutoscaleNotificationResponse `pulumi:"notifications"`
	Profiles               []AutoscaleProfileResponse      `pulumi:"profiles"`
	Tags                   map[string]string               `pulumi:"tags"`
	TargetResourceLocation *string                         `pulumi:"targetResourceLocation"`
	TargetResourceUri      *string                         `pulumi:"targetResourceUri"`
	Type                   string                          `pulumi:"type"`
}


func (val *LookupAutoscaleSettingResult) Defaults() *LookupAutoscaleSettingResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Enabled) {
		enabled_ := false
		tmp.Enabled = &enabled_
	}
	return &tmp
}

func LookupAutoscaleSettingOutput(ctx *pulumi.Context, args LookupAutoscaleSettingOutputArgs, opts ...pulumi.InvokeOption) LookupAutoscaleSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutoscaleSettingResult, error) {
			args := v.(LookupAutoscaleSettingArgs)
			r, err := LookupAutoscaleSetting(ctx, &args, opts...)
			var s LookupAutoscaleSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutoscaleSettingResultOutput)
}

type LookupAutoscaleSettingOutputArgs struct {
	AutoscaleSettingName pulumi.StringInput `pulumi:"autoscaleSettingName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAutoscaleSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoscaleSettingArgs)(nil)).Elem()
}


type LookupAutoscaleSettingResultOutput struct{ *pulumi.OutputState }

func (LookupAutoscaleSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoscaleSettingResult)(nil)).Elem()
}

func (o LookupAutoscaleSettingResultOutput) ToLookupAutoscaleSettingResultOutput() LookupAutoscaleSettingResultOutput {
	return o
}

func (o LookupAutoscaleSettingResultOutput) ToLookupAutoscaleSettingResultOutputWithContext(ctx context.Context) LookupAutoscaleSettingResultOutput {
	return o
}

func (o LookupAutoscaleSettingResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAutoscaleSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAutoscaleSettingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAutoscaleSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAutoscaleSettingResultOutput) Notifications() AutoscaleNotificationResponseArrayOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) []AutoscaleNotificationResponse { return v.Notifications }).(AutoscaleNotificationResponseArrayOutput)
}

func (o LookupAutoscaleSettingResultOutput) Profiles() AutoscaleProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) []AutoscaleProfileResponse { return v.Profiles }).(AutoscaleProfileResponseArrayOutput)
}

func (o LookupAutoscaleSettingResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAutoscaleSettingResultOutput) TargetResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) *string { return v.TargetResourceLocation }).(pulumi.StringPtrOutput)
}

func (o LookupAutoscaleSettingResultOutput) TargetResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) *string { return v.TargetResourceUri }).(pulumi.StringPtrOutput)
}

func (o LookupAutoscaleSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutoscaleSettingResultOutput{})
}
