


package v20200701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVMHost(ctx *pulumi.Context, args *ListVMHostArgs, opts ...pulumi.InvokeOption) (*ListVMHostResult, error) {
	var rv ListVMHostResult
	err := ctx.Invoke("azure-native:elastic/v20200701:listVMHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVMHostArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListVMHostResult struct {
	NextLink *string               `pulumi:"nextLink"`
	Value    []VMResourcesResponse `pulumi:"value"`
}

func ListVMHostOutput(ctx *pulumi.Context, args ListVMHostOutputArgs, opts ...pulumi.InvokeOption) ListVMHostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVMHostResult, error) {
			args := v.(ListVMHostArgs)
			r, err := ListVMHost(ctx, &args, opts...)
			var s ListVMHostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVMHostResultOutput)
}

type ListVMHostOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListVMHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVMHostArgs)(nil)).Elem()
}


type ListVMHostResultOutput struct{ *pulumi.OutputState }

func (ListVMHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVMHostResult)(nil)).Elem()
}

func (o ListVMHostResultOutput) ToListVMHostResultOutput() ListVMHostResultOutput {
	return o
}

func (o ListVMHostResultOutput) ToListVMHostResultOutputWithContext(ctx context.Context) ListVMHostResultOutput {
	return o
}

func (o ListVMHostResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVMHostResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListVMHostResultOutput) Value() VMResourcesResponseArrayOutput {
	return o.ApplyT(func(v ListVMHostResult) []VMResourcesResponse { return v.Value }).(VMResourcesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVMHostResultOutput{})
}
