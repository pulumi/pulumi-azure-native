


package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDeviceFailoverSets(ctx *pulumi.Context, args *ListDeviceFailoverSetsArgs, opts ...pulumi.InvokeOption) (*ListDeviceFailoverSetsResult, error) {
	var rv ListDeviceFailoverSetsResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:listDeviceFailoverSets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDeviceFailoverSetsArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDeviceFailoverSetsResult struct {
	Value []FailoverSetResponse `pulumi:"value"`
}

func ListDeviceFailoverSetsOutput(ctx *pulumi.Context, args ListDeviceFailoverSetsOutputArgs, opts ...pulumi.InvokeOption) ListDeviceFailoverSetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDeviceFailoverSetsResult, error) {
			args := v.(ListDeviceFailoverSetsArgs)
			r, err := ListDeviceFailoverSets(ctx, &args, opts...)
			var s ListDeviceFailoverSetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDeviceFailoverSetsResultOutput)
}

type ListDeviceFailoverSetsOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDeviceFailoverSetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeviceFailoverSetsArgs)(nil)).Elem()
}


type ListDeviceFailoverSetsResultOutput struct{ *pulumi.OutputState }

func (ListDeviceFailoverSetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeviceFailoverSetsResult)(nil)).Elem()
}

func (o ListDeviceFailoverSetsResultOutput) ToListDeviceFailoverSetsResultOutput() ListDeviceFailoverSetsResultOutput {
	return o
}

func (o ListDeviceFailoverSetsResultOutput) ToListDeviceFailoverSetsResultOutputWithContext(ctx context.Context) ListDeviceFailoverSetsResultOutput {
	return o
}

func (o ListDeviceFailoverSetsResultOutput) Value() FailoverSetResponseArrayOutput {
	return o.ApplyT(func(v ListDeviceFailoverSetsResult) []FailoverSetResponse { return v.Value }).(FailoverSetResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDeviceFailoverSetsResultOutput{})
}
