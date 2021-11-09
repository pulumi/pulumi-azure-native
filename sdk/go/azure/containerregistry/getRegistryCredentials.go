


package containerregistry

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRegistryCredentials(ctx *pulumi.Context, args *GetRegistryCredentialsArgs, opts ...pulumi.InvokeOption) (*GetRegistryCredentialsResult, error) {
	var rv GetRegistryCredentialsResult
	err := ctx.Invoke("azure-native:containerregistry:getRegistryCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetRegistryCredentialsArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetRegistryCredentialsResult struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}
