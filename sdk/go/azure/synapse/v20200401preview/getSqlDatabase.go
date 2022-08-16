


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSqlDatabase(ctx *pulumi.Context, args *LookupSqlDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupSqlDatabaseResult, error) {
	var rv LookupSqlDatabaseResult
	err := ctx.Invoke("azure-native:synapse/v20200401preview:getSqlDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlDatabaseArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SqlDatabaseName   string `pulumi:"sqlDatabaseName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupSqlDatabaseResult struct {
	Collation         *string                           `pulumi:"collation"`
	DataRetention     *SqlDatabaseDataRetentionResponse `pulumi:"dataRetention"`
	DatabaseGuid      string                            `pulumi:"databaseGuid"`
	Id                string                            `pulumi:"id"`
	Location          string                            `pulumi:"location"`
	Name              string                            `pulumi:"name"`
	Status            string                            `pulumi:"status"`
	StorageRedundancy *string                           `pulumi:"storageRedundancy"`
	SystemData        SystemDataResponse                `pulumi:"systemData"`
	Tags              map[string]string                 `pulumi:"tags"`
	Type              string                            `pulumi:"type"`
}

func LookupSqlDatabaseOutput(ctx *pulumi.Context, args LookupSqlDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupSqlDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlDatabaseResult, error) {
			args := v.(LookupSqlDatabaseArgs)
			r, err := LookupSqlDatabase(ctx, &args, opts...)
			var s LookupSqlDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlDatabaseResultOutput)
}

type LookupSqlDatabaseOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlDatabaseName   pulumi.StringInput `pulumi:"sqlDatabaseName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDatabaseArgs)(nil)).Elem()
}


type LookupSqlDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupSqlDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlDatabaseResult)(nil)).Elem()
}

func (o LookupSqlDatabaseResultOutput) ToLookupSqlDatabaseResultOutput() LookupSqlDatabaseResultOutput {
	return o
}

func (o LookupSqlDatabaseResultOutput) ToLookupSqlDatabaseResultOutputWithContext(ctx context.Context) LookupSqlDatabaseResultOutput {
	return o
}

func (o LookupSqlDatabaseResultOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlDatabaseResult) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o LookupSqlDatabaseResultOutput) DataRetention() SqlDatabaseDataRetentionResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlDatabaseResult) *SqlDatabaseDataRetentionResponse { return v.DataRetention }).(SqlDatabaseDataRetentionResponsePtrOutput)
}

func (o LookupSqlDatabaseResultOutput) DatabaseGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDatabaseResult) string { return v.DatabaseGuid }).(pulumi.StringOutput)
}

func (o LookupSqlDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlDatabaseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDatabaseResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSqlDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlDatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDatabaseResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSqlDatabaseResultOutput) StorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlDatabaseResult) *string { return v.StorageRedundancy }).(pulumi.StringPtrOutput)
}

func (o LookupSqlDatabaseResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlDatabaseResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSqlDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlDatabaseResultOutput{})
}
