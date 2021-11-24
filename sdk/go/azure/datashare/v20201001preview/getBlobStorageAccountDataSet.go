


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobStorageAccountDataSet(ctx *pulumi.Context, args *LookupBlobStorageAccountDataSetArgs, opts ...pulumi.InvokeOption) (*LookupBlobStorageAccountDataSetResult, error) {
	var rv LookupBlobStorageAccountDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getBlobStorageAccountDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobStorageAccountDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupBlobStorageAccountDataSetResult struct {
	DataSetId                string                           `pulumi:"dataSetId"`
	Id                       string                           `pulumi:"id"`
	Kind                     string                           `pulumi:"kind"`
	Location                 string                           `pulumi:"location"`
	Name                     string                           `pulumi:"name"`
	Paths                    []BlobStorageAccountPathResponse `pulumi:"paths"`
	StorageAccountResourceId string                           `pulumi:"storageAccountResourceId"`
	SystemData               SystemDataResponse               `pulumi:"systemData"`
	Type                     string                           `pulumi:"type"`
}
