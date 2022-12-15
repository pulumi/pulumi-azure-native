


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAuthorizationServerSecrets(ctx *pulumi.Context, args *ListAuthorizationServerSecretsArgs, opts ...pulumi.InvokeOption) (*ListAuthorizationServerSecretsResult, error) {
	var rv ListAuthorizationServerSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:listAuthorizationServerSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAuthorizationServerSecretsArgs struct {
	Authsid           string `pulumi:"authsid"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListAuthorizationServerSecretsResult struct {
	ClientSecret          *string `pulumi:"clientSecret"`
	ResourceOwnerPassword *string `pulumi:"resourceOwnerPassword"`
	ResourceOwnerUsername *string `pulumi:"resourceOwnerUsername"`
}

func ListAuthorizationServerSecretsOutput(ctx *pulumi.Context, args ListAuthorizationServerSecretsOutputArgs, opts ...pulumi.InvokeOption) ListAuthorizationServerSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAuthorizationServerSecretsResult, error) {
			args := v.(ListAuthorizationServerSecretsArgs)
			r, err := ListAuthorizationServerSecrets(ctx, &args, opts...)
			var s ListAuthorizationServerSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAuthorizationServerSecretsResultOutput)
}

type ListAuthorizationServerSecretsOutputArgs struct {
	Authsid           pulumi.StringInput `pulumi:"authsid"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (ListAuthorizationServerSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAuthorizationServerSecretsArgs)(nil)).Elem()
}


type ListAuthorizationServerSecretsResultOutput struct{ *pulumi.OutputState }

func (ListAuthorizationServerSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAuthorizationServerSecretsResult)(nil)).Elem()
}

func (o ListAuthorizationServerSecretsResultOutput) ToListAuthorizationServerSecretsResultOutput() ListAuthorizationServerSecretsResultOutput {
	return o
}

func (o ListAuthorizationServerSecretsResultOutput) ToListAuthorizationServerSecretsResultOutputWithContext(ctx context.Context) ListAuthorizationServerSecretsResultOutput {
	return o
}

func (o ListAuthorizationServerSecretsResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListAuthorizationServerSecretsResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListAuthorizationServerSecretsResultOutput) ResourceOwnerPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListAuthorizationServerSecretsResult) *string { return v.ResourceOwnerPassword }).(pulumi.StringPtrOutput)
}

func (o ListAuthorizationServerSecretsResultOutput) ResourceOwnerUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListAuthorizationServerSecretsResult) *string { return v.ResourceOwnerUsername }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAuthorizationServerSecretsResultOutput{})
}
