


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlDBTableDataSetMapping(ctx *pulumi.Context, args *LookupSqlDBTableDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupSqlDBTableDataSetMappingResult, error) {
	var rv LookupSqlDBTableDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getSqlDBTableDataSetMapping", args, &rv, opts...)
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
	DataSetId            string `pulumi:"dataSetId"`
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	DatabaseName         string `pulumi:"databaseName"`
	Id                   string `pulumi:"id"`
	Kind                 string `pulumi:"kind"`
	Name                 string `pulumi:"name"`
	ProvisioningState    string `pulumi:"provisioningState"`
	SchemaName           string `pulumi:"schemaName"`
	SqlServerResourceId  string `pulumi:"sqlServerResourceId"`
	TableName            string `pulumi:"tableName"`
	Type                 string `pulumi:"type"`
}

func LookupSqlDBTableDataSetMappingOutput(ctx *pulumi.Context, args LookupSqlDBTableDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupSqlDBTableDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlDBTableDataSetMappingResult, error) {
			args := v.(LookupSqlDBTableDataSetMappingArgs)
			r, err := LookupSqlDBTableDataSetMapping(ctx, &args, opts...)
			var s LookupSqlDBTableDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlDBTableDataSetMappingResultOutput)
}

type LookupSqlDBTableDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupSqlDBTableDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDBTableDataSetMappingArgs)(nil)).Elem()
}


type LookupSqlDBTableDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupSqlDBTableDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDBTableDataSetMappingResult)(nil)).Elem()
}

func (o LookupSqlDBTableDataSetMappingResultOutput) ToLookupSqlDBTableDataSetMappingResultOutput() LookupSqlDBTableDataSetMappingResultOutput {
	return o
}

func (o LookupSqlDBTableDataSetMappingResultOutput) ToLookupSqlDBTableDataSetMappingResultOutputWithContext(ctx context.Context) LookupSqlDBTableDataSetMappingResultOutput {
	return o
}

func (o LookupSqlDBTableDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetMappingResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetMappingResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetMappingResultOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetMappingResult) string { return v.SchemaName }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetMappingResultOutput) SqlServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetMappingResult) string { return v.SqlServerResourceId }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetMappingResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetMappingResult) string { return v.TableName }).(pulumi.StringOutput)
}

func (o LookupSqlDBTableDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDBTableDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlDBTableDataSetMappingResultOutput{})
}
