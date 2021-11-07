


package v20180201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBuildTask(ctx *pulumi.Context, args *LookupBuildTaskArgs, opts ...pulumi.InvokeOption) (*LookupBuildTaskResult, error) {
	var rv LookupBuildTaskResult
	err := ctx.Invoke("azure-native:containerregistry/v20180201preview:getBuildTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBuildTaskArgs struct {
	BuildTaskName     string `pulumi:"buildTaskName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBuildTaskResult struct {
	Alias             string                             `pulumi:"alias"`
	CreationDate      string                             `pulumi:"creationDate"`
	Id                string                             `pulumi:"id"`
	Location          string                             `pulumi:"location"`
	Name              string                             `pulumi:"name"`
	Platform          PlatformPropertiesResponse         `pulumi:"platform"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	SourceRepository  SourceRepositoryPropertiesResponse `pulumi:"sourceRepository"`
	Status            *string                            `pulumi:"status"`
	Tags              map[string]string                  `pulumi:"tags"`
	Timeout           *int                               `pulumi:"timeout"`
	Type              string                             `pulumi:"type"`
}
