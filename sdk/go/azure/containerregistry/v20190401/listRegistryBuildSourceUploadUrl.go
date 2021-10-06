


package v20190401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRegistryBuildSourceUploadUrl(ctx *pulumi.Context, args *ListRegistryBuildSourceUploadUrlArgs, opts ...pulumi.InvokeOption) (*ListRegistryBuildSourceUploadUrlResult, error) {
	var rv ListRegistryBuildSourceUploadUrlResult
	err := ctx.Invoke("azure-native:containerregistry/v20190401:listRegistryBuildSourceUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRegistryBuildSourceUploadUrlArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListRegistryBuildSourceUploadUrlResult struct {
	RelativePath *string `pulumi:"relativePath"`
	UploadUrl    *string `pulumi:"uploadUrl"`
}
