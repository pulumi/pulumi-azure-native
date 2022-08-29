


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListComputeKeys(ctx *pulumi.Context, args *ListComputeKeysArgs, opts ...pulumi.InvokeOption) (*ListComputeKeysResult, error) {
	var rv ListComputeKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220501:listComputeKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListComputeKeysArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListComputeKeysResult struct {
	ComputeType string `pulumi:"computeType"`
}

func ListComputeKeysOutput(ctx *pulumi.Context, args ListComputeKeysOutputArgs, opts ...pulumi.InvokeOption) ListComputeKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListComputeKeysResult, error) {
			args := v.(ListComputeKeysArgs)
			r, err := ListComputeKeys(ctx, &args, opts...)
			var s ListComputeKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListComputeKeysResultOutput)
}

type ListComputeKeysOutputArgs struct {
	ComputeName       pulumi.StringInput `pulumi:"computeName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListComputeKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListComputeKeysArgs)(nil)).Elem()
}


type ListComputeKeysResultOutput struct{ *pulumi.OutputState }

func (ListComputeKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListComputeKeysResult)(nil)).Elem()
}

func (o ListComputeKeysResultOutput) ToListComputeKeysResultOutput() ListComputeKeysResultOutput {
	return o
}

func (o ListComputeKeysResultOutput) ToListComputeKeysResultOutputWithContext(ctx context.Context) ListComputeKeysResultOutput {
	return o
}

func (o ListComputeKeysResultOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v ListComputeKeysResult) string { return v.ComputeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListComputeKeysResultOutput{})
}
