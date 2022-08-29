


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMachineLearningComputeKeys(ctx *pulumi.Context, args *ListMachineLearningComputeKeysArgs, opts ...pulumi.InvokeOption) (*ListMachineLearningComputeKeysResult, error) {
	var rv ListMachineLearningComputeKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200101:listMachineLearningComputeKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMachineLearningComputeKeysArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListMachineLearningComputeKeysResult struct {
	ComputeType string `pulumi:"computeType"`
}

func ListMachineLearningComputeKeysOutput(ctx *pulumi.Context, args ListMachineLearningComputeKeysOutputArgs, opts ...pulumi.InvokeOption) ListMachineLearningComputeKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMachineLearningComputeKeysResult, error) {
			args := v.(ListMachineLearningComputeKeysArgs)
			r, err := ListMachineLearningComputeKeys(ctx, &args, opts...)
			var s ListMachineLearningComputeKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMachineLearningComputeKeysResultOutput)
}

type ListMachineLearningComputeKeysOutputArgs struct {
	ComputeName       pulumi.StringInput `pulumi:"computeName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListMachineLearningComputeKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMachineLearningComputeKeysArgs)(nil)).Elem()
}


type ListMachineLearningComputeKeysResultOutput struct{ *pulumi.OutputState }

func (ListMachineLearningComputeKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMachineLearningComputeKeysResult)(nil)).Elem()
}

func (o ListMachineLearningComputeKeysResultOutput) ToListMachineLearningComputeKeysResultOutput() ListMachineLearningComputeKeysResultOutput {
	return o
}

func (o ListMachineLearningComputeKeysResultOutput) ToListMachineLearningComputeKeysResultOutputWithContext(ctx context.Context) ListMachineLearningComputeKeysResultOutput {
	return o
}

func (o ListMachineLearningComputeKeysResultOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v ListMachineLearningComputeKeysResult) string { return v.ComputeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMachineLearningComputeKeysResultOutput{})
}
