package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbformariadb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbformariadb.NewServer(ctx, "server", &dbformariadb.ServerArgs{
			Location: pulumi.String("westus"),
			Properties: dbformariadb.ServerPropertiesForReplica{
				CreateMode:     "Replica",
				SourceServerId: "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/MasterResourceGroup/providers/Microsoft.DBforMariaDB/servers/masterserver",
			},
			ResourceGroupName: pulumi.String("TargetResourceGroup"),
			ServerName:        pulumi.String("targetserver"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
