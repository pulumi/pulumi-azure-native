


package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBandwidthSchedule(ctx *pulumi.Context, args *LookupBandwidthScheduleArgs, opts ...pulumi.InvokeOption) (*LookupBandwidthScheduleResult, error) {
	var rv LookupBandwidthScheduleResult
	err := ctx.Invoke("azure-native:databoxedge/v20230101preview:getBandwidthSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBandwidthScheduleArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBandwidthScheduleResult struct {
	Days       []string           `pulumi:"days"`
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	RateInMbps int                `pulumi:"rateInMbps"`
	Start      string             `pulumi:"start"`
	Stop       string             `pulumi:"stop"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupBandwidthScheduleOutput(ctx *pulumi.Context, args LookupBandwidthScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupBandwidthScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBandwidthScheduleResult, error) {
			args := v.(LookupBandwidthScheduleArgs)
			r, err := LookupBandwidthSchedule(ctx, &args, opts...)
			var s LookupBandwidthScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBandwidthScheduleResultOutput)
}

type LookupBandwidthScheduleOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBandwidthScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBandwidthScheduleArgs)(nil)).Elem()
}


type LookupBandwidthScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupBandwidthScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBandwidthScheduleResult)(nil)).Elem()
}

func (o LookupBandwidthScheduleResultOutput) ToLookupBandwidthScheduleResultOutput() LookupBandwidthScheduleResultOutput {
	return o
}

func (o LookupBandwidthScheduleResultOutput) ToLookupBandwidthScheduleResultOutputWithContext(ctx context.Context) LookupBandwidthScheduleResultOutput {
	return o
}

func (o LookupBandwidthScheduleResultOutput) Days() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBandwidthScheduleResult) []string { return v.Days }).(pulumi.StringArrayOutput)
}

func (o LookupBandwidthScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBandwidthScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBandwidthScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBandwidthScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBandwidthScheduleResultOutput) RateInMbps() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBandwidthScheduleResult) int { return v.RateInMbps }).(pulumi.IntOutput)
}

func (o LookupBandwidthScheduleResultOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBandwidthScheduleResult) string { return v.Start }).(pulumi.StringOutput)
}

func (o LookupBandwidthScheduleResultOutput) Stop() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBandwidthScheduleResult) string { return v.Stop }).(pulumi.StringOutput)
}

func (o LookupBandwidthScheduleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBandwidthScheduleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBandwidthScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBandwidthScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBandwidthScheduleResultOutput{})
}
