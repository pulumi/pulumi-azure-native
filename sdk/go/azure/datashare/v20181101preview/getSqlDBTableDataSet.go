


package v20181101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlDBTableDataSet(ctx *pulumi.Context, args *LookupSqlDBTableDataSetArgs, opts ...pulumi.InvokeOption) (*LookupSqlDBTableDataSetResult, error) {
	var rv LookupSqlDBTableDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20181101preview:getSqlDBTableDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlDBTableDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupSqlDBTableDataSetResult struct {
	DataSetId           string `pulumi:"dataSetId"`
	DatabaseName        string `pulumi:"databaseName"`
	Id                  string `pulumi:"id"`
	Kind                string `pulumi:"kind"`
	Name                string `pulumi:"name"`
	SchemaName          string `pulumi:"schemaName"`
	SqlServerResourceId string `pulumi:"sqlServerResourceId"`
	TableName           string `pulumi:"tableName"`
	Type                string `pulumi:"type"`
}
