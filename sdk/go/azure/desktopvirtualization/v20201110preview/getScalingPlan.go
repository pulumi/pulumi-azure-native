


package v20201110preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupScalingPlan(ctx *pulumi.Context, args *LookupScalingPlanArgs, opts ...pulumi.InvokeOption) (*LookupScalingPlanResult, error) {
	var rv LookupScalingPlanResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20201110preview:getScalingPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScalingPlanArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScalingPlanName   string `pulumi:"scalingPlanName"`
}


type LookupScalingPlanResult struct {
	Description        *string                            `pulumi:"description"`
	ExclusionTag       *string                            `pulumi:"exclusionTag"`
	FriendlyName       *string                            `pulumi:"friendlyName"`
	HostPoolReferences []ScalingHostPoolReferenceResponse `pulumi:"hostPoolReferences"`
	HostPoolType       *string                            `pulumi:"hostPoolType"`
	Id                 string                             `pulumi:"id"`
	Location           string                             `pulumi:"location"`
	Name               string                             `pulumi:"name"`
	Schedules          []ScalingScheduleResponse          `pulumi:"schedules"`
	Tags               map[string]string                  `pulumi:"tags"`
	TimeZone           *string                            `pulumi:"timeZone"`
	Type               string                             `pulumi:"type"`
}

func LookupScalingPlanOutput(ctx *pulumi.Context, args LookupScalingPlanOutputArgs, opts ...pulumi.InvokeOption) LookupScalingPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScalingPlanResult, error) {
			args := v.(LookupScalingPlanArgs)
			r, err := LookupScalingPlan(ctx, &args, opts...)
			var s LookupScalingPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScalingPlanResultOutput)
}

type LookupScalingPlanOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ScalingPlanName   pulumi.StringInput `pulumi:"scalingPlanName"`
}

func (LookupScalingPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPlanArgs)(nil)).Elem()
}


type LookupScalingPlanResultOutput struct{ *pulumi.OutputState }

func (LookupScalingPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPlanResult)(nil)).Elem()
}

func (o LookupScalingPlanResultOutput) ToLookupScalingPlanResultOutput() LookupScalingPlanResultOutput {
	return o
}

func (o LookupScalingPlanResultOutput) ToLookupScalingPlanResultOutputWithContext(ctx context.Context) LookupScalingPlanResultOutput {
	return o
}

func (o LookupScalingPlanResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanResultOutput) ExclusionTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.ExclusionTag }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanResultOutput) HostPoolReferences() ScalingHostPoolReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) []ScalingHostPoolReferenceResponse { return v.HostPoolReferences }).(ScalingHostPoolReferenceResponseArrayOutput)
}

func (o LookupScalingPlanResultOutput) HostPoolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.HostPoolType }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScalingPlanResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupScalingPlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScalingPlanResultOutput) Schedules() ScalingScheduleResponseArrayOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) []ScalingScheduleResponse { return v.Schedules }).(ScalingScheduleResponseArrayOutput)
}

func (o LookupScalingPlanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupScalingPlanResultOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScalingPlanResultOutput{})
}
