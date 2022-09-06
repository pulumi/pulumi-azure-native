


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListProductsAndConfigurationProductFamilies(ctx *pulumi.Context, args *ListProductsAndConfigurationProductFamiliesArgs, opts ...pulumi.InvokeOption) (*ListProductsAndConfigurationProductFamiliesResult, error) {
	var rv ListProductsAndConfigurationProductFamiliesResult
	err := ctx.Invoke("azure-native:edgeorder/v20220501preview:listProductsAndConfigurationProductFamilies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProductsAndConfigurationProductFamiliesArgs struct {
	CustomerSubscriptionDetails *CustomerSubscriptionDetails    `pulumi:"customerSubscriptionDetails"`
	Expand                      *string                         `pulumi:"expand"`
	FilterableProperties        map[string][]FilterableProperty `pulumi:"filterableProperties"`
	SkipToken                   *string                         `pulumi:"skipToken"`
}


type ListProductsAndConfigurationProductFamiliesResult struct {
	NextLink *string                 `pulumi:"nextLink"`
	Value    []ProductFamilyResponse `pulumi:"value"`
}

func ListProductsAndConfigurationProductFamiliesOutput(ctx *pulumi.Context, args ListProductsAndConfigurationProductFamiliesOutputArgs, opts ...pulumi.InvokeOption) ListProductsAndConfigurationProductFamiliesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListProductsAndConfigurationProductFamiliesResult, error) {
			args := v.(ListProductsAndConfigurationProductFamiliesArgs)
			r, err := ListProductsAndConfigurationProductFamilies(ctx, &args, opts...)
			var s ListProductsAndConfigurationProductFamiliesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListProductsAndConfigurationProductFamiliesResultOutput)
}

type ListProductsAndConfigurationProductFamiliesOutputArgs struct {
	CustomerSubscriptionDetails CustomerSubscriptionDetailsPtrInput `pulumi:"customerSubscriptionDetails"`
	Expand                      pulumi.StringPtrInput               `pulumi:"expand"`
	FilterableProperties        FilterablePropertyArrayMapInput     `pulumi:"filterableProperties"`
	SkipToken                   pulumi.StringPtrInput               `pulumi:"skipToken"`
}

func (ListProductsAndConfigurationProductFamiliesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductsAndConfigurationProductFamiliesArgs)(nil)).Elem()
}


type ListProductsAndConfigurationProductFamiliesResultOutput struct{ *pulumi.OutputState }

func (ListProductsAndConfigurationProductFamiliesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductsAndConfigurationProductFamiliesResult)(nil)).Elem()
}

func (o ListProductsAndConfigurationProductFamiliesResultOutput) ToListProductsAndConfigurationProductFamiliesResultOutput() ListProductsAndConfigurationProductFamiliesResultOutput {
	return o
}

func (o ListProductsAndConfigurationProductFamiliesResultOutput) ToListProductsAndConfigurationProductFamiliesResultOutputWithContext(ctx context.Context) ListProductsAndConfigurationProductFamiliesResultOutput {
	return o
}

func (o ListProductsAndConfigurationProductFamiliesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListProductsAndConfigurationProductFamiliesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListProductsAndConfigurationProductFamiliesResultOutput) Value() ProductFamilyResponseArrayOutput {
	return o.ApplyT(func(v ListProductsAndConfigurationProductFamiliesResult) []ProductFamilyResponse { return v.Value }).(ProductFamilyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProductsAndConfigurationProductFamiliesResultOutput{})
}
