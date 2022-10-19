


package v20200218preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMachineLearningComputeNodes(ctx *pulumi.Context, args *ListMachineLearningComputeNodesArgs, opts ...pulumi.InvokeOption) (*ListMachineLearningComputeNodesResult, error) {
	var rv ListMachineLearningComputeNodesResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200218preview:listMachineLearningComputeNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMachineLearningComputeNodesArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListMachineLearningComputeNodesResult struct {
	ComputeType string                              `pulumi:"computeType"`
	NextLink    string                              `pulumi:"nextLink"`
	Nodes       []AmlComputeNodeInformationResponse `pulumi:"nodes"`
}

func ListMachineLearningComputeNodesOutput(ctx *pulumi.Context, args ListMachineLearningComputeNodesOutputArgs, opts ...pulumi.InvokeOption) ListMachineLearningComputeNodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMachineLearningComputeNodesResult, error) {
			args := v.(ListMachineLearningComputeNodesArgs)
			r, err := ListMachineLearningComputeNodes(ctx, &args, opts...)
			var s ListMachineLearningComputeNodesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMachineLearningComputeNodesResultOutput)
}

type ListMachineLearningComputeNodesOutputArgs struct {
	ComputeName       pulumi.StringInput `pulumi:"computeName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListMachineLearningComputeNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMachineLearningComputeNodesArgs)(nil)).Elem()
}


type ListMachineLearningComputeNodesResultOutput struct{ *pulumi.OutputState }

func (ListMachineLearningComputeNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMachineLearningComputeNodesResult)(nil)).Elem()
}

func (o ListMachineLearningComputeNodesResultOutput) ToListMachineLearningComputeNodesResultOutput() ListMachineLearningComputeNodesResultOutput {
	return o
}

func (o ListMachineLearningComputeNodesResultOutput) ToListMachineLearningComputeNodesResultOutputWithContext(ctx context.Context) ListMachineLearningComputeNodesResultOutput {
	return o
}

func (o ListMachineLearningComputeNodesResultOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v ListMachineLearningComputeNodesResult) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o ListMachineLearningComputeNodesResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListMachineLearningComputeNodesResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListMachineLearningComputeNodesResultOutput) Nodes() AmlComputeNodeInformationResponseArrayOutput {
	return o.ApplyT(func(v ListMachineLearningComputeNodesResult) []AmlComputeNodeInformationResponse { return v.Nodes }).(AmlComputeNodeInformationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMachineLearningComputeNodesResultOutput{})
}
