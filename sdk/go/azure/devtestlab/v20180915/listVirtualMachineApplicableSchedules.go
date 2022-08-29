


package v20180915

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVirtualMachineApplicableSchedules(ctx *pulumi.Context, args *ListVirtualMachineApplicableSchedulesArgs, opts ...pulumi.InvokeOption) (*ListVirtualMachineApplicableSchedulesResult, error) {
	var rv ListVirtualMachineApplicableSchedulesResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:listVirtualMachineApplicableSchedules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListVirtualMachineApplicableSchedulesArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListVirtualMachineApplicableSchedulesResult struct {
	Id             string            `pulumi:"id"`
	LabVmsShutdown *ScheduleResponse `pulumi:"labVmsShutdown"`
	LabVmsStartup  *ScheduleResponse `pulumi:"labVmsStartup"`
	Location       *string           `pulumi:"location"`
	Name           string            `pulumi:"name"`
	Tags           map[string]string `pulumi:"tags"`
	Type           string            `pulumi:"type"`
}


func (val *ListVirtualMachineApplicableSchedulesResult) Defaults() *ListVirtualMachineApplicableSchedulesResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.LabVmsShutdown = tmp.LabVmsShutdown.Defaults()

	tmp.LabVmsStartup = tmp.LabVmsStartup.Defaults()

	return &tmp
}

func ListVirtualMachineApplicableSchedulesOutput(ctx *pulumi.Context, args ListVirtualMachineApplicableSchedulesOutputArgs, opts ...pulumi.InvokeOption) ListVirtualMachineApplicableSchedulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVirtualMachineApplicableSchedulesResult, error) {
			args := v.(ListVirtualMachineApplicableSchedulesArgs)
			r, err := ListVirtualMachineApplicableSchedules(ctx, &args, opts...)
			var s ListVirtualMachineApplicableSchedulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVirtualMachineApplicableSchedulesResultOutput)
}

type ListVirtualMachineApplicableSchedulesOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListVirtualMachineApplicableSchedulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVirtualMachineApplicableSchedulesArgs)(nil)).Elem()
}


type ListVirtualMachineApplicableSchedulesResultOutput struct{ *pulumi.OutputState }

func (ListVirtualMachineApplicableSchedulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVirtualMachineApplicableSchedulesResult)(nil)).Elem()
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) ToListVirtualMachineApplicableSchedulesResultOutput() ListVirtualMachineApplicableSchedulesResultOutput {
	return o
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) ToListVirtualMachineApplicableSchedulesResultOutputWithContext(ctx context.Context) ListVirtualMachineApplicableSchedulesResultOutput {
	return o
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) LabVmsShutdown() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) *ScheduleResponse { return v.LabVmsShutdown }).(ScheduleResponsePtrOutput)
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) LabVmsStartup() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) *ScheduleResponse { return v.LabVmsStartup }).(ScheduleResponsePtrOutput)
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVirtualMachineApplicableSchedulesResultOutput{})
}
