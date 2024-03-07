package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewDaprSubscription(ctx, "daprSubscription", &app.DaprSubscriptionArgs{
			BulkSubscribe: &app.DaprSubscriptionBulkSubscribeOptionsArgs{
				Enabled:            pulumi.Bool(true),
				MaxAwaitDurationMs: pulumi.Int(500),
				MaxMessagesCount:   pulumi.Int(123),
			},
			EnvironmentName:   pulumi.String("myenvironment"),
			Name:              pulumi.String("mysubscription"),
			PubsubName:        pulumi.String("mypubsubcomponent"),
			ResourceGroupName: pulumi.String("examplerg"),
			Routes: &app.DaprSubscriptionRoutesArgs{
				Default: pulumi.String("/products"),
			},
			Scopes: pulumi.StringArray{
				pulumi.String("warehouseapp"),
				pulumi.String("customersupportapp"),
			},
			Topic: pulumi.String("inventory"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
