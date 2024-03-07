package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/webpubsub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := webpubsub.NewWebPubSub(ctx, "webPubSub", &webpubsub.WebPubSubArgs{
			DisableAadAuth:   pulumi.Bool(false),
			DisableLocalAuth: pulumi.Bool(false),
			Identity: &webpubsub.ManagedIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			LiveTraceConfiguration: &webpubsub.LiveTraceConfigurationArgs{
				Categories: webpubsub.LiveTraceCategoryArray{
					&webpubsub.LiveTraceCategoryArgs{
						Enabled: pulumi.String("true"),
						Name:    pulumi.String("ConnectivityLogs"),
					},
				},
				Enabled: pulumi.String("false"),
			},
			Location: pulumi.String("eastus"),
			NetworkACLs: &webpubsub.WebPubSubNetworkACLsArgs{
				DefaultAction: pulumi.String("Deny"),
				PrivateEndpoints: webpubsub.PrivateEndpointACLArray{
					&webpubsub.PrivateEndpointACLArgs{
						Allow: pulumi.StringArray{
							pulumi.String("ServerConnection"),
						},
						Name: pulumi.String("mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
					},
				},
				PublicNetwork: &webpubsub.NetworkACLArgs{
					Allow: pulumi.StringArray{
						pulumi.String("ClientConnection"),
					},
				},
			},
			PublicNetworkAccess: pulumi.String("Enabled"),
			ResourceGroupName:   pulumi.String("myResourceGroup"),
			ResourceName:        pulumi.String("myWebPubSubService"),
			Sku: &webpubsub.ResourceSkuArgs{
				Capacity: pulumi.Int(1),
				Name:     pulumi.String("Premium_P1"),
				Tier:     pulumi.String("Premium"),
			},
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			Tls: &webpubsub.WebPubSubTlsSettingsArgs{
				ClientCertEnabled: pulumi.Bool(false),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
