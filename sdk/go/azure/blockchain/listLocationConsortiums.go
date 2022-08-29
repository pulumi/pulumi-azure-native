


package blockchain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListLocationConsortiums(ctx *pulumi.Context, args *ListLocationConsortiumsArgs, opts ...pulumi.InvokeOption) (*ListLocationConsortiumsResult, error) {
	var rv ListLocationConsortiumsResult
	err := ctx.Invoke("azure-native:blockchain:listLocationConsortiums", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLocationConsortiumsArgs struct {
	LocationName string `pulumi:"locationName"`
}


type ListLocationConsortiumsResult struct {
	Value []ConsortiumResponse `pulumi:"value"`
}

func ListLocationConsortiumsOutput(ctx *pulumi.Context, args ListLocationConsortiumsOutputArgs, opts ...pulumi.InvokeOption) ListLocationConsortiumsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListLocationConsortiumsResult, error) {
			args := v.(ListLocationConsortiumsArgs)
			r, err := ListLocationConsortiums(ctx, &args, opts...)
			var s ListLocationConsortiumsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListLocationConsortiumsResultOutput)
}

type ListLocationConsortiumsOutputArgs struct {
	LocationName pulumi.StringInput `pulumi:"locationName"`
}

func (ListLocationConsortiumsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLocationConsortiumsArgs)(nil)).Elem()
}


type ListLocationConsortiumsResultOutput struct{ *pulumi.OutputState }

func (ListLocationConsortiumsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLocationConsortiumsResult)(nil)).Elem()
}

func (o ListLocationConsortiumsResultOutput) ToListLocationConsortiumsResultOutput() ListLocationConsortiumsResultOutput {
	return o
}

func (o ListLocationConsortiumsResultOutput) ToListLocationConsortiumsResultOutputWithContext(ctx context.Context) ListLocationConsortiumsResultOutput {
	return o
}

func (o ListLocationConsortiumsResultOutput) Value() ConsortiumResponseArrayOutput {
	return o.ApplyT(func(v ListLocationConsortiumsResult) []ConsortiumResponse { return v.Value }).(ConsortiumResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListLocationConsortiumsResultOutput{})
}
