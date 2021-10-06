


package v20180201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBuildTaskSourceRepositoryProperties(ctx *pulumi.Context, args *ListBuildTaskSourceRepositoryPropertiesArgs, opts ...pulumi.InvokeOption) (*ListBuildTaskSourceRepositoryPropertiesResult, error) {
	var rv ListBuildTaskSourceRepositoryPropertiesResult
	err := ctx.Invoke("azure-native:containerregistry/v20180201preview:listBuildTaskSourceRepositoryProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBuildTaskSourceRepositoryPropertiesArgs struct {
	BuildTaskName     string `pulumi:"buildTaskName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListBuildTaskSourceRepositoryPropertiesResult struct {
	IsCommitTriggerEnabled      *bool                          `pulumi:"isCommitTriggerEnabled"`
	RepositoryUrl               string                         `pulumi:"repositoryUrl"`
	SourceControlAuthProperties *SourceControlAuthInfoResponse `pulumi:"sourceControlAuthProperties"`
	SourceControlType           string                         `pulumi:"sourceControlType"`
}
