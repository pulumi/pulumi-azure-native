


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAssetStreamingLocators(ctx *pulumi.Context, args *ListAssetStreamingLocatorsArgs, opts ...pulumi.InvokeOption) (*ListAssetStreamingLocatorsResult, error) {
	var rv ListAssetStreamingLocatorsResult
	err := ctx.Invoke("azure-native:media/v20210601:listAssetStreamingLocators", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAssetStreamingLocatorsArgs struct {
	AccountName       string `pulumi:"accountName"`
	AssetName         string `pulumi:"assetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAssetStreamingLocatorsResult struct {
	StreamingLocators []AssetStreamingLocatorResponse `pulumi:"streamingLocators"`
}
