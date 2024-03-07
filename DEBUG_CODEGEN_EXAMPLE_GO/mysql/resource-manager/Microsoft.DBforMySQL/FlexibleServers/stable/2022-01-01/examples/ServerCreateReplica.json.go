package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbformysql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbformysql.NewServer(ctx, "server", &dbformysql.ServerArgs{
			CreateMode:             pulumi.String("Replica"),
			Location:               pulumi.String("SoutheastAsia"),
			ResourceGroupName:      pulumi.String("testgr"),
			ServerName:             pulumi.String("replica-server"),
			SourceServerResourceId: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testgr/providers/Microsoft.DBforMySQL/flexibleServers/source-server"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
