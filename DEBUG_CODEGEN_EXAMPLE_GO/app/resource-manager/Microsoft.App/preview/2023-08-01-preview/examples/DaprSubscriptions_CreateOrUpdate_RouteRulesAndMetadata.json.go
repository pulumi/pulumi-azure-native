package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewDaprSubscription(ctx, "daprSubscription", &app.DaprSubscriptionArgs{
			EnvironmentName: pulumi.String("myenvironment"),
			Metadata: pulumi.StringMap{
				"foo":   pulumi.String("bar"),
				"hello": pulumi.String("world"),
			},
			Name:              pulumi.String("mysubscription"),
			PubsubName:        pulumi.String("mypubsubcomponent"),
			ResourceGroupName: pulumi.String("examplerg"),
			Routes: app.DaprSubscriptionRoutesResponse{
				Default: pulumi.String("/products"),
				Rules: app.DaprSubscriptionRouteRuleArray{
					&app.DaprSubscriptionRouteRuleArgs{
						Match: pulumi.String("event.type == 'widget'"),
						Path:  pulumi.String("/widgets"),
					},
					&app.DaprSubscriptionRouteRuleArgs{
						Match: pulumi.String("event.type == 'gadget'"),
						Path:  pulumi.String("/gadgets"),
					},
				},
			},
			Topic: pulumi.String("inventory"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
