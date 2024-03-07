package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewSyncMember(ctx, "syncMember", &sql.SyncMemberArgs{
			DatabaseName:                      pulumi.String("syncgroupcrud-7421"),
			DatabaseType:                      pulumi.String("AzureSqlDatabase"),
			ResourceGroupName:                 pulumi.String("syncgroupcrud-65440"),
			ServerName:                        pulumi.String("syncgroupcrud-8475"),
			SyncDirection:                     pulumi.String("Bidirectional"),
			SyncGroupName:                     pulumi.String("syncgroupcrud-3187"),
			SyncMemberAzureDatabaseResourceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-65440/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
			SyncMemberName:                    pulumi.String("syncmembercrud-4879"),
			UsePrivateLinkConnection:          pulumi.Bool(true),
			UserName:                          pulumi.String("myUser"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
