


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlPoolSensitivityLabel(ctx *pulumi.Context, args *LookupSqlPoolSensitivityLabelArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolSensitivityLabelResult, error) {
	var rv LookupSqlPoolSensitivityLabelResult
	err := ctx.Invoke("azure-native:synapse/v20201201:getSqlPoolSensitivityLabel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolSensitivityLabelArgs struct {
	ColumnName             string `pulumi:"columnName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SchemaName             string `pulumi:"schemaName"`
	SensitivityLabelSource string `pulumi:"sensitivityLabelSource"`
	SqlPoolName            string `pulumi:"sqlPoolName"`
	TableName              string `pulumi:"tableName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type LookupSqlPoolSensitivityLabelResult struct {
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
