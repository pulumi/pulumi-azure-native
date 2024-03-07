package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewResourceGuardProxy(ctx, "resourceGuardProxy", &recoveryservices.ResourceGuardProxyArgs{
			Properties: &recoveryservices.ResourceGuardProxyBaseArgs{
				ResourceGuardResourceId: pulumi.String("/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew"),
			},
			ResourceGroupName:      pulumi.String("SampleResourceGroup"),
			ResourceGuardProxyName: pulumi.String("swaggerExample"),
			VaultName:              pulumi.String("sampleVault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
