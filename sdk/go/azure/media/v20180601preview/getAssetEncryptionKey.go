


package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAssetEncryptionKey(ctx *pulumi.Context, args *GetAssetEncryptionKeyArgs, opts ...pulumi.InvokeOption) (*GetAssetEncryptionKeyResult, error) {
	var rv GetAssetEncryptionKeyResult
	err := ctx.Invoke("azure-native:media/v20180601preview:getAssetEncryptionKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAssetEncryptionKeyArgs struct {
	AccountName       string `pulumi:"accountName"`
	AssetName         string `pulumi:"assetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetAssetEncryptionKeyResult struct {
	StorageEncryptionKey *string `pulumi:"storageEncryptionKey"`
}
