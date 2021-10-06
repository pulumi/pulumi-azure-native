


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlDWTableDataSetMapping(ctx *pulumi.Context, args *LookupSqlDWTableDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupSqlDWTableDataSetMappingResult, error) {
	var rv LookupSqlDWTableDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getSqlDWTableDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlDWTableDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupSqlDWTableDataSetMappingResult struct {
	DataSetId            string             `pulumi:"dataSetId"`
	DataSetMappingStatus string             `pulumi:"dataSetMappingStatus"`
	DataWarehouseName    string             `pulumi:"dataWarehouseName"`
	Id                   string             `pulumi:"id"`
	Kind                 string             `pulumi:"kind"`
	Name                 string             `pulumi:"name"`
	ProvisioningState    string             `pulumi:"provisioningState"`
	SchemaName           string             `pulumi:"schemaName"`
	SqlServerResourceId  string             `pulumi:"sqlServerResourceId"`
	SystemData           SystemDataResponse `pulumi:"systemData"`
	TableName            string             `pulumi:"tableName"`
	Type                 string             `pulumi:"type"`
}
