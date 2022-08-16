


package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConfigurations(ctx *pulumi.Context, args *ListConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListConfigurationsResult, error) {
	var rv ListConfigurationsResult
	err := ctx.Invoke("azure-native:edgeorder/v20201201preview:listConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConfigurationsArgs struct {
	ConfigurationFilters        []ConfigurationFilters       `pulumi:"configurationFilters"`
	CustomerSubscriptionDetails *CustomerSubscriptionDetails `pulumi:"customerSubscriptionDetails"`
	SkipToken                   *string                      `pulumi:"skipToken"`
}


type ListConfigurationsResult struct {
	NextLink *string                 `pulumi:"nextLink"`
	Value    []ConfigurationResponse `pulumi:"value"`
}

func ListConfigurationsOutput(ctx *pulumi.Context, args ListConfigurationsOutputArgs, opts ...pulumi.InvokeOption) ListConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConfigurationsResult, error) {
			args := v.(ListConfigurationsArgs)
			r, err := ListConfigurations(ctx, &args, opts...)
			var s ListConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConfigurationsResultOutput)
}

type ListConfigurationsOutputArgs struct {
	ConfigurationFilters        ConfigurationFiltersArrayInput      `pulumi:"configurationFilters"`
	CustomerSubscriptionDetails CustomerSubscriptionDetailsPtrInput `pulumi:"customerSubscriptionDetails"`
	SkipToken                   pulumi.StringPtrInput               `pulumi:"skipToken"`
}

func (ListConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigurationsArgs)(nil)).Elem()
}


type ListConfigurationsResultOutput struct{ *pulumi.OutputState }

func (ListConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigurationsResult)(nil)).Elem()
}

func (o ListConfigurationsResultOutput) ToListConfigurationsResultOutput() ListConfigurationsResultOutput {
	return o
}

func (o ListConfigurationsResultOutput) ToListConfigurationsResultOutputWithContext(ctx context.Context) ListConfigurationsResultOutput {
	return o
}

func (o ListConfigurationsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListConfigurationsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListConfigurationsResultOutput) Value() ConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListConfigurationsResult) []ConfigurationResponse { return v.Value }).(ConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConfigurationsResultOutput{})
}
