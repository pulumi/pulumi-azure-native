


package datamigration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSqlMigrationServiceMonitoringData(ctx *pulumi.Context, args *ListSqlMigrationServiceMonitoringDataArgs, opts ...pulumi.InvokeOption) (*ListSqlMigrationServiceMonitoringDataResult, error) {
	var rv ListSqlMigrationServiceMonitoringDataResult
	err := ctx.Invoke("azure-native:datamigration:listSqlMigrationServiceMonitoringData", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSqlMigrationServiceMonitoringDataArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SqlMigrationServiceName string `pulumi:"sqlMigrationServiceName"`
}


type ListSqlMigrationServiceMonitoringDataResult struct {
	Name  string                       `pulumi:"name"`
	Nodes []NodeMonitoringDataResponse `pulumi:"nodes"`
}

func ListSqlMigrationServiceMonitoringDataOutput(ctx *pulumi.Context, args ListSqlMigrationServiceMonitoringDataOutputArgs, opts ...pulumi.InvokeOption) ListSqlMigrationServiceMonitoringDataResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSqlMigrationServiceMonitoringDataResult, error) {
			args := v.(ListSqlMigrationServiceMonitoringDataArgs)
			r, err := ListSqlMigrationServiceMonitoringData(ctx, &args, opts...)
			var s ListSqlMigrationServiceMonitoringDataResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSqlMigrationServiceMonitoringDataResultOutput)
}

type ListSqlMigrationServiceMonitoringDataOutputArgs struct {
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlMigrationServiceName pulumi.StringInput `pulumi:"sqlMigrationServiceName"`
}

func (ListSqlMigrationServiceMonitoringDataOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSqlMigrationServiceMonitoringDataArgs)(nil)).Elem()
}


type ListSqlMigrationServiceMonitoringDataResultOutput struct{ *pulumi.OutputState }

func (ListSqlMigrationServiceMonitoringDataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSqlMigrationServiceMonitoringDataResult)(nil)).Elem()
}

func (o ListSqlMigrationServiceMonitoringDataResultOutput) ToListSqlMigrationServiceMonitoringDataResultOutput() ListSqlMigrationServiceMonitoringDataResultOutput {
	return o
}

func (o ListSqlMigrationServiceMonitoringDataResultOutput) ToListSqlMigrationServiceMonitoringDataResultOutputWithContext(ctx context.Context) ListSqlMigrationServiceMonitoringDataResultOutput {
	return o
}

func (o ListSqlMigrationServiceMonitoringDataResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListSqlMigrationServiceMonitoringDataResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListSqlMigrationServiceMonitoringDataResultOutput) Nodes() NodeMonitoringDataResponseArrayOutput {
	return o.ApplyT(func(v ListSqlMigrationServiceMonitoringDataResult) []NodeMonitoringDataResponse { return v.Nodes }).(NodeMonitoringDataResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSqlMigrationServiceMonitoringDataResultOutput{})
}
