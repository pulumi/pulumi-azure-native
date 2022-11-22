


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListActiveConnectivityConfigurations(ctx *pulumi.Context, args *ListActiveConnectivityConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListActiveConnectivityConfigurationsResult, error) {
	var rv ListActiveConnectivityConfigurationsResult
	err := ctx.Invoke("azure-native:network/v20220201preview:listActiveConnectivityConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveConnectivityConfigurationsArgs struct {
	NetworkManagerName string   `pulumi:"networkManagerName"`
	Regions            []string `pulumi:"regions"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SkipToken          *string  `pulumi:"skipToken"`
}


type ListActiveConnectivityConfigurationsResult struct {
	SkipToken *string                                   `pulumi:"skipToken"`
	Value     []ActiveConnectivityConfigurationResponse `pulumi:"value"`
}

func ListActiveConnectivityConfigurationsOutput(ctx *pulumi.Context, args ListActiveConnectivityConfigurationsOutputArgs, opts ...pulumi.InvokeOption) ListActiveConnectivityConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListActiveConnectivityConfigurationsResult, error) {
			args := v.(ListActiveConnectivityConfigurationsArgs)
			r, err := ListActiveConnectivityConfigurations(ctx, &args, opts...)
			var s ListActiveConnectivityConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListActiveConnectivityConfigurationsResultOutput)
}

type ListActiveConnectivityConfigurationsOutputArgs struct {
	NetworkManagerName pulumi.StringInput      `pulumi:"networkManagerName"`
	Regions            pulumi.StringArrayInput `pulumi:"regions"`
	ResourceGroupName  pulumi.StringInput      `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput   `pulumi:"skipToken"`
}

func (ListActiveConnectivityConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveConnectivityConfigurationsArgs)(nil)).Elem()
}


type ListActiveConnectivityConfigurationsResultOutput struct{ *pulumi.OutputState }

func (ListActiveConnectivityConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveConnectivityConfigurationsResult)(nil)).Elem()
}

func (o ListActiveConnectivityConfigurationsResultOutput) ToListActiveConnectivityConfigurationsResultOutput() ListActiveConnectivityConfigurationsResultOutput {
	return o
}

func (o ListActiveConnectivityConfigurationsResultOutput) ToListActiveConnectivityConfigurationsResultOutputWithContext(ctx context.Context) ListActiveConnectivityConfigurationsResultOutput {
	return o
}

func (o ListActiveConnectivityConfigurationsResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListActiveConnectivityConfigurationsResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListActiveConnectivityConfigurationsResultOutput) Value() ActiveConnectivityConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListActiveConnectivityConfigurationsResult) []ActiveConnectivityConfigurationResponse {
		return v.Value
	}).(ActiveConnectivityConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListActiveConnectivityConfigurationsResultOutput{})
}
