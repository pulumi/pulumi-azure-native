


package datashare

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobFolderDataSet(ctx *pulumi.Context, args *LookupBlobFolderDataSetArgs, opts ...pulumi.InvokeOption) (*LookupBlobFolderDataSetResult, error) {
	var rv LookupBlobFolderDataSetResult
	err := ctx.Invoke("azure-native:datashare:getBlobFolderDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobFolderDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupBlobFolderDataSetResult struct {
	ContainerName      string             `pulumi:"containerName"`
	DataSetId          string             `pulumi:"dataSetId"`
	Id                 string             `pulumi:"id"`
	Kind               string             `pulumi:"kind"`
	Name               string             `pulumi:"name"`
	Prefix             string             `pulumi:"prefix"`
	ResourceGroup      string             `pulumi:"resourceGroup"`
	StorageAccountName string             `pulumi:"storageAccountName"`
	SubscriptionId     string             `pulumi:"subscriptionId"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
}
