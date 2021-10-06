


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAsset(ctx *pulumi.Context, args *LookupAssetArgs, opts ...pulumi.InvokeOption) (*LookupAssetResult, error) {
	var rv LookupAssetResult
	err := ctx.Invoke("azure-native:media/v20210601:getAsset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssetArgs struct {
	AccountName       string `pulumi:"accountName"`
	AssetName         string `pulumi:"assetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAssetResult struct {
	AlternateId             *string            `pulumi:"alternateId"`
	AssetId                 string             `pulumi:"assetId"`
	Container               *string            `pulumi:"container"`
	Created                 string             `pulumi:"created"`
	Description             *string            `pulumi:"description"`
	Id                      string             `pulumi:"id"`
	LastModified            string             `pulumi:"lastModified"`
	Name                    string             `pulumi:"name"`
	StorageAccountName      *string            `pulumi:"storageAccountName"`
	StorageEncryptionFormat string             `pulumi:"storageEncryptionFormat"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	Type                    string             `pulumi:"type"`
}
