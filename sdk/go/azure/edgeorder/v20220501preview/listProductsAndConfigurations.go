


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListProductsAndConfigurations(ctx *pulumi.Context, args *ListProductsAndConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListProductsAndConfigurationsResult, error) {
	var rv ListProductsAndConfigurationsResult
	err := ctx.Invoke("azure-native:edgeorder/v20220501preview:listProductsAndConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProductsAndConfigurationsArgs struct {
	ConfigurationFilter         *ConfigurationFilter         `pulumi:"configurationFilter"`
	CustomerSubscriptionDetails *CustomerSubscriptionDetails `pulumi:"customerSubscriptionDetails"`
	SkipToken                   *string                      `pulumi:"skipToken"`
}


type ListProductsAndConfigurationsResult struct {
	NextLink *string                 `pulumi:"nextLink"`
	Value    []ConfigurationResponse `pulumi:"value"`
}

func ListProductsAndConfigurationsOutput(ctx *pulumi.Context, args ListProductsAndConfigurationsOutputArgs, opts ...pulumi.InvokeOption) ListProductsAndConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListProductsAndConfigurationsResult, error) {
			args := v.(ListProductsAndConfigurationsArgs)
			r, err := ListProductsAndConfigurations(ctx, &args, opts...)
			var s ListProductsAndConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListProductsAndConfigurationsResultOutput)
}

type ListProductsAndConfigurationsOutputArgs struct {
	ConfigurationFilter         ConfigurationFilterPtrInput         `pulumi:"configurationFilter"`
	CustomerSubscriptionDetails CustomerSubscriptionDetailsPtrInput `pulumi:"customerSubscriptionDetails"`
	SkipToken                   pulumi.StringPtrInput               `pulumi:"skipToken"`
}

func (ListProductsAndConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductsAndConfigurationsArgs)(nil)).Elem()
}


type ListProductsAndConfigurationsResultOutput struct{ *pulumi.OutputState }

func (ListProductsAndConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductsAndConfigurationsResult)(nil)).Elem()
}

func (o ListProductsAndConfigurationsResultOutput) ToListProductsAndConfigurationsResultOutput() ListProductsAndConfigurationsResultOutput {
	return o
}

func (o ListProductsAndConfigurationsResultOutput) ToListProductsAndConfigurationsResultOutputWithContext(ctx context.Context) ListProductsAndConfigurationsResultOutput {
	return o
}

func (o ListProductsAndConfigurationsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListProductsAndConfigurationsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListProductsAndConfigurationsResultOutput) Value() ConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListProductsAndConfigurationsResult) []ConfigurationResponse { return v.Value }).(ConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProductsAndConfigurationsResultOutput{})
}
