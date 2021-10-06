


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIdentityProviderSecrets(ctx *pulumi.Context, args *ListIdentityProviderSecretsArgs, opts ...pulumi.InvokeOption) (*ListIdentityProviderSecretsResult, error) {
	var rv ListIdentityProviderSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20210401preview:listIdentityProviderSecrets", args, &rv, opts...)
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
