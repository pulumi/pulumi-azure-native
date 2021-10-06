


package containerregistry

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBuildStepBuildArguments(ctx *pulumi.Context, args *ListBuildStepBuildArgumentsArgs, opts ...pulumi.InvokeOption) (*ListBuildStepBuildArgumentsResult, error) {
	var rv ListBuildStepBuildArgumentsResult
	err := ctx.Invoke("azure-native:containerregistry:listBuildStepBuildArguments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBuildStepBuildArgumentsArgs struct {
	BuildTaskName     string `pulumi:"buildTaskName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StepName          string `pulumi:"stepName"`
}


type ListBuildStepBuildArgumentsResult struct {
	NextLink *string                 `pulumi:"nextLink"`
	Value    []BuildArgumentResponse `pulumi:"value"`
}
