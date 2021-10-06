


package datashare

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2FolderDataSet(ctx *pulumi.Context, args *LookupADLSGen2FolderDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FolderDataSetResult, error) {
	var rv LookupADLSGen2FolderDataSetResult
	err := ctx.Invoke("azure-native:datashare:getADLSGen2FolderDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FolderDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupADLSGen2FolderDataSetResult struct {
	DataSetId          string             `pulumi:"dataSetId"`
	FileSystem         string             `pulumi:"fileSystem"`
	FolderPath         string             `pulumi:"folderPath"`
	Id                 string             `pulumi:"id"`
	Kind               string             `pulumi:"kind"`
	Name               string             `pulumi:"name"`
	ResourceGroup      string             `pulumi:"resourceGroup"`
	StorageAccountName string             `pulumi:"storageAccountName"`
	SubscriptionId     string             `pulumi:"subscriptionId"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
}
