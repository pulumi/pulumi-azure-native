


package datashare

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlDWTableDataSet(ctx *pulumi.Context, args *LookupSqlDWTableDataSetArgs, opts ...pulumi.InvokeOption) (*LookupSqlDWTableDataSetResult, error) {
	var rv LookupSqlDWTableDataSetResult
	err := ctx.Invoke("azure-native:datashare:getSqlDWTableDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlDWTableDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupSqlDWTableDataSetResult struct {
	DataSetId           string             `pulumi:"dataSetId"`
	DataWarehouseName   string             `pulumi:"dataWarehouseName"`
	Id                  string             `pulumi:"id"`
	Kind                string             `pulumi:"kind"`
	Name                string             `pulumi:"name"`
	SchemaName          string             `pulumi:"schemaName"`
	SqlServerResourceId string             `pulumi:"sqlServerResourceId"`
	SystemData          SystemDataResponse `pulumi:"systemData"`
	TableName           string             `pulumi:"tableName"`
	Type                string             `pulumi:"type"`
}
