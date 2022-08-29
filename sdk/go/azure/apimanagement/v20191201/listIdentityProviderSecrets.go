


package v20191201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIdentityProviderSecrets(ctx *pulumi.Context, args *ListIdentityProviderSecretsArgs, opts ...pulumi.InvokeOption) (*ListIdentityProviderSecretsResult, error) {
	var rv ListIdentityProviderSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201:listIdentityProviderSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIdentityProviderSecretsArgs struct {
	IdentityProviderName string `pulumi:"identityProviderName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	ServiceName          string `pulumi:"serviceName"`
}


type ListIdentityProviderSecretsResult struct {
	ClientSecret *string `pulumi:"clientSecret"`
}

func ListIdentityProviderSecretsOutput(ctx *pulumi.Context, args ListIdentityProviderSecretsOutputArgs, opts ...pulumi.InvokeOption) ListIdentityProviderSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIdentityProviderSecretsResult, error) {
			args := v.(ListIdentityProviderSecretsArgs)
			r, err := ListIdentityProviderSecrets(ctx, &args, opts...)
			var s ListIdentityProviderSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIdentityProviderSecretsResultOutput)
}

type ListIdentityProviderSecretsOutputArgs struct {
	IdentityProviderName pulumi.StringInput `pulumi:"identityProviderName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName          pulumi.StringInput `pulumi:"serviceName"`
}

func (ListIdentityProviderSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIdentityProviderSecretsArgs)(nil)).Elem()
}


type ListIdentityProviderSecretsResultOutput struct{ *pulumi.OutputState }

func (ListIdentityProviderSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIdentityProviderSecretsResult)(nil)).Elem()
}

func (o ListIdentityProviderSecretsResultOutput) ToListIdentityProviderSecretsResultOutput() ListIdentityProviderSecretsResultOutput {
	return o
}

func (o ListIdentityProviderSecretsResultOutput) ToListIdentityProviderSecretsResultOutputWithContext(ctx context.Context) ListIdentityProviderSecretsResultOutput {
	return o
}

func (o ListIdentityProviderSecretsResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIdentityProviderSecretsResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIdentityProviderSecretsResultOutput{})
}
