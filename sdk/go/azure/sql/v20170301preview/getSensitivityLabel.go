


package v20170301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSensitivityLabel(ctx *pulumi.Context, args *LookupSensitivityLabelArgs, opts ...pulumi.InvokeOption) (*LookupSensitivityLabelResult, error) {
	var rv LookupSensitivityLabelResult
	err := ctx.Invoke("azure-native:sql/v20170301preview:getSensitivityLabel", args, &rv, opts...)
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
	Id                string  `pulumi:"id"`
	InformationType   *string `pulumi:"informationType"`
	InformationTypeId *string `pulumi:"informationTypeId"`
	IsDisabled        bool    `pulumi:"isDisabled"`
	LabelId           *string `pulumi:"labelId"`
	LabelName         *string `pulumi:"labelName"`
	Name              string  `pulumi:"name"`
	Rank              *string `pulumi:"rank"`
	Type              string  `pulumi:"type"`
}
