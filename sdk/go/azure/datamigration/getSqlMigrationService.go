


package datamigration

import (
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
