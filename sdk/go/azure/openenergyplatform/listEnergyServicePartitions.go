


package openenergyplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEnergyServicePartitions(ctx *pulumi.Context, args *ListEnergyServicePartitionsArgs, opts ...pulumi.InvokeOption) (*ListEnergyServicePartitionsResult, error) {
	var rv ListEnergyServicePartitionsResult
	err := ctx.Invoke("azure-native:openenergyplatform:listEnergyServicePartitions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEnergyServicePartitionsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListEnergyServicePartitionsResult struct {
	DataPartitionInfo []DataPartitionPropertiesResponse `pulumi:"dataPartitionInfo"`
}

func ListEnergyServicePartitionsOutput(ctx *pulumi.Context, args ListEnergyServicePartitionsOutputArgs, opts ...pulumi.InvokeOption) ListEnergyServicePartitionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEnergyServicePartitionsResult, error) {
			args := v.(ListEnergyServicePartitionsArgs)
			r, err := ListEnergyServicePartitions(ctx, &args, opts...)
			var s ListEnergyServicePartitionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEnergyServicePartitionsResultOutput)
}

type ListEnergyServicePartitionsOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListEnergyServicePartitionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEnergyServicePartitionsArgs)(nil)).Elem()
}


type ListEnergyServicePartitionsResultOutput struct{ *pulumi.OutputState }

func (ListEnergyServicePartitionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEnergyServicePartitionsResult)(nil)).Elem()
}

func (o ListEnergyServicePartitionsResultOutput) ToListEnergyServicePartitionsResultOutput() ListEnergyServicePartitionsResultOutput {
	return o
}

func (o ListEnergyServicePartitionsResultOutput) ToListEnergyServicePartitionsResultOutputWithContext(ctx context.Context) ListEnergyServicePartitionsResultOutput {
	return o
}

func (o ListEnergyServicePartitionsResultOutput) DataPartitionInfo() DataPartitionPropertiesResponseArrayOutput {
	return o.ApplyT(func(v ListEnergyServicePartitionsResult) []DataPartitionPropertiesResponse {
		return v.DataPartitionInfo
	}).(DataPartitionPropertiesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEnergyServicePartitionsResultOutput{})
}
