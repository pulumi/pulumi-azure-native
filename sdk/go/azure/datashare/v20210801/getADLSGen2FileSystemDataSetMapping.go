


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2FileSystemDataSetMapping(ctx *pulumi.Context, args *LookupADLSGen2FileSystemDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FileSystemDataSetMappingResult, error) {
	var rv LookupADLSGen2FileSystemDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getADLSGen2FileSystemDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FileSystemDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupADLSGen2FileSystemDataSetMappingResult struct {
	DataSetId            string             `pulumi:"dataSetId"`
	DataSetMappingStatus string             `pulumi:"dataSetMappingStatus"`
	FileSystem           string             `pulumi:"fileSystem"`
	Id                   string             `pulumi:"id"`
	Kind                 string             `pulumi:"kind"`
	Name                 string             `pulumi:"name"`
	ProvisioningState    string             `pulumi:"provisioningState"`
	ResourceGroup        string             `pulumi:"resourceGroup"`
	StorageAccountName   string             `pulumi:"storageAccountName"`
	SubscriptionId       string             `pulumi:"subscriptionId"`
	SystemData           SystemDataResponse `pulumi:"systemData"`
	Type                 string             `pulumi:"type"`
}
