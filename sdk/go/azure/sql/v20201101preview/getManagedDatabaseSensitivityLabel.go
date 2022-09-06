


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedDatabaseSensitivityLabel(ctx *pulumi.Context, args *LookupManagedDatabaseSensitivityLabelArgs, opts ...pulumi.InvokeOption) (*LookupManagedDatabaseSensitivityLabelResult, error) {
	var rv LookupManagedDatabaseSensitivityLabelResult
	err := ctx.Invoke("azure-native:sql/v20201101preview:getManagedDatabaseSensitivityLabel", args, &rv, opts...)
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

func LookupManagedDatabaseSensitivityLabelOutput(ctx *pulumi.Context, args LookupManagedDatabaseSensitivityLabelOutputArgs, opts ...pulumi.InvokeOption) LookupManagedDatabaseSensitivityLabelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedDatabaseSensitivityLabelResult, error) {
			args := v.(LookupManagedDatabaseSensitivityLabelArgs)
			r, err := LookupManagedDatabaseSensitivityLabel(ctx, &args, opts...)
			var s LookupManagedDatabaseSensitivityLabelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedDatabaseSensitivityLabelResultOutput)
}

type LookupManagedDatabaseSensitivityLabelOutputArgs struct {
	ColumnName             pulumi.StringInput `pulumi:"columnName"`
	DatabaseName           pulumi.StringInput `pulumi:"databaseName"`
	ManagedInstanceName    pulumi.StringInput `pulumi:"managedInstanceName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SchemaName             pulumi.StringInput `pulumi:"schemaName"`
	SensitivityLabelSource pulumi.StringInput `pulumi:"sensitivityLabelSource"`
	TableName              pulumi.StringInput `pulumi:"tableName"`
}

func (LookupManagedDatabaseSensitivityLabelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedDatabaseSensitivityLabelArgs)(nil)).Elem()
}


type LookupManagedDatabaseSensitivityLabelResultOutput struct{ *pulumi.OutputState }

func (LookupManagedDatabaseSensitivityLabelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedDatabaseSensitivityLabelResult)(nil)).Elem()
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) ToLookupManagedDatabaseSensitivityLabelResultOutput() LookupManagedDatabaseSensitivityLabelResultOutput {
	return o
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) ToLookupManagedDatabaseSensitivityLabelResultOutputWithContext(ctx context.Context) LookupManagedDatabaseSensitivityLabelResultOutput {
	return o
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.ColumnName }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) InformationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) *string { return v.InformationType }).(pulumi.StringPtrOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) InformationTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) *string { return v.InformationTypeId }).(pulumi.StringPtrOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) bool { return v.IsDisabled }).(pulumi.BoolOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) LabelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) *string { return v.LabelId }).(pulumi.StringPtrOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) LabelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) *string { return v.LabelName }).(pulumi.StringPtrOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) Rank() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) *string { return v.Rank }).(pulumi.StringPtrOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.SchemaName }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.TableName }).(pulumi.StringOutput)
}

func (o LookupManagedDatabaseSensitivityLabelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseSensitivityLabelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedDatabaseSensitivityLabelResultOutput{})
}
