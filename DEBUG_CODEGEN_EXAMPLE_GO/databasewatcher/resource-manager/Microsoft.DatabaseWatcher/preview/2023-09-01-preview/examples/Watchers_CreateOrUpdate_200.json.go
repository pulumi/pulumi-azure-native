package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databasewatcher/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databasewatcher.NewWatcher(ctx, "watcher", &databasewatcher.WatcherArgs{
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("apiTest-ddat4p"),
			WatcherName:       pulumi.String("databasemo3ej9ih"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
