


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOpenIdConnectProviderSecrets(ctx *pulumi.Context, args *ListOpenIdConnectProviderSecretsArgs, opts ...pulumi.InvokeOption) (*ListOpenIdConnectProviderSecretsResult, error) {
	var rv ListOpenIdConnectProviderSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20201201:listOpenIdConnectProviderSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOpenIdConnectProviderSecretsArgs struct {
	Opid              string `pulumi:"opid"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListOpenIdConnectProviderSecretsResult struct {
	ClientSecret *string `pulumi:"clientSecret"`
}

func ListOpenIdConnectProviderSecretsOutput(ctx *pulumi.Context, args ListOpenIdConnectProviderSecretsOutputArgs, opts ...pulumi.InvokeOption) ListOpenIdConnectProviderSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOpenIdConnectProviderSecretsResult, error) {
			args := v.(ListOpenIdConnectProviderSecretsArgs)
			r, err := ListOpenIdConnectProviderSecrets(ctx, &args, opts...)
			var s ListOpenIdConnectProviderSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListOpenIdConnectProviderSecretsResultOutput)
}

type ListOpenIdConnectProviderSecretsOutputArgs struct {
	Opid              pulumi.StringInput `pulumi:"opid"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (ListOpenIdConnectProviderSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOpenIdConnectProviderSecretsArgs)(nil)).Elem()
}


type ListOpenIdConnectProviderSecretsResultOutput struct{ *pulumi.OutputState }

func (ListOpenIdConnectProviderSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOpenIdConnectProviderSecretsResult)(nil)).Elem()
}

func (o ListOpenIdConnectProviderSecretsResultOutput) ToListOpenIdConnectProviderSecretsResultOutput() ListOpenIdConnectProviderSecretsResultOutput {
	return o
}

func (o ListOpenIdConnectProviderSecretsResultOutput) ToListOpenIdConnectProviderSecretsResultOutputWithContext(ctx context.Context) ListOpenIdConnectProviderSecretsResultOutput {
	return o
}

func (o ListOpenIdConnectProviderSecretsResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOpenIdConnectProviderSecretsResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOpenIdConnectProviderSecretsResultOutput{})
}
