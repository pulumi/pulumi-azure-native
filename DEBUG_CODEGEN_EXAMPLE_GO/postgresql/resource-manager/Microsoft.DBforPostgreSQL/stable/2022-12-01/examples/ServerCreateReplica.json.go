package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewServer(ctx, "server", &dbforpostgresql.ServerArgs{
			CreateMode:             pulumi.String("Replica"),
			Location:               pulumi.String("westus"),
			PointInTimeUTC:         pulumi.String("2021-06-27T00:04:59.4078005+00:00"),
			ResourceGroupName:      pulumi.String("testrg"),
			ServerName:             pulumi.String("pgtestsvc5rep"),
			SourceServerResourceId: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/sourcepgservername"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
