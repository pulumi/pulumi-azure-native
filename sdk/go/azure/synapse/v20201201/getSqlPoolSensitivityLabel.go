


package v20201201

import (
	"context"
	"reflect"

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

func LookupSqlPoolSensitivityLabelOutput(ctx *pulumi.Context, args LookupSqlPoolSensitivityLabelOutputArgs, opts ...pulumi.InvokeOption) LookupSqlPoolSensitivityLabelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlPoolSensitivityLabelResult, error) {
			args := v.(LookupSqlPoolSensitivityLabelArgs)
			r, err := LookupSqlPoolSensitivityLabel(ctx, &args, opts...)
			var s LookupSqlPoolSensitivityLabelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlPoolSensitivityLabelResultOutput)
}

type LookupSqlPoolSensitivityLabelOutputArgs struct {
	ColumnName             pulumi.StringInput `pulumi:"columnName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SchemaName             pulumi.StringInput `pulumi:"schemaName"`
	SensitivityLabelSource pulumi.StringInput `pulumi:"sensitivityLabelSource"`
	SqlPoolName            pulumi.StringInput `pulumi:"sqlPoolName"`
	TableName              pulumi.StringInput `pulumi:"tableName"`
	WorkspaceName          pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlPoolSensitivityLabelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolSensitivityLabelArgs)(nil)).Elem()
}


type LookupSqlPoolSensitivityLabelResultOutput struct{ *pulumi.OutputState }

func (LookupSqlPoolSensitivityLabelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolSensitivityLabelResult)(nil)).Elem()
}

func (o LookupSqlPoolSensitivityLabelResultOutput) ToLookupSqlPoolSensitivityLabelResultOutput() LookupSqlPoolSensitivityLabelResultOutput {
	return o
}

func (o LookupSqlPoolSensitivityLabelResultOutput) ToLookupSqlPoolSensitivityLabelResultOutputWithContext(ctx context.Context) LookupSqlPoolSensitivityLabelResultOutput {
	return o
}

func (o LookupSqlPoolSensitivityLabelResultOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.ColumnName }).(pulumi.StringOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) InformationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) *string { return v.InformationType }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) InformationTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) *string { return v.InformationTypeId }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) bool { return v.IsDisabled }).(pulumi.BoolOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) LabelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) *string { return v.LabelId }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) LabelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) *string { return v.LabelName }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) Rank() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) *string { return v.Rank }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.SchemaName }).(pulumi.StringOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.TableName }).(pulumi.StringOutput)
}

func (o LookupSqlPoolSensitivityLabelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolSensitivityLabelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlPoolSensitivityLabelResultOutput{})
}
