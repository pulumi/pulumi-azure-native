


package v20201101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRegistryCredentials(ctx *pulumi.Context, args *ListRegistryCredentialsArgs, opts ...pulumi.InvokeOption) (*ListRegistryCredentialsResult, error) {
	var rv ListRegistryCredentialsResult
	err := ctx.Invoke("azure-native:containerregistry/v20201101preview:listRegistryCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRegistryCredentialsArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListRegistryCredentialsResult struct {
	Passwords []RegistryPasswordResponse `pulumi:"passwords"`
	Username  *string                    `pulumi:"username"`
}
