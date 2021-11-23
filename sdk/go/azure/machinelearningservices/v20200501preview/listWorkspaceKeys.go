


package v20200501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceKeys(ctx *pulumi.Context, args *ListWorkspaceKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceKeysResult, error) {
	var rv ListWorkspaceKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200501preview:listWorkspaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type ListWorkspaceKeysResult struct {
	AppInsightsInstrumentationKey string                                `pulumi:"appInsightsInstrumentationKey"`
	ContainerRegistryCredentials  RegistryListCredentialsResultResponse `pulumi:"containerRegistryCredentials"`
	UserStorageKey                string                                `pulumi:"userStorageKey"`
	UserStorageResourceId         string                                `pulumi:"userStorageResourceId"`
}
