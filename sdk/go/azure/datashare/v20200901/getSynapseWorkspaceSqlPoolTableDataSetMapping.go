


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSynapseWorkspaceSqlPoolTableDataSetMapping(ctx *pulumi.Context, args *LookupSynapseWorkspaceSqlPoolTableDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult, error) {
	var rv LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20200901:getSynapseWorkspaceSqlPoolTableDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSynapseWorkspaceSqlPoolTableDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult struct {
	DataSetId                              string             `pulumi:"dataSetId"`
	DataSetMappingStatus                   string             `pulumi:"dataSetMappingStatus"`
	Id                                     string             `pulumi:"id"`
	Kind                                   string             `pulumi:"kind"`
	Name                                   string             `pulumi:"name"`
	ProvisioningState                      string             `pulumi:"provisioningState"`
	SynapseWorkspaceSqlPoolTableResourceId string             `pulumi:"synapseWorkspaceSqlPoolTableResourceId"`
	SystemData                             SystemDataResponse `pulumi:"systemData"`
	Type                                   string             `pulumi:"type"`
}
