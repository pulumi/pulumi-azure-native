


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSensitivityLabel(ctx *pulumi.Context, args *LookupSensitivityLabelArgs, opts ...pulumi.InvokeOption) (*LookupSensitivityLabelResult, error) {
	var rv LookupSensitivityLabelResult
	err := ctx.Invoke("azure-native:sql/v20220201preview:getSensitivityLabel", args, &rv, opts...)
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

func LookupSensitivityLabelOutput(ctx *pulumi.Context, args LookupSensitivityLabelOutputArgs, opts ...pulumi.InvokeOption) LookupSensitivityLabelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSensitivityLabelResult, error) {
			args := v.(LookupSensitivityLabelArgs)
			r, err := LookupSensitivityLabel(ctx, &args, opts...)
			var s LookupSensitivityLabelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSensitivityLabelResultOutput)
}

type LookupSensitivityLabelOutputArgs struct {
	ColumnName             pulumi.StringInput `pulumi:"columnName"`
	DatabaseName           pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SchemaName             pulumi.StringInput `pulumi:"schemaName"`
	SensitivityLabelSource pulumi.StringInput `pulumi:"sensitivityLabelSource"`
	ServerName             pulumi.StringInput `pulumi:"serverName"`
	TableName              pulumi.StringInput `pulumi:"tableName"`
}

func (LookupSensitivityLabelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSensitivityLabelArgs)(nil)).Elem()
}


type LookupSensitivityLabelResultOutput struct{ *pulumi.OutputState }

func (LookupSensitivityLabelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSensitivityLabelResult)(nil)).Elem()
}

func (o LookupSensitivityLabelResultOutput) ToLookupSensitivityLabelResultOutput() LookupSensitivityLabelResultOutput {
	return o
}

func (o LookupSensitivityLabelResultOutput) ToLookupSensitivityLabelResultOutputWithContext(ctx context.Context) LookupSensitivityLabelResultOutput {
	return o
}

func (o LookupSensitivityLabelResultOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) string { return v.ColumnName }).(pulumi.StringOutput)
}

func (o LookupSensitivityLabelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSensitivityLabelResultOutput) InformationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) *string { return v.InformationType }).(pulumi.StringPtrOutput)
}

func (o LookupSensitivityLabelResultOutput) InformationTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) *string { return v.InformationTypeId }).(pulumi.StringPtrOutput)
}

func (o LookupSensitivityLabelResultOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) bool { return v.IsDisabled }).(pulumi.BoolOutput)
}

func (o LookupSensitivityLabelResultOutput) LabelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) *string { return v.LabelId }).(pulumi.StringPtrOutput)
}

func (o LookupSensitivityLabelResultOutput) LabelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) *string { return v.LabelName }).(pulumi.StringPtrOutput)
}

func (o LookupSensitivityLabelResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o LookupSensitivityLabelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSensitivityLabelResultOutput) Rank() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) *string { return v.Rank }).(pulumi.StringPtrOutput)
}

func (o LookupSensitivityLabelResultOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) string { return v.SchemaName }).(pulumi.StringOutput)
}

func (o LookupSensitivityLabelResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) string { return v.TableName }).(pulumi.StringOutput)
}

func (o LookupSensitivityLabelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSensitivityLabelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSensitivityLabelResultOutput{})
}
