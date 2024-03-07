package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logic.NewIntegrationServiceEnvironment(ctx, "integrationServiceEnvironment", &logic.IntegrationServiceEnvironmentArgs{
			IntegrationServiceEnvironmentName: pulumi.String("testIntegrationServiceEnvironment"),
			Location:                          pulumi.String("brazilsouth"),
			Properties: &logic.IntegrationServiceEnvironmentPropertiesArgs{
				EncryptionConfiguration: &logic.IntegrationServiceEnvironmenEncryptionConfigurationArgs{
					EncryptionKeyReference: &logic.IntegrationServiceEnvironmenEncryptionKeyReferenceArgs{
						KeyName: pulumi.String("testKeyName"),
						KeyVault: &logic.ResourceReferenceArgs{
							Id: pulumi.String("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.KeyVault/vaults/testKeyVault"),
						},
						KeyVersion: pulumi.String("13b261d30b984753869902d7f47f4d55"),
					},
				},
				NetworkConfiguration: &logic.NetworkConfigurationArgs{
					AccessEndpoint: &logic.IntegrationServiceEnvironmentAccessEndpointArgs{
						Type: pulumi.String("Internal"),
					},
					Subnets: logic.ResourceReferenceArray{
						&logic.ResourceReferenceArgs{
							Id: pulumi.String("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s1"),
						},
						&logic.ResourceReferenceArgs{
							Id: pulumi.String("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s2"),
						},
						&logic.ResourceReferenceArgs{
							Id: pulumi.String("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s3"),
						},
						&logic.ResourceReferenceArgs{
							Id: pulumi.String("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s4"),
						},
					},
				},
			},
			ResourceGroup: pulumi.String("testResourceGroup"),
			Sku: &logic.IntegrationServiceEnvironmentSkuArgs{
				Capacity: pulumi.Int(2),
				Name:     pulumi.String("Premium"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
