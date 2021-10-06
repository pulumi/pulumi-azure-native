


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobDataSet(ctx *pulumi.Context, args *LookupBlobDataSetArgs, opts ...pulumi.InvokeOption) (*LookupBlobDataSetResult, error) {
	var rv LookupBlobDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getBlobDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupBlobDataSetResult struct {
	ContainerName      string `pulumi:"containerName"`
	DataSetId          string `pulumi:"dataSetId"`
	FilePath           string `pulumi:"filePath"`
	Id                 string `pulumi:"id"`
	Kind               string `pulumi:"kind"`
	Name               string `pulumi:"name"`
	ResourceGroup      string `pulumi:"resourceGroup"`
	StorageAccountName string `pulumi:"storageAccountName"`
	SubscriptionId     string `pulumi:"subscriptionId"`
	Type               string `pulumi:"type"`
}
