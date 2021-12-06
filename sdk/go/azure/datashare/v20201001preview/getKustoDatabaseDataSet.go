


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoDatabaseDataSet(ctx *pulumi.Context, args *LookupKustoDatabaseDataSetArgs, opts ...pulumi.InvokeOption) (*LookupKustoDatabaseDataSetResult, error) {
	var rv LookupKustoDatabaseDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getKustoDatabaseDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoDatabaseDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupKustoDatabaseDataSetResult struct {
	DataSetId               string             `pulumi:"dataSetId"`
	Id                      string             `pulumi:"id"`
	Kind                    string             `pulumi:"kind"`
	KustoDatabaseResourceId string             `pulumi:"kustoDatabaseResourceId"`
	Location                string             `pulumi:"location"`
	Name                    string             `pulumi:"name"`
	ProvisioningState       string             `pulumi:"provisioningState"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	Type                    string             `pulumi:"type"`
}
