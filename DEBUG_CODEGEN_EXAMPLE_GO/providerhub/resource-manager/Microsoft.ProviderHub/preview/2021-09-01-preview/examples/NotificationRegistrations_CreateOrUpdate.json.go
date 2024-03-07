package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/providerhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := providerhub.NewNotificationRegistration(ctx, "notificationRegistration", &providerhub.NotificationRegistrationArgs{
			NotificationRegistrationName: pulumi.String("fooNotificationRegistration"),
			Properties: providerhub.NotificationRegistrationResponseProperties{
				IncludedEvents: pulumi.StringArray{
					pulumi.String("*/write"),
					pulumi.String("Microsoft.Contoso/employees/delete"),
				},
				MessageScope: pulumi.String("RegisteredSubscriptions"),
				NotificationEndpoints: providerhub.NotificationEndpointArray{
					&providerhub.NotificationEndpointArgs{
						Locations: pulumi.StringArray{
							pulumi.String(""),
							pulumi.String("East US"),
						},
						NotificationDestination: pulumi.String("/subscriptions/ac6bcfb5-3dc1-491f-95a6-646b89bf3e88/resourceGroups/mgmtexp-eastus/providers/Microsoft.EventHub/namespaces/unitedstates-mgmtexpint/eventhubs/armlinkednotifications"),
					},
					&providerhub.NotificationEndpointArgs{
						Locations: pulumi.StringArray{
							pulumi.String("North Europe"),
						},
						NotificationDestination: pulumi.String("/subscriptions/ac6bcfb5-3dc1-491f-95a6-646b89bf3e88/resourceGroups/mgmtexp-northeurope/providers/Microsoft.EventHub/namespaces/europe-mgmtexpint/eventhubs/armlinkednotifications"),
					},
				},
				NotificationMode: pulumi.String("EventHub"),
			},
			ProviderNamespace: pulumi.String("Microsoft.Contoso"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
