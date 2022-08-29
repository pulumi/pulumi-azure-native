


package datadog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMonitorDefaultKey(ctx *pulumi.Context, args *GetMonitorDefaultKeyArgs, opts ...pulumi.InvokeOption) (*GetMonitorDefaultKeyResult, error) {
	var rv GetMonitorDefaultKeyResult
	err := ctx.Invoke("azure-native:datadog:getMonitorDefaultKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetMonitorDefaultKeyArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type GetMonitorDefaultKeyResult struct {
	Created   *string `pulumi:"created"`
	CreatedBy *string `pulumi:"createdBy"`
	Key       string  `pulumi:"key"`
	Name      *string `pulumi:"name"`
}

func GetMonitorDefaultKeyOutput(ctx *pulumi.Context, args GetMonitorDefaultKeyOutputArgs, opts ...pulumi.InvokeOption) GetMonitorDefaultKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMonitorDefaultKeyResult, error) {
			args := v.(GetMonitorDefaultKeyArgs)
			r, err := GetMonitorDefaultKey(ctx, &args, opts...)
			var s GetMonitorDefaultKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMonitorDefaultKeyResultOutput)
}

type GetMonitorDefaultKeyOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetMonitorDefaultKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitorDefaultKeyArgs)(nil)).Elem()
}

type GetMonitorDefaultKeyResultOutput struct{ *pulumi.OutputState }

func (GetMonitorDefaultKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitorDefaultKeyResult)(nil)).Elem()
}

func (o GetMonitorDefaultKeyResultOutput) ToGetMonitorDefaultKeyResultOutput() GetMonitorDefaultKeyResultOutput {
	return o
}

func (o GetMonitorDefaultKeyResultOutput) ToGetMonitorDefaultKeyResultOutputWithContext(ctx context.Context) GetMonitorDefaultKeyResultOutput {
	return o
}

func (o GetMonitorDefaultKeyResultOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorDefaultKeyResult) *string { return v.Created }).(pulumi.StringPtrOutput)
}

func (o GetMonitorDefaultKeyResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorDefaultKeyResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o GetMonitorDefaultKeyResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitorDefaultKeyResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o GetMonitorDefaultKeyResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorDefaultKeyResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitorDefaultKeyResultOutput{})
}
