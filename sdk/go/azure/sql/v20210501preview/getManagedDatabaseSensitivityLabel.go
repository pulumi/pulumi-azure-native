


package v20210501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedDatabaseSensitivityLabel(ctx *pulumi.Context, args *LookupManagedDatabaseSensitivityLabelArgs, opts ...pulumi.InvokeOption) (*LookupManagedDatabaseSensitivityLabelResult, error) {
	var rv LookupManagedDatabaseSensitivityLabelResult
	err := ctx.Invoke("azure-native:sql/v20210501preview:getManagedDatabaseSensitivityLabel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedDatabaseSensitivityLabelArgs struct {
	ColumnName             string `pulumi:"columnName"`
	DatabaseName           string `pulumi:"databaseName"`
	ManagedInstanceName    string `pulumi:"managedInstanceName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SchemaName             string `pulumi:"schemaName"`
	SensitivityLabelSource string `pulumi:"sensitivityLabelSource"`
	TableName              string `pulumi:"tableName"`
}


type LookupManagedDatabaseSensitivityLabelResult struct {
	ColumnName        string  `pulumi:"columnName"`
	Id                string  `pulumi:"id"`
	InformationType   *string `pulumi:"informationType"`
	InformationTypeId *string `pulumi:"informationTypeId"`
	IsDisabled        bool    `pulumi:"isDisabled"`
	LabelId           *string `pulumi:"labelId"`
	LabelName         *string `pulumi:"labelName"`
	ManagedBy         string  `pulumi:"managedBy"`
	Name              string  `pulumi:"name"`
	Rank              *string `pulumi:"rank"`
	SchemaName        string  `pulumi:"schemaName"`
	TableName         string  `pulumi:"tableName"`
	Type              string  `pulumi:"type"`
}
