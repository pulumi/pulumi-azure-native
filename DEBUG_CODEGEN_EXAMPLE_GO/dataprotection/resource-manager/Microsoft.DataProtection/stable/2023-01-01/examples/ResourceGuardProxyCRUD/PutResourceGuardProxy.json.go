package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dataprotection/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dataprotection.NewDppResourceGuardProxy(ctx, "dppResourceGuardProxy", &dataprotection.DppResourceGuardProxyArgs{
			Properties: &dataprotection.ResourceGuardProxyBaseArgs{
				ResourceGuardResourceId: pulumi.String("/subscriptions/f9e67185-f313-4e79-aa71-6458d429369d/resourceGroups/ResourceGuardSecurityAdminRG/providers/Microsoft.DataProtection/resourceGuards/ResourceGuardTestResource"),
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
