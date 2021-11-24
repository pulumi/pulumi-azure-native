


package v20180901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRunLogSasUrl(ctx *pulumi.Context, args *ListRunLogSasUrlArgs, opts ...pulumi.InvokeOption) (*ListRunLogSasUrlResult, error) {
	var rv ListRunLogSasUrlResult
	err := ctx.Invoke("azure-native:containerregistry/v20180901:listRunLogSasUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRunLogSasUrlArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RunId             string `pulumi:"runId"`
}


type ListRunLogSasUrlResult struct {
	LogLink *string `pulumi:"logLink"`
}
