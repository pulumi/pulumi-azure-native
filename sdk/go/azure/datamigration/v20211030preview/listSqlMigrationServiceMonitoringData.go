


package v20211030preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSqlMigrationServiceMonitoringData(ctx *pulumi.Context, args *ListSqlMigrationServiceMonitoringDataArgs, opts ...pulumi.InvokeOption) (*ListSqlMigrationServiceMonitoringDataResult, error) {
	var rv ListSqlMigrationServiceMonitoringDataResult
	err := ctx.Invoke("azure-native:datamigration/v20211030preview:listSqlMigrationServiceMonitoringData", args, &rv, opts...)
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
