


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAuthorizationServerSecrets(ctx *pulumi.Context, args *ListAuthorizationServerSecretsArgs, opts ...pulumi.InvokeOption) (*ListAuthorizationServerSecretsResult, error) {
	var rv ListAuthorizationServerSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20210801:listAuthorizationServerSecrets", args, &rv, opts...)
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
