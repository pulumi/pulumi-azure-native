package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestack/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestack.NewLinkedSubscription(ctx, "linkedSubscription", &azurestack.LinkedSubscriptionArgs{
			LinkedSubscriptionId:   pulumi.String("104fbb77-2b0e-476a-83de-65ad8acd1f0b"),
			LinkedSubscriptionName: pulumi.String("testLinkedSubscription"),
			Location:               pulumi.String("eastus"),
			RegistrationResourceId: pulumi.String("/subscriptions/dd8597b4-8739-4467-8b10-f8679f62bfbf/resourceGroups/azurestack/providers/Microsoft.AzureStack/registrations/testRegistration"),
			ResourceGroup:          pulumi.String("azurestack"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
