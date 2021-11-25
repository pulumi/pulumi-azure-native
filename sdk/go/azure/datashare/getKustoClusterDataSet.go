


package datashare

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoClusterDataSet(ctx *pulumi.Context, args *LookupKustoClusterDataSetArgs, opts ...pulumi.InvokeOption) (*LookupKustoClusterDataSetResult, error) {
	var rv LookupKustoClusterDataSetResult
	err := ctx.Invoke("azure-native:datashare:getKustoClusterDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoClusterDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupKustoClusterDataSetResult struct {
	DataSetId              string             `pulumi:"dataSetId"`
	Id                     string             `pulumi:"id"`
	Kind                   string             `pulumi:"kind"`
	KustoClusterResourceId string             `pulumi:"kustoClusterResourceId"`
	Location               string             `pulumi:"location"`
	Name                   string             `pulumi:"name"`
	ProvisioningState      string             `pulumi:"provisioningState"`
	SystemData             SystemDataResponse `pulumi:"systemData"`
	Type                   string             `pulumi:"type"`
}
