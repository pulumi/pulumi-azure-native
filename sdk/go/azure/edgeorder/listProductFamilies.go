


package edgeorder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListProductFamilies(ctx *pulumi.Context, args *ListProductFamiliesArgs, opts ...pulumi.InvokeOption) (*ListProductFamiliesResult, error) {
	var rv ListProductFamiliesResult
	err := ctx.Invoke("azure-native:edgeorder:listProductFamilies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProductFamiliesArgs struct {
	CustomerSubscriptionDetails *CustomerSubscriptionDetails    `pulumi:"customerSubscriptionDetails"`
	Expand                      *string                         `pulumi:"expand"`
	FilterableProperties        map[string][]FilterableProperty `pulumi:"filterableProperties"`
	SkipToken                   *string                         `pulumi:"skipToken"`
}


type ListProductFamiliesResult struct {
	NextLink *string                 `pulumi:"nextLink"`
	Value    []ProductFamilyResponse `pulumi:"value"`
}

func ListProductFamiliesOutput(ctx *pulumi.Context, args ListProductFamiliesOutputArgs, opts ...pulumi.InvokeOption) ListProductFamiliesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListProductFamiliesResult, error) {
			args := v.(ListProductFamiliesArgs)
			r, err := ListProductFamilies(ctx, &args, opts...)
			var s ListProductFamiliesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListProductFamiliesResultOutput)
}

type ListProductFamiliesOutputArgs struct {
	CustomerSubscriptionDetails CustomerSubscriptionDetailsPtrInput `pulumi:"customerSubscriptionDetails"`
	Expand                      pulumi.StringPtrInput               `pulumi:"expand"`
	FilterableProperties        FilterablePropertyArrayMapInput     `pulumi:"filterableProperties"`
	SkipToken                   pulumi.StringPtrInput               `pulumi:"skipToken"`
}

func (ListProductFamiliesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductFamiliesArgs)(nil)).Elem()
}


type ListProductFamiliesResultOutput struct{ *pulumi.OutputState }

func (ListProductFamiliesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductFamiliesResult)(nil)).Elem()
}

func (o ListProductFamiliesResultOutput) ToListProductFamiliesResultOutput() ListProductFamiliesResultOutput {
	return o
}

func (o ListProductFamiliesResultOutput) ToListProductFamiliesResultOutputWithContext(ctx context.Context) ListProductFamiliesResultOutput {
	return o
}

func (o ListProductFamiliesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListProductFamiliesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListProductFamiliesResultOutput) Value() ProductFamilyResponseArrayOutput {
	return o.ApplyT(func(v ListProductFamiliesResult) []ProductFamilyResponse { return v.Value }).(ProductFamilyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProductFamiliesResultOutput{})
}
