package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewSyncAgent(ctx, "syncAgent", &sql.SyncAgentArgs{
			ResourceGroupName: pulumi.String("syncagentcrud-65440"),
			ServerName:        pulumi.String("syncagentcrud-8475"),
			SyncAgentName:     pulumi.String("syncagentcrud-3187"),
			SyncDatabaseId:    pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-Onebox/providers/Microsoft.Sql/servers/syncagentcrud-8475/databases/sync"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
