


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListScheduleApplicable(ctx *pulumi.Context, args *ListScheduleApplicableArgs, opts ...pulumi.InvokeOption) (*ListScheduleApplicableResult, error) {
	var rv ListScheduleApplicableResult
	err := ctx.Invoke("azure-native:devtestlab:listScheduleApplicable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListScheduleApplicableArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListScheduleApplicableResult struct {
	NextLink *string            `pulumi:"nextLink"`
	Value    []ScheduleResponse `pulumi:"value"`
}

func ListScheduleApplicableOutput(ctx *pulumi.Context, args ListScheduleApplicableOutputArgs, opts ...pulumi.InvokeOption) ListScheduleApplicableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListScheduleApplicableResult, error) {
			args := v.(ListScheduleApplicableArgs)
			r, err := ListScheduleApplicable(ctx, &args, opts...)
			var s ListScheduleApplicableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListScheduleApplicableResultOutput)
}

type ListScheduleApplicableOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListScheduleApplicableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListScheduleApplicableArgs)(nil)).Elem()
}


type ListScheduleApplicableResultOutput struct{ *pulumi.OutputState }

func (ListScheduleApplicableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListScheduleApplicableResult)(nil)).Elem()
}

func (o ListScheduleApplicableResultOutput) ToListScheduleApplicableResultOutput() ListScheduleApplicableResultOutput {
	return o
}

func (o ListScheduleApplicableResultOutput) ToListScheduleApplicableResultOutputWithContext(ctx context.Context) ListScheduleApplicableResultOutput {
	return o
}

func (o ListScheduleApplicableResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListScheduleApplicableResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListScheduleApplicableResultOutput) Value() ScheduleResponseArrayOutput {
	return o.ApplyT(func(v ListScheduleApplicableResult) []ScheduleResponse { return v.Value }).(ScheduleResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListScheduleApplicableResultOutput{})
}
