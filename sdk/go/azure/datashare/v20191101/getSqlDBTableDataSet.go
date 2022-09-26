


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlDBTableDataSet(ctx *pulumi.Context, args *LookupSqlDBTableDataSetArgs, opts ...pulumi.InvokeOption) (*LookupSqlDBTableDataSetResult, error) {
	var rv LookupSqlDBTableDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getSqlDBTableDataSet", args, &rv, opts...)
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

func LookupSqlDBTableDataSetOutput(ctx *pulumi.Context, args LookupSqlDBTableDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupSqlDBTableDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlDBTableDataSetResult, error) {
			args := v.(LookupSqlDBTableDataSetArgs)
			r, err := LookupSqlDBTableDataSet(ctx, &args, opts...)
			var s LookupSqlDBTableDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlDBTableDataSetResultOutput)
}

type LookupSqlDBTableDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupSqlDBTableDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDBTableDataSetArgs)(nil)).Elem()
}


type LookupSqlDBTableDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupSqlDBTableDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDBTableDataSetResult)(nil)).Elem()
}

func (o LookupSqlDBTableDataSetResultOutput) ToLookupSqlDBTableDataSetResultOutput() LookupSqlDBTableDataSetResultOutput {
	return o
}

func (o LookupSqlDBTableDataSetResultOutput) ToLookupSqlDBTableDataSetResultOutputWithContext(ctx context.Context) LookupSqlDBTableDataSetResultOutput {
	return o
}

func (o LookupSqlDBTableDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetResultOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetResult) string { return v.SchemaName }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetResultOutput) SqlServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetResult) string { return v.SqlServerResourceId }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetResult) string { return v.TableName }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlDBTableDataSetResultOutput{})
}
