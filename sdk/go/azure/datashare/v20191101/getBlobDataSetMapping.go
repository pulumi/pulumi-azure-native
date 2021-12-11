


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobDataSetMapping(ctx *pulumi.Context, args *LookupBlobDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupBlobDataSetMappingResult, error) {
	var rv LookupBlobDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getBlobDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupBlobDataSetMappingResult struct {
	ContainerName        string  `pulumi:"containerName"`
	DataSetId            string  `pulumi:"dataSetId"`
	DataSetMappingStatus string  `pulumi:"dataSetMappingStatus"`
	FilePath             string  `pulumi:"filePath"`
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
