


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEffectiveConnectivityConfiguration(ctx *pulumi.Context, args *ListEffectiveConnectivityConfigurationArgs, opts ...pulumi.InvokeOption) (*ListEffectiveConnectivityConfigurationResult, error) {
	var rv ListEffectiveConnectivityConfigurationResult
	err := ctx.Invoke("azure-native:network/v20210201preview:listEffectiveConnectivityConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEffectiveConnectivityConfigurationArgs struct {
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SkipToken          *string `pulumi:"skipToken"`
	VirtualNetworkName string  `pulumi:"virtualNetworkName"`
}


type ListEffectiveConnectivityConfigurationResult struct {
	SkipToken *string                                      `pulumi:"skipToken"`
	Value     []EffectiveConnectivityConfigurationResponse `pulumi:"value"`
}

func ListEffectiveConnectivityConfigurationOutput(ctx *pulumi.Context, args ListEffectiveConnectivityConfigurationOutputArgs, opts ...pulumi.InvokeOption) ListEffectiveConnectivityConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEffectiveConnectivityConfigurationResult, error) {
			args := v.(ListEffectiveConnectivityConfigurationArgs)
			r, err := ListEffectiveConnectivityConfiguration(ctx, &args, opts...)
			var s ListEffectiveConnectivityConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEffectiveConnectivityConfigurationResultOutput)
}

type ListEffectiveConnectivityConfigurationOutputArgs struct {
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput `pulumi:"skipToken"`
	VirtualNetworkName pulumi.StringInput    `pulumi:"virtualNetworkName"`
}

func (ListEffectiveConnectivityConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEffectiveConnectivityConfigurationArgs)(nil)).Elem()
}


type ListEffectiveConnectivityConfigurationResultOutput struct{ *pulumi.OutputState }

func (ListEffectiveConnectivityConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEffectiveConnectivityConfigurationResult)(nil)).Elem()
}

func (o ListEffectiveConnectivityConfigurationResultOutput) ToListEffectiveConnectivityConfigurationResultOutput() ListEffectiveConnectivityConfigurationResultOutput {
	return o
}

func (o ListEffectiveConnectivityConfigurationResultOutput) ToListEffectiveConnectivityConfigurationResultOutputWithContext(ctx context.Context) ListEffectiveConnectivityConfigurationResultOutput {
	return o
}

func (o ListEffectiveConnectivityConfigurationResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListEffectiveConnectivityConfigurationResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListEffectiveConnectivityConfigurationResultOutput) Value() EffectiveConnectivityConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListEffectiveConnectivityConfigurationResult) []EffectiveConnectivityConfigurationResponse {
		return v.Value
	}).(EffectiveConnectivityConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEffectiveConnectivityConfigurationResultOutput{})
}
