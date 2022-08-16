


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListComputeNodes(ctx *pulumi.Context, args *ListComputeNodesArgs, opts ...pulumi.InvokeOption) (*ListComputeNodesResult, error) {
	var rv ListComputeNodesResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220101preview:listComputeNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListComputeNodesArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListComputeNodesResult struct {
	NextLink string                              `pulumi:"nextLink"`
	Nodes    []AmlComputeNodeInformationResponse `pulumi:"nodes"`
}

func ListComputeNodesOutput(ctx *pulumi.Context, args ListComputeNodesOutputArgs, opts ...pulumi.InvokeOption) ListComputeNodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListComputeNodesResult, error) {
			args := v.(ListComputeNodesArgs)
			r, err := ListComputeNodes(ctx, &args, opts...)
			var s ListComputeNodesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListComputeNodesResultOutput)
}

type ListComputeNodesOutputArgs struct {
	ComputeName       pulumi.StringInput `pulumi:"computeName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListComputeNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListComputeNodesArgs)(nil)).Elem()
}


type ListComputeNodesResultOutput struct{ *pulumi.OutputState }

func (ListComputeNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListComputeNodesResult)(nil)).Elem()
}

func (o ListComputeNodesResultOutput) ToListComputeNodesResultOutput() ListComputeNodesResultOutput {
	return o
}

func (o ListComputeNodesResultOutput) ToListComputeNodesResultOutputWithContext(ctx context.Context) ListComputeNodesResultOutput {
	return o
}

func (o ListComputeNodesResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListComputeNodesResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListComputeNodesResultOutput) Nodes() AmlComputeNodeInformationResponseArrayOutput {
	return o.ApplyT(func(v ListComputeNodesResult) []AmlComputeNodeInformationResponse { return v.Nodes }).(AmlComputeNodeInformationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListComputeNodesResultOutput{})
}
