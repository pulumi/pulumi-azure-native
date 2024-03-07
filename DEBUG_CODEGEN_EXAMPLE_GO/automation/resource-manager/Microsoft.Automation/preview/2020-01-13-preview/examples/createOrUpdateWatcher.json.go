package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewWatcher(ctx, "watcher", &automation.WatcherArgs{
			AutomationAccountName:       pulumi.String("MyTestAutomationAccount"),
			Description:                 pulumi.String("This is a test watcher."),
			ExecutionFrequencyInSeconds: pulumi.Float64(60),
			ResourceGroupName:           pulumi.String("rg"),
			ScriptName:                  pulumi.String("MyTestWatcherRunbook"),
			ScriptRunOn:                 pulumi.String("MyTestHybridWorkerGroup"),
			Tags:                        nil,
			WatcherName:                 pulumi.String("MyTestWatcher"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
