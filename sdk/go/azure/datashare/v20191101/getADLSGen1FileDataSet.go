


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen1FileDataSet(ctx *pulumi.Context, args *LookupADLSGen1FileDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen1FileDataSetResult, error) {
	var rv LookupADLSGen1FileDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getADLSGen1FileDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen1FileDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupADLSGen1FileDataSetResult struct {
	AccountName    string `pulumi:"accountName"`
	DataSetId      string `pulumi:"dataSetId"`
	FileName       string `pulumi:"fileName"`
	FolderPath     string `pulumi:"folderPath"`
	Id             string `pulumi:"id"`
	Kind           string `pulumi:"kind"`
	Name           string `pulumi:"name"`
	ResourceGroup  string `pulumi:"resourceGroup"`
	SubscriptionId string `pulumi:"subscriptionId"`
	Type           string `pulumi:"type"`
}
