


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2FileDataSetMapping(ctx *pulumi.Context, args *LookupADLSGen2FileDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FileDataSetMappingResult, error) {
	var rv LookupADLSGen2FileDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getADLSGen2FileDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FileDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupADLSGen2FileDataSetMappingResult struct {
	DataSetId            string  `pulumi:"dataSetId"`
	DataSetMappingStatus string  `pulumi:"dataSetMappingStatus"`
	FilePath             string  `pulumi:"filePath"`
	FileSystem           string  `pulumi:"fileSystem"`
	Id                   string  `pulumi:"id"`
	Kind                 string  `pulumi:"kind"`
	Name                 string  `pulumi:"name"`
	OutputType           *string `pulumi:"outputType"`
	ProvisioningState    string  `pulumi:"provisioningState"`
	ResourceGroup        string  `pulumi:"resourceGroup"`
	StorageAccountName   string  `pulumi:"storageAccountName"`
	SubscriptionId       string  `pulumi:"subscriptionId"`
	Type                 string  `pulumi:"type"`
}
