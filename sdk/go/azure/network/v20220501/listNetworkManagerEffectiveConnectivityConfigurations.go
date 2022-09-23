


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNetworkManagerEffectiveConnectivityConfigurations(ctx *pulumi.Context, args *ListNetworkManagerEffectiveConnectivityConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListNetworkManagerEffectiveConnectivityConfigurationsResult, error) {
	var rv ListNetworkManagerEffectiveConnectivityConfigurationsResult
	err := ctx.Invoke("azure-native:network/v20220501:listNetworkManagerEffectiveConnectivityConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNetworkManagerEffectiveConnectivityConfigurationsArgs struct {
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SkipToken          *string `pulumi:"skipToken"`
	Top                *int    `pulumi:"top"`
	VirtualNetworkName string  `pulumi:"virtualNetworkName"`
}


type ListNetworkManagerEffectiveConnectivityConfigurationsResult struct {
	SkipToken *string                                      `pulumi:"skipToken"`
	Value     []EffectiveConnectivityConfigurationResponse `pulumi:"value"`
}

func ListNetworkManagerEffectiveConnectivityConfigurationsOutput(ctx *pulumi.Context, args ListNetworkManagerEffectiveConnectivityConfigurationsOutputArgs, opts ...pulumi.InvokeOption) ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNetworkManagerEffectiveConnectivityConfigurationsResult, error) {
			args := v.(ListNetworkManagerEffectiveConnectivityConfigurationsArgs)
			r, err := ListNetworkManagerEffectiveConnectivityConfigurations(ctx, &args, opts...)
			var s ListNetworkManagerEffectiveConnectivityConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput)
}

type ListNetworkManagerEffectiveConnectivityConfigurationsOutputArgs struct {
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput `pulumi:"skipToken"`
	Top                pulumi.IntPtrInput    `pulumi:"top"`
	VirtualNetworkName pulumi.StringInput    `pulumi:"virtualNetworkName"`
}

func (ListNetworkManagerEffectiveConnectivityConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerEffectiveConnectivityConfigurationsArgs)(nil)).Elem()
}


type ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput struct{ *pulumi.OutputState }

func (ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerEffectiveConnectivityConfigurationsResult)(nil)).Elem()
}

func (o ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput) ToListNetworkManagerEffectiveConnectivityConfigurationsResultOutput() ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput {
	return o
}

func (o ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput) ToListNetworkManagerEffectiveConnectivityConfigurationsResultOutputWithContext(ctx context.Context) ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput {
	return o
}

func (o ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNetworkManagerEffectiveConnectivityConfigurationsResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput) Value() EffectiveConnectivityConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListNetworkManagerEffectiveConnectivityConfigurationsResult) []EffectiveConnectivityConfigurationResponse {
		return v.Value
	}).(EffectiveConnectivityConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput{})
}
