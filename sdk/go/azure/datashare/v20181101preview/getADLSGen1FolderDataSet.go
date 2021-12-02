


package v20181101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen1FolderDataSet(ctx *pulumi.Context, args *LookupADLSGen1FolderDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen1FolderDataSetResult, error) {
	var rv LookupADLSGen1FolderDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20181101preview:getADLSGen1FolderDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen1FolderDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupADLSGen1FolderDataSetResult struct {
	AccountName    string `pulumi:"accountName"`
	DataSetId      string `pulumi:"dataSetId"`
	FolderPath     string `pulumi:"folderPath"`
	Id             string `pulumi:"id"`
	Kind           string `pulumi:"kind"`
	Name           string `pulumi:"name"`
	ResourceGroup  string `pulumi:"resourceGroup"`
	SubscriptionId string `pulumi:"subscriptionId"`
	Type           string `pulumi:"type"`
}
