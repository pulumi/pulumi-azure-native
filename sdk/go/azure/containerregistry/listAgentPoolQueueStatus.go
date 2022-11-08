


package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAgentPoolQueueStatus(ctx *pulumi.Context, args *ListAgentPoolQueueStatusArgs, opts ...pulumi.InvokeOption) (*ListAgentPoolQueueStatusResult, error) {
	var rv ListAgentPoolQueueStatusResult
	err := ctx.Invoke("azure-native:containerregistry:listAgentPoolQueueStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAgentPoolQueueStatusArgs struct {
	AgentPoolName     string `pulumi:"agentPoolName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAgentPoolQueueStatusResult struct {
	Count *int `pulumi:"count"`
}

func ListAgentPoolQueueStatusOutput(ctx *pulumi.Context, args ListAgentPoolQueueStatusOutputArgs, opts ...pulumi.InvokeOption) ListAgentPoolQueueStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAgentPoolQueueStatusResult, error) {
			args := v.(ListAgentPoolQueueStatusArgs)
			r, err := ListAgentPoolQueueStatus(ctx, &args, opts...)
			var s ListAgentPoolQueueStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAgentPoolQueueStatusResultOutput)
}

type ListAgentPoolQueueStatusOutputArgs struct {
	AgentPoolName     pulumi.StringInput `pulumi:"agentPoolName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAgentPoolQueueStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAgentPoolQueueStatusArgs)(nil)).Elem()
}


type ListAgentPoolQueueStatusResultOutput struct{ *pulumi.OutputState }

func (ListAgentPoolQueueStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAgentPoolQueueStatusResult)(nil)).Elem()
}

func (o ListAgentPoolQueueStatusResultOutput) ToListAgentPoolQueueStatusResultOutput() ListAgentPoolQueueStatusResultOutput {
	return o
}

func (o ListAgentPoolQueueStatusResultOutput) ToListAgentPoolQueueStatusResultOutputWithContext(ctx context.Context) ListAgentPoolQueueStatusResultOutput {
	return o
}

func (o ListAgentPoolQueueStatusResultOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ListAgentPoolQueueStatusResult) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAgentPoolQueueStatusResultOutput{})
}
