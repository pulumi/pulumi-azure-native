


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListActiveConnectivityConfiguration(ctx *pulumi.Context, args *ListActiveConnectivityConfigurationArgs, opts ...pulumi.InvokeOption) (*ListActiveConnectivityConfigurationResult, error) {
	var rv ListActiveConnectivityConfigurationResult
	err := ctx.Invoke("azure-native:network/v20210201preview:listActiveConnectivityConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveConnectivityConfigurationArgs struct {
	NetworkManagerName string   `pulumi:"networkManagerName"`
	Regions            []string `pulumi:"regions"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SkipToken          *string  `pulumi:"skipToken"`
}


type ListActiveConnectivityConfigurationResult struct {
	SkipToken *string                                   `pulumi:"skipToken"`
	Value     []ActiveConnectivityConfigurationResponse `pulumi:"value"`
}

func ListActiveConnectivityConfigurationOutput(ctx *pulumi.Context, args ListActiveConnectivityConfigurationOutputArgs, opts ...pulumi.InvokeOption) ListActiveConnectivityConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListActiveConnectivityConfigurationResult, error) {
			args := v.(ListActiveConnectivityConfigurationArgs)
			r, err := ListActiveConnectivityConfiguration(ctx, &args, opts...)
			var s ListActiveConnectivityConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListActiveConnectivityConfigurationResultOutput)
}

type ListActiveConnectivityConfigurationOutputArgs struct {
	NetworkManagerName pulumi.StringInput      `pulumi:"networkManagerName"`
	Regions            pulumi.StringArrayInput `pulumi:"regions"`
	ResourceGroupName  pulumi.StringInput      `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput   `pulumi:"skipToken"`
}

func (ListActiveConnectivityConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveConnectivityConfigurationArgs)(nil)).Elem()
}


type ListActiveConnectivityConfigurationResultOutput struct{ *pulumi.OutputState }

func (ListActiveConnectivityConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveConnectivityConfigurationResult)(nil)).Elem()
}

func (o ListActiveConnectivityConfigurationResultOutput) ToListActiveConnectivityConfigurationResultOutput() ListActiveConnectivityConfigurationResultOutput {
	return o
}

func (o ListActiveConnectivityConfigurationResultOutput) ToListActiveConnectivityConfigurationResultOutputWithContext(ctx context.Context) ListActiveConnectivityConfigurationResultOutput {
	return o
}

func (o ListActiveConnectivityConfigurationResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListActiveConnectivityConfigurationResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListActiveConnectivityConfigurationResultOutput) Value() ActiveConnectivityConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListActiveConnectivityConfigurationResult) []ActiveConnectivityConfigurationResponse {
		return v.Value
	}).(ActiveConnectivityConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListActiveConnectivityConfigurationResultOutput{})
}
