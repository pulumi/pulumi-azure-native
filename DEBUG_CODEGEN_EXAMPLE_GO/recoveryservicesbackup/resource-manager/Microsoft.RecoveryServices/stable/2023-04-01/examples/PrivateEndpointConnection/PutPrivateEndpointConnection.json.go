package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &recoveryservices.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b"),
			Properties: recoveryservices.PrivateEndpointConnectionResponse{
				GroupIds: pulumi.StringArray{
					pulumi.String("AzureBackup_secondary"),
				},
				PrivateEndpoint: &recoveryservices.PrivateEndpointArgs{
					Id: pulumi.String("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/gaallaRG/providers/Microsoft.Network/privateEndpoints/gaallatestpe3"),
				},
				PrivateLinkServiceConnectionState: &recoveryservices.PrivateLinkServiceConnectionStateArgs{
					Description: pulumi.String("Approved by johndoe@company.com"),
					Status:      pulumi.String("Approved"),
				},
				ProvisioningState: pulumi.String("Succeeded"),
			},
			ResourceGroupName: pulumi.String("gaallaRG"),
			VaultName:         pulumi.String("gaallavaultbvtd2msi"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
