package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewVault(ctx, "vault", &recoveryservices.VaultArgs{
			Identity: &recoveryservices.IdentityDataArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Location: pulumi.String("West US"),
			Properties: &recoveryservices.VaultPropertiesArgs{
				PublicNetworkAccess: pulumi.String("Enabled"),
			},
			ResourceGroupName: pulumi.String("Default-RecoveryServices-ResourceGroup"),
			Sku: &recoveryservices.SkuArgs{
				Name: pulumi.String("Standard"),
			},
			VaultName: pulumi.String("swaggerExample"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
