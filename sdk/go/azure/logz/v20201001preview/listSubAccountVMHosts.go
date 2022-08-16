


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSubAccountVMHosts(ctx *pulumi.Context, args *ListSubAccountVMHostsArgs, opts ...pulumi.InvokeOption) (*ListSubAccountVMHostsResult, error) {
	var rv ListSubAccountVMHostsResult
	err := ctx.Invoke("azure-native:logz/v20201001preview:listSubAccountVMHosts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSubAccountVMHostsArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SubAccountName    string `pulumi:"subAccountName"`
}


type ListSubAccountVMHostsResult struct {
	NextLink *string               `pulumi:"nextLink"`
	Value    []VMResourcesResponse `pulumi:"value"`
}

func ListSubAccountVMHostsOutput(ctx *pulumi.Context, args ListSubAccountVMHostsOutputArgs, opts ...pulumi.InvokeOption) ListSubAccountVMHostsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSubAccountVMHostsResult, error) {
			args := v.(ListSubAccountVMHostsArgs)
			r, err := ListSubAccountVMHosts(ctx, &args, opts...)
			var s ListSubAccountVMHostsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSubAccountVMHostsResultOutput)
}

type ListSubAccountVMHostsOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SubAccountName    pulumi.StringInput `pulumi:"subAccountName"`
}

func (ListSubAccountVMHostsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubAccountVMHostsArgs)(nil)).Elem()
}


type ListSubAccountVMHostsResultOutput struct{ *pulumi.OutputState }

func (ListSubAccountVMHostsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubAccountVMHostsResult)(nil)).Elem()
}

func (o ListSubAccountVMHostsResultOutput) ToListSubAccountVMHostsResultOutput() ListSubAccountVMHostsResultOutput {
	return o
}

func (o ListSubAccountVMHostsResultOutput) ToListSubAccountVMHostsResultOutputWithContext(ctx context.Context) ListSubAccountVMHostsResultOutput {
	return o
}

func (o ListSubAccountVMHostsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSubAccountVMHostsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListSubAccountVMHostsResultOutput) Value() VMResourcesResponseArrayOutput {
	return o.ApplyT(func(v ListSubAccountVMHostsResult) []VMResourcesResponse { return v.Value }).(VMResourcesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSubAccountVMHostsResultOutput{})
}
