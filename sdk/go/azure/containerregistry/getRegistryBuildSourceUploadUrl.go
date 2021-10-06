


package containerregistry

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRegistryBuildSourceUploadUrl(ctx *pulumi.Context, args *GetRegistryBuildSourceUploadUrlArgs, opts ...pulumi.InvokeOption) (*GetRegistryBuildSourceUploadUrlResult, error) {
	var rv GetRegistryBuildSourceUploadUrlResult
	err := ctx.Invoke("azure-native:containerregistry:getRegistryBuildSourceUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetRegistryBuildSourceUploadUrlArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetRegistryBuildSourceUploadUrlResult struct {
	RelativePath *string `pulumi:"relativePath"`
	UploadUrl    *string `pulumi:"uploadUrl"`
}
