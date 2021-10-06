


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSynapseWorkspaceSqlPoolTableDataSet(ctx *pulumi.Context, args *LookupSynapseWorkspaceSqlPoolTableDataSetArgs, opts ...pulumi.InvokeOption) (*LookupSynapseWorkspaceSqlPoolTableDataSetResult, error) {
	var rv LookupSynapseWorkspaceSqlPoolTableDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20200901:getSynapseWorkspaceSqlPoolTableDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSynapseWorkspaceSqlPoolTableDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupSynapseWorkspaceSqlPoolTableDataSetResult struct {
	DataSetId                              string             `pulumi:"dataSetId"`
	Id                                     string             `pulumi:"id"`
	Kind                                   string             `pulumi:"kind"`
	Name                                   string             `pulumi:"name"`
	SynapseWorkspaceSqlPoolTableResourceId string             `pulumi:"synapseWorkspaceSqlPoolTableResourceId"`
	SystemData                             SystemDataResponse `pulumi:"systemData"`
	Type                                   string             `pulumi:"type"`
}
