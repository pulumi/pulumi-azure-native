


package datamigration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlMigrationService(ctx *pulumi.Context, args *LookupSqlMigrationServiceArgs, opts ...pulumi.InvokeOption) (*LookupSqlMigrationServiceResult, error) {
	var rv LookupSqlMigrationServiceResult
	err := ctx.Invoke("azure-native:datamigration:getSqlMigrationService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlMigrationServiceArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SqlMigrationServiceName string `pulumi:"sqlMigrationServiceName"`
}


type LookupSqlMigrationServiceResult struct {
	Id                      string             `pulumi:"id"`
	IntegrationRuntimeState string             `pulumi:"integrationRuntimeState"`
	Location                *string            `pulumi:"location"`
	Name                    string             `pulumi:"name"`
	ProvisioningState       string             `pulumi:"provisioningState"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	Tags                    map[string]string  `pulumi:"tags"`
	Type                    string             `pulumi:"type"`
}

func LookupSqlMigrationServiceOutput(ctx *pulumi.Context, args LookupSqlMigrationServiceOutputArgs, opts ...pulumi.InvokeOption) LookupSqlMigrationServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlMigrationServiceResult, error) {
			args := v.(LookupSqlMigrationServiceArgs)
			r, err := LookupSqlMigrationService(ctx, &args, opts...)
			var s LookupSqlMigrationServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlMigrationServiceResultOutput)
}

type LookupSqlMigrationServiceOutputArgs struct {
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlMigrationServiceName pulumi.StringInput `pulumi:"sqlMigrationServiceName"`
}

func (LookupSqlMigrationServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlMigrationServiceArgs)(nil)).Elem()
}


type LookupSqlMigrationServiceResultOutput struct{ *pulumi.OutputState }

func (LookupSqlMigrationServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlMigrationServiceResult)(nil)).Elem()
}

func (o LookupSqlMigrationServiceResultOutput) ToLookupSqlMigrationServiceResultOutput() LookupSqlMigrationServiceResultOutput {
	return o
}

func (o LookupSqlMigrationServiceResultOutput) ToLookupSqlMigrationServiceResultOutputWithContext(ctx context.Context) LookupSqlMigrationServiceResultOutput {
	return o
}

func (o LookupSqlMigrationServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlMigrationServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlMigrationServiceResultOutput) IntegrationRuntimeState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlMigrationServiceResult) string { return v.IntegrationRuntimeState }).(pulumi.StringOutput)
}

func (o LookupSqlMigrationServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlMigrationServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSqlMigrationServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlMigrationServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlMigrationServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlMigrationServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSqlMigrationServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlMigrationServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSqlMigrationServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlMigrationServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlMigrationServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlMigrationServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlMigrationServiceResultOutput{})
}
