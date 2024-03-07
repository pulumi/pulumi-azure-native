package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewSyncGroup(ctx, "syncGroup", &sql.SyncGroupArgs{
			ConflictResolutionPolicy: pulumi.String("HubWin"),
			DatabaseName:             pulumi.String("syncgroupcrud-4328"),
			HubDatabaseUserName:      pulumi.String("hubUser"),
			Interval:                 -1,
			ResourceGroupName:        pulumi.String("syncgroupcrud-65440"),
			ServerName:               pulumi.String("syncgroupcrud-8475"),
			SyncDatabaseId:           pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
			SyncGroupName:            pulumi.String("syncgroupcrud-3187"),
			UsePrivateLinkConnection: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
