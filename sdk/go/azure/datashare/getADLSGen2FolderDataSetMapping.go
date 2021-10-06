


package datashare

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2FolderDataSetMapping(ctx *pulumi.Context, args *LookupADLSGen2FolderDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FolderDataSetMappingResult, error) {
	var rv LookupADLSGen2FolderDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare:getADLSGen2FolderDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FolderDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupADLSGen2FolderDataSetMappingResult struct {
	DataSetId            string             `pulumi:"dataSetId"`
	DataSetMappingStatus string             `pulumi:"dataSetMappingStatus"`
	FileSystem           string             `pulumi:"fileSystem"`
	FolderPath           string             `pulumi:"folderPath"`
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
