


package v20220615preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlServerDatabase(ctx *pulumi.Context, args *LookupSqlServerDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerDatabaseResult, error) {
	var rv LookupSqlServerDatabaseResult
	err := ctx.Invoke("azure-native:azurearcdata/v20220615preview:getSqlServerDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlServerDatabaseArgs struct {
	DatabaseName          string `pulumi:"databaseName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SqlServerInstanceName string `pulumi:"sqlServerInstanceName"`
}


type LookupSqlServerDatabaseResult struct {
	Id         string                                      `pulumi:"id"`
	Location   string                                      `pulumi:"location"`
	Name       string                                      `pulumi:"name"`
	Properties SqlServerDatabaseResourcePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                          `pulumi:"systemData"`
	Tags       map[string]string                           `pulumi:"tags"`
	Type       string                                      `pulumi:"type"`
}

func LookupSqlServerDatabaseOutput(ctx *pulumi.Context, args LookupSqlServerDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupSqlServerDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlServerDatabaseResult, error) {
			args := v.(LookupSqlServerDatabaseArgs)
			r, err := LookupSqlServerDatabase(ctx, &args, opts...)
			var s LookupSqlServerDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlServerDatabaseResultOutput)
}

type LookupSqlServerDatabaseOutputArgs struct {
	DatabaseName          pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlServerInstanceName pulumi.StringInput `pulumi:"sqlServerInstanceName"`
}

func (LookupSqlServerDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerDatabaseArgs)(nil)).Elem()
}


type LookupSqlServerDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupSqlServerDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerDatabaseResult)(nil)).Elem()
}

func (o LookupSqlServerDatabaseResultOutput) ToLookupSqlServerDatabaseResultOutput() LookupSqlServerDatabaseResultOutput {
	return o
}

func (o LookupSqlServerDatabaseResultOutput) ToLookupSqlServerDatabaseResultOutputWithContext(ctx context.Context) LookupSqlServerDatabaseResultOutput {
	return o
}

func (o LookupSqlServerDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlServerDatabaseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSqlServerDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlServerDatabaseResultOutput) Properties() SqlServerDatabaseResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) SqlServerDatabaseResourcePropertiesResponse { return v.Properties }).(SqlServerDatabaseResourcePropertiesResponseOutput)
}

func (o LookupSqlServerDatabaseResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSqlServerDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlServerDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlServerDatabaseResultOutput{})
}
