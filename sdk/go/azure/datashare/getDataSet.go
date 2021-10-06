


package datashare

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDataSet(ctx *pulumi.Context, args *LookupDataSetArgs, opts ...pulumi.InvokeOption) (*LookupDataSetResult, error) {
	var rv LookupDataSetResult
	err := ctx.Invoke("azure-native:datashare:getDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupDataSetResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}
