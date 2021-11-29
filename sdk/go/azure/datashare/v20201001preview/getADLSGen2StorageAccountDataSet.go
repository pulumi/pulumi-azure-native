


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupADLSGen2StorageAccountDataSet(ctx *pulumi.Context, args *LookupADLSGen2StorageAccountDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2StorageAccountDataSetResult, error) {
	var rv LookupADLSGen2StorageAccountDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getADLSGen2StorageAccountDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2StorageAccountDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupADLSGen2StorageAccountDataSetResult struct {
	DataSetId                string                               `pulumi:"dataSetId"`
	Id                       string                               `pulumi:"id"`
	Kind                     string                               `pulumi:"kind"`
	Location                 string                               `pulumi:"location"`
	Name                     string                               `pulumi:"name"`
	Paths                    []ADLSGen2StorageAccountPathResponse `pulumi:"paths"`
	StorageAccountResourceId string                               `pulumi:"storageAccountResourceId"`
	SystemData               SystemDataResponse                   `pulumi:"systemData"`
	Type                     string                               `pulumi:"type"`
}
