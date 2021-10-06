


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationMigrationItem(ctx *pulumi.Context, args *LookupReplicationMigrationItemArgs, opts ...pulumi.InvokeOption) (*LookupReplicationMigrationItemResult, error) {
	var rv LookupReplicationMigrationItemResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210301:getReplicationMigrationItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationMigrationItemArgs struct {
	FabricName              string `pulumi:"fabricName"`
	MigrationItemName       string `pulumi:"migrationItemName"`
	ProtectionContainerName string `pulumi:"protectionContainerName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	ResourceName            string `pulumi:"resourceName"`
}


type LookupReplicationMigrationItemResult struct {
	Id         string                          `pulumi:"id"`
	Location   *string                         `pulumi:"location"`
	Name       string                          `pulumi:"name"`
	Properties MigrationItemPropertiesResponse `pulumi:"properties"`
	Type       string                          `pulumi:"type"`
}
