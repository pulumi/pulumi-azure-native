


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlDWTableDataSetMapping(ctx *pulumi.Context, args *LookupSqlDWTableDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupSqlDWTableDataSetMappingResult, error) {
	var rv LookupSqlDWTableDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getSqlDWTableDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlDWTableDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupSqlDWTableDataSetMappingResult struct {
	DataSetId            string             `pulumi:"dataSetId"`
	DataSetMappingStatus string             `pulumi:"dataSetMappingStatus"`
	DataWarehouseName    string             `pulumi:"dataWarehouseName"`
	Id                   string             `pulumi:"id"`
	Kind                 string             `pulumi:"kind"`
	Name                 string             `pulumi:"name"`
	ProvisioningState    string             `pulumi:"provisioningState"`
	SchemaName           string             `pulumi:"schemaName"`
	SqlServerResourceId  string             `pulumi:"sqlServerResourceId"`
	SystemData           SystemDataResponse `pulumi:"systemData"`
	TableName            string             `pulumi:"tableName"`
	Type                 string             `pulumi:"type"`
}

func LookupSqlDWTableDataSetMappingOutput(ctx *pulumi.Context, args LookupSqlDWTableDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupSqlDWTableDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlDWTableDataSetMappingResult, error) {
			args := v.(LookupSqlDWTableDataSetMappingArgs)
			r, err := LookupSqlDWTableDataSetMapping(ctx, &args, opts...)
			var s LookupSqlDWTableDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlDWTableDataSetMappingResultOutput)
}

type LookupSqlDWTableDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupSqlDWTableDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDWTableDataSetMappingArgs)(nil)).Elem()
}


type LookupSqlDWTableDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupSqlDWTableDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDWTableDataSetMappingResult)(nil)).Elem()
}

func (o LookupSqlDWTableDataSetMappingResultOutput) ToLookupSqlDWTableDataSetMappingResultOutput() LookupSqlDWTableDataSetMappingResultOutput {
	return o
}

func (o LookupSqlDWTableDataSetMappingResultOutput) ToLookupSqlDWTableDataSetMappingResultOutputWithContext(ctx context.Context) LookupSqlDWTableDataSetMappingResultOutput {
	return o
}

func (o LookupSqlDWTableDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetMappingResultOutput) DataWarehouseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.DataWarehouseName }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetMappingResultOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.SchemaName }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetMappingResultOutput) SqlServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.SqlServerResourceId }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSqlDWTableDataSetMappingResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.TableName }).(pulumi.StringOutput)
}

func (o LookupSqlDWTableDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDWTableDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlDWTableDataSetMappingResultOutput{})
}
