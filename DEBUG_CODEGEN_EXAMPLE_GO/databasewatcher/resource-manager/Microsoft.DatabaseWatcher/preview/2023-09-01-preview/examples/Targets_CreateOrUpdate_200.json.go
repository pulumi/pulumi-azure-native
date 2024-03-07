package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databasewatcher/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databasewatcher.NewTarget(ctx, "target", &databasewatcher.TargetArgs{
			ConnectionServerName:     pulumi.String("sqlServero1ihe2"),
			ResourceGroupName:        pulumi.String("apiTest-ddat4p"),
			TargetAuthenticationType: pulumi.String("Aad"),
			TargetName:               pulumi.String("monitoringh22eed"),
			TargetType:               pulumi.String("SqlDb"),
			WatcherName:              pulumi.String("databasemo3ej9ih"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
