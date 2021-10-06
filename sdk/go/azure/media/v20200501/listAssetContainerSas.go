


package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAssetContainerSas(ctx *pulumi.Context, args *ListAssetContainerSasArgs, opts ...pulumi.InvokeOption) (*ListAssetContainerSasResult, error) {
	var rv ListAssetContainerSasResult
	err := ctx.Invoke("azure-native:media/v20200501:listAssetContainerSas", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAssetContainerSasArgs struct {
	AccountName       string  `pulumi:"accountName"`
	AssetName         string  `pulumi:"assetName"`
	ExpiryTime        *string `pulumi:"expiryTime"`
	Permissions       *string `pulumi:"permissions"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type ListAssetContainerSasResult struct {
	AssetContainerSasUrls []string `pulumi:"assetContainerSasUrls"`
}
