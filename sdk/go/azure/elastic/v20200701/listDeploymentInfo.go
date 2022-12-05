


package v20200701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDeploymentInfo(ctx *pulumi.Context, args *ListDeploymentInfoArgs, opts ...pulumi.InvokeOption) (*ListDeploymentInfoResult, error) {
	var rv ListDeploymentInfoResult
	err := ctx.Invoke("azure-native:elastic/v20200701:listDeploymentInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDeploymentInfoArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDeploymentInfoResult struct {
	DiskCapacity   string `pulumi:"diskCapacity"`
	MemoryCapacity string `pulumi:"memoryCapacity"`
	Status         string `pulumi:"status"`
	Version        string `pulumi:"version"`
}

func ListDeploymentInfoOutput(ctx *pulumi.Context, args ListDeploymentInfoOutputArgs, opts ...pulumi.InvokeOption) ListDeploymentInfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDeploymentInfoResult, error) {
			args := v.(ListDeploymentInfoArgs)
			r, err := ListDeploymentInfo(ctx, &args, opts...)
			var s ListDeploymentInfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDeploymentInfoResultOutput)
}

type ListDeploymentInfoOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDeploymentInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeploymentInfoArgs)(nil)).Elem()
}


type ListDeploymentInfoResultOutput struct{ *pulumi.OutputState }

func (ListDeploymentInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeploymentInfoResult)(nil)).Elem()
}

func (o ListDeploymentInfoResultOutput) ToListDeploymentInfoResultOutput() ListDeploymentInfoResultOutput {
	return o
}

func (o ListDeploymentInfoResultOutput) ToListDeploymentInfoResultOutputWithContext(ctx context.Context) ListDeploymentInfoResultOutput {
	return o
}

func (o ListDeploymentInfoResultOutput) DiskCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v ListDeploymentInfoResult) string { return v.DiskCapacity }).(pulumi.StringOutput)
}

func (o ListDeploymentInfoResultOutput) MemoryCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v ListDeploymentInfoResult) string { return v.MemoryCapacity }).(pulumi.StringOutput)
}

func (o ListDeploymentInfoResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ListDeploymentInfoResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o ListDeploymentInfoResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ListDeploymentInfoResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDeploymentInfoResultOutput{})
}
