


package v20191201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOpenIdConnectProviderSecrets(ctx *pulumi.Context, args *ListOpenIdConnectProviderSecretsArgs, opts ...pulumi.InvokeOption) (*ListOpenIdConnectProviderSecretsResult, error) {
	var rv ListOpenIdConnectProviderSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201preview:listOpenIdConnectProviderSecrets", args, &rv, opts...)
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
