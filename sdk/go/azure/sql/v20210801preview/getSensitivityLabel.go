


package v20210801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSensitivityLabel(ctx *pulumi.Context, args *LookupSensitivityLabelArgs, opts ...pulumi.InvokeOption) (*LookupSensitivityLabelResult, error) {
	var rv LookupSensitivityLabelResult
	err := ctx.Invoke("azure-native:sql/v20210801preview:getSensitivityLabel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSensitivityLabelArgs struct {
	ColumnName             string `pulumi:"columnName"`
	DatabaseName           string `pulumi:"databaseName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SchemaName             string `pulumi:"schemaName"`
	SensitivityLabelSource string `pulumi:"sensitivityLabelSource"`
	ServerName             string `pulumi:"serverName"`
	TableName              string `pulumi:"tableName"`
}


type LookupSensitivityLabelResult struct {
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
