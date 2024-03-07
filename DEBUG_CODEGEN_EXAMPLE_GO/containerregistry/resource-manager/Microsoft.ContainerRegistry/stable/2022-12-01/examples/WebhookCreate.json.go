package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewWebhook(ctx, "webhook", &containerregistry.WebhookArgs{
			Actions: pulumi.StringArray{
				pulumi.String("push"),
			},
			CustomHeaders: pulumi.StringMap{
				"Authorization": pulumi.String("******"),
			},
			Location:          pulumi.String("westus"),
			RegistryName:      pulumi.String("myRegistry"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Scope:             pulumi.String("myRepository"),
			ServiceUri:        pulumi.String("http://myservice.com"),
			Status:            pulumi.String("enabled"),
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
			WebhookName: pulumi.String("myWebhook"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
