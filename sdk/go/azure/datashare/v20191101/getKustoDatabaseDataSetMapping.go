


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoDatabaseDataSetMapping(ctx *pulumi.Context, args *LookupKustoDatabaseDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupKustoDatabaseDataSetMappingResult, error) {
	var rv LookupKustoDatabaseDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getKustoDatabaseDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoDatabaseDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupKustoDatabaseDataSetMappingResult struct {
	DataSetId              string `pulumi:"dataSetId"`
	DataSetMappingStatus   string `pulumi:"dataSetMappingStatus"`
	Id                     string `pulumi:"id"`
	Kind                   string `pulumi:"kind"`
	KustoClusterResourceId string `pulumi:"kustoClusterResourceId"`
	Location               string `pulumi:"location"`
	Name                   string `pulumi:"name"`
	ProvisioningState      string `pulumi:"provisioningState"`
	Type                   string `pulumi:"type"`
}
