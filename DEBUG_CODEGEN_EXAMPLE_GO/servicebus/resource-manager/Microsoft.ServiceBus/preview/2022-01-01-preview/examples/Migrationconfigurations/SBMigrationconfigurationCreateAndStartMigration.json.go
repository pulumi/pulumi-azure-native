package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicebus.NewMigrationConfig(ctx, "migrationConfig", &servicebus.MigrationConfigArgs{
			ConfigName:        pulumi.String("$default"),
			NamespaceName:     pulumi.String("sdk-Namespace-41"),
			PostMigrationName: pulumi.String("sdk-PostMigration-5919"),
			ResourceGroupName: pulumi.String("ResourceGroup"),
			TargetNamespace:   pulumi.String("/subscriptions/SubscriptionId/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-4028"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
