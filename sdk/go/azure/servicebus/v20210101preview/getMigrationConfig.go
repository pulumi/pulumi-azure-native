


package v20210101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMigrationConfig(ctx *pulumi.Context, args *LookupMigrationConfigArgs, opts ...pulumi.InvokeOption) (*LookupMigrationConfigResult, error) {
	var rv LookupMigrationConfigResult
	err := ctx.Invoke("azure-native:servicebus/v20210101preview:getMigrationConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMigrationConfigArgs struct {
	ConfigName        string `pulumi:"configName"`
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMigrationConfigResult struct {
	Id                                string             `pulumi:"id"`
	MigrationState                    string             `pulumi:"migrationState"`
	Name                              string             `pulumi:"name"`
	PendingReplicationOperationsCount float64            `pulumi:"pendingReplicationOperationsCount"`
	PostMigrationName                 string             `pulumi:"postMigrationName"`
	ProvisioningState                 string             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse `pulumi:"systemData"`
	TargetNamespace                   string             `pulumi:"targetNamespace"`
	Type                              string             `pulumi:"type"`
}
