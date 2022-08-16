


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListServiceFabricApplicableSchedules(ctx *pulumi.Context, args *ListServiceFabricApplicableSchedulesArgs, opts ...pulumi.InvokeOption) (*ListServiceFabricApplicableSchedulesResult, error) {
	var rv ListServiceFabricApplicableSchedulesResult
	err := ctx.Invoke("azure-native:devtestlab:listServiceFabricApplicableSchedules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListServiceFabricApplicableSchedulesArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	UserName          string `pulumi:"userName"`
}


type ListServiceFabricApplicableSchedulesResult struct {
	Id             string            `pulumi:"id"`
	LabVmsShutdown *ScheduleResponse `pulumi:"labVmsShutdown"`
	LabVmsStartup  *ScheduleResponse `pulumi:"labVmsStartup"`
	Location       *string           `pulumi:"location"`
	Name           string            `pulumi:"name"`
	Tags           map[string]string `pulumi:"tags"`
	Type           string            `pulumi:"type"`
}


func (val *ListServiceFabricApplicableSchedulesResult) Defaults() *ListServiceFabricApplicableSchedulesResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.LabVmsShutdown = tmp.LabVmsShutdown.Defaults()

	tmp.LabVmsStartup = tmp.LabVmsStartup.Defaults()

	return &tmp
}

func ListServiceFabricApplicableSchedulesOutput(ctx *pulumi.Context, args ListServiceFabricApplicableSchedulesOutputArgs, opts ...pulumi.InvokeOption) ListServiceFabricApplicableSchedulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListServiceFabricApplicableSchedulesResult, error) {
			args := v.(ListServiceFabricApplicableSchedulesArgs)
			r, err := ListServiceFabricApplicableSchedules(ctx, &args, opts...)
			var s ListServiceFabricApplicableSchedulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListServiceFabricApplicableSchedulesResultOutput)
}

type ListServiceFabricApplicableSchedulesOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	UserName          pulumi.StringInput `pulumi:"userName"`
}

func (ListServiceFabricApplicableSchedulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServiceFabricApplicableSchedulesArgs)(nil)).Elem()
}


type ListServiceFabricApplicableSchedulesResultOutput struct{ *pulumi.OutputState }

func (ListServiceFabricApplicableSchedulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServiceFabricApplicableSchedulesResult)(nil)).Elem()
}

func (o ListServiceFabricApplicableSchedulesResultOutput) ToListServiceFabricApplicableSchedulesResultOutput() ListServiceFabricApplicableSchedulesResultOutput {
	return o
}

func (o ListServiceFabricApplicableSchedulesResultOutput) ToListServiceFabricApplicableSchedulesResultOutputWithContext(ctx context.Context) ListServiceFabricApplicableSchedulesResultOutput {
	return o
}

func (o ListServiceFabricApplicableSchedulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListServiceFabricApplicableSchedulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListServiceFabricApplicableSchedulesResultOutput) LabVmsShutdown() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListServiceFabricApplicableSchedulesResult) *ScheduleResponse { return v.LabVmsShutdown }).(ScheduleResponsePtrOutput)
}

func (o ListServiceFabricApplicableSchedulesResultOutput) LabVmsStartup() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListServiceFabricApplicableSchedulesResult) *ScheduleResponse { return v.LabVmsStartup }).(ScheduleResponsePtrOutput)
}

func (o ListServiceFabricApplicableSchedulesResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListServiceFabricApplicableSchedulesResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ListServiceFabricApplicableSchedulesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListServiceFabricApplicableSchedulesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListServiceFabricApplicableSchedulesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListServiceFabricApplicableSchedulesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListServiceFabricApplicableSchedulesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListServiceFabricApplicableSchedulesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServiceFabricApplicableSchedulesResultOutput{})
}
