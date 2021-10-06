


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlDBTableDataSetMapping(ctx *pulumi.Context, args *LookupSqlDBTableDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupSqlDBTableDataSetMappingResult, error) {
	var rv LookupSqlDBTableDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getSqlDBTableDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlDBTableDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupSqlDBTableDataSetMappingResult struct {
	DataSetId            string             `pulumi:"dataSetId"`
	DataSetMappingStatus string             `pulumi:"dataSetMappingStatus"`
	DatabaseName         string             `pulumi:"databaseName"`
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
