


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApiManagementServiceSsoToken(ctx *pulumi.Context, args *GetApiManagementServiceSsoTokenArgs, opts ...pulumi.InvokeOption) (*GetApiManagementServiceSsoTokenResult, error) {
	var rv GetApiManagementServiceSsoTokenResult
	err := ctx.Invoke("azure-native:apimanagement/v20201201:getApiManagementServiceSsoToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetApiManagementServiceSsoTokenArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type GetApiManagementServiceSsoTokenResult struct {
	RedirectUri *string `pulumi:"redirectUri"`
}
