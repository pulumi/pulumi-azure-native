


package datamigration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseMigrationsSqlDb(ctx *pulumi.Context, args *LookupDatabaseMigrationsSqlDbArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseMigrationsSqlDbResult, error) {
	var rv LookupDatabaseMigrationsSqlDbResult
	err := ctx.Invoke("azure-native:datamigration:getDatabaseMigrationsSqlDb", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseMigrationsSqlDbArgs struct {
	Expand               *string `pulumi:"expand"`
	MigrationOperationId *string `pulumi:"migrationOperationId"`
	ResourceGroupName    string  `pulumi:"resourceGroupName"`
	SqlDbInstanceName    string  `pulumi:"sqlDbInstanceName"`
	TargetDbName         string  `pulumi:"targetDbName"`
}


type LookupDatabaseMigrationsSqlDbResult struct {
	Id         string                                   `pulumi:"id"`
	Name       string                                   `pulumi:"name"`
	Properties DatabaseMigrationPropertiesSqlDbResponse `pulumi:"properties"`
	SystemData SystemDataResponse                       `pulumi:"systemData"`
	Type       string                                   `pulumi:"type"`
}

func LookupDatabaseMigrationsSqlDbOutput(ctx *pulumi.Context, args LookupDatabaseMigrationsSqlDbOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseMigrationsSqlDbResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseMigrationsSqlDbResult, error) {
			args := v.(LookupDatabaseMigrationsSqlDbArgs)
			r, err := LookupDatabaseMigrationsSqlDb(ctx, &args, opts...)
			var s LookupDatabaseMigrationsSqlDbResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseMigrationsSqlDbResultOutput)
}

type LookupDatabaseMigrationsSqlDbOutputArgs struct {
	Expand               pulumi.StringPtrInput `pulumi:"expand"`
	MigrationOperationId pulumi.StringPtrInput `pulumi:"migrationOperationId"`
	ResourceGroupName    pulumi.StringInput    `pulumi:"resourceGroupName"`
	SqlDbInstanceName    pulumi.StringInput    `pulumi:"sqlDbInstanceName"`
	TargetDbName         pulumi.StringInput    `pulumi:"targetDbName"`
}

func (LookupDatabaseMigrationsSqlDbOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseMigrationsSqlDbArgs)(nil)).Elem()
}


type LookupDatabaseMigrationsSqlDbResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseMigrationsSqlDbResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseMigrationsSqlDbResult)(nil)).Elem()
}

func (o LookupDatabaseMigrationsSqlDbResultOutput) ToLookupDatabaseMigrationsSqlDbResultOutput() LookupDatabaseMigrationsSqlDbResultOutput {
	return o
}

func (o LookupDatabaseMigrationsSqlDbResultOutput) ToLookupDatabaseMigrationsSqlDbResultOutputWithContext(ctx context.Context) LookupDatabaseMigrationsSqlDbResultOutput {
	return o
}

func (o LookupDatabaseMigrationsSqlDbResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseMigrationsSqlDbResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseMigrationsSqlDbResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseMigrationsSqlDbResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseMigrationsSqlDbResultOutput) Properties() DatabaseMigrationPropertiesSqlDbResponseOutput {
	return o.ApplyT(func(v LookupDatabaseMigrationsSqlDbResult) DatabaseMigrationPropertiesSqlDbResponse {
		return v.Properties
	}).(DatabaseMigrationPropertiesSqlDbResponseOutput)
}

func (o LookupDatabaseMigrationsSqlDbResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDatabaseMigrationsSqlDbResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDatabaseMigrationsSqlDbResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseMigrationsSqlDbResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseMigrationsSqlDbResultOutput{})
}
