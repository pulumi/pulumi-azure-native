


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlDWTableDataSet(ctx *pulumi.Context, args *LookupSqlDWTableDataSetArgs, opts ...pulumi.InvokeOption) (*LookupSqlDWTableDataSetResult, error) {
	var rv LookupSqlDWTableDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getSqlDWTableDataSet", args, &rv, opts...)
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

func LookupSqlDWTableDataSetOutput(ctx *pulumi.Context, args LookupSqlDWTableDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupSqlDWTableDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlDWTableDataSetResult, error) {
			args := v.(LookupSqlDWTableDataSetArgs)
			r, err := LookupSqlDWTableDataSet(ctx, &args, opts...)
			var s LookupSqlDWTableDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlDWTableDataSetResultOutput)
}

type LookupSqlDWTableDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupSqlDWTableDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDWTableDataSetArgs)(nil)).Elem()
}


type LookupSqlDWTableDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupSqlDWTableDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDWTableDataSetResult)(nil)).Elem()
}

func (o LookupSqlDWTableDataSetResultOutput) ToLookupSqlDWTableDataSetResultOutput() LookupSqlDWTableDataSetResultOutput {
	return o
}

func (o LookupSqlDWTableDataSetResultOutput) ToLookupSqlDWTableDataSetResultOutputWithContext(ctx context.Context) LookupSqlDWTableDataSetResultOutput {
	return o
}

func (o LookupSqlDWTableDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetResultOutput) DataWarehouseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.DataWarehouseName }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetResultOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.SchemaName }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetResultOutput) SqlServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.SqlServerResourceId }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSqlDWTableDataSetResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.TableName }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlDWTableDataSetResultOutput{})
}
