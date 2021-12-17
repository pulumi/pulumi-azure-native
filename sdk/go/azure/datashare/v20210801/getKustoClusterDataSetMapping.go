


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoClusterDataSetMapping(ctx *pulumi.Context, args *LookupKustoClusterDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupKustoClusterDataSetMappingResult, error) {
	var rv LookupKustoClusterDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getKustoClusterDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoClusterDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupKustoClusterDataSetMappingResult struct {
	DataSetId              string             `pulumi:"dataSetId"`
	DataSetMappingStatus   string             `pulumi:"dataSetMappingStatus"`
	Id                     string             `pulumi:"id"`
	Kind                   string             `pulumi:"kind"`
	KustoClusterResourceId string             `pulumi:"kustoClusterResourceId"`
	Location               string             `pulumi:"location"`
	Name                   string             `pulumi:"name"`
	ProvisioningState      string             `pulumi:"provisioningState"`
	SystemData             SystemDataResponse `pulumi:"systemData"`
	Type                   string             `pulumi:"type"`
}
