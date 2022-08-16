


package engagementfabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAccountChannelTypes(ctx *pulumi.Context, args *ListAccountChannelTypesArgs, opts ...pulumi.InvokeOption) (*ListAccountChannelTypesResult, error) {
	var rv ListAccountChannelTypesResult
	err := ctx.Invoke("azure-native:engagementfabric:listAccountChannelTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountChannelTypesArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAccountChannelTypesResult struct {
	Value []ChannelTypeDescriptionResponse `pulumi:"value"`
}

func ListAccountChannelTypesOutput(ctx *pulumi.Context, args ListAccountChannelTypesOutputArgs, opts ...pulumi.InvokeOption) ListAccountChannelTypesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAccountChannelTypesResult, error) {
			args := v.(ListAccountChannelTypesArgs)
			r, err := ListAccountChannelTypes(ctx, &args, opts...)
			var s ListAccountChannelTypesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAccountChannelTypesResultOutput)
}

type ListAccountChannelTypesOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAccountChannelTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountChannelTypesArgs)(nil)).Elem()
}


type ListAccountChannelTypesResultOutput struct{ *pulumi.OutputState }

func (ListAccountChannelTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountChannelTypesResult)(nil)).Elem()
}

func (o ListAccountChannelTypesResultOutput) ToListAccountChannelTypesResultOutput() ListAccountChannelTypesResultOutput {
	return o
}

func (o ListAccountChannelTypesResultOutput) ToListAccountChannelTypesResultOutputWithContext(ctx context.Context) ListAccountChannelTypesResultOutput {
	return o
}

func (o ListAccountChannelTypesResultOutput) Value() ChannelTypeDescriptionResponseArrayOutput {
	return o.ApplyT(func(v ListAccountChannelTypesResult) []ChannelTypeDescriptionResponse { return v.Value }).(ChannelTypeDescriptionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAccountChannelTypesResultOutput{})
}
