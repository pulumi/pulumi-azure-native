


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOnlineEndpointToken(ctx *pulumi.Context, args *GetOnlineEndpointTokenArgs, opts ...pulumi.InvokeOption) (*GetOnlineEndpointTokenResult, error) {
	var rv GetOnlineEndpointTokenResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getOnlineEndpointToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetOnlineEndpointTokenArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type GetOnlineEndpointTokenResult struct {
	AccessToken         *string  `pulumi:"accessToken"`
	ExpiryTimeUtc       *float64 `pulumi:"expiryTimeUtc"`
	RefreshAfterTimeUtc *float64 `pulumi:"refreshAfterTimeUtc"`
	TokenType           *string  `pulumi:"tokenType"`
}
