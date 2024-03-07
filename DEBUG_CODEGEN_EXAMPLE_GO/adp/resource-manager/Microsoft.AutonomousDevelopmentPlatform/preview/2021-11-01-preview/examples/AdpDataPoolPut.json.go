package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/autonomousdevelopmentplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := autonomousdevelopmentplatform.NewDataPool(ctx, "dataPool", &autonomousdevelopmentplatform.DataPoolArgs{
			AccountName:  pulumi.String("sampleacct"),
			DataPoolName: pulumi.String("sampledp"),
			Locations: autonomousdevelopmentplatform.DataPoolLocationArray{
				&autonomousdevelopmentplatform.DataPoolLocationArgs{
					Encryption: &autonomousdevelopmentplatform.DataPoolEncryptionArgs{
						KeyName:              pulumi.String("key1"),
						KeyVaultUri:          pulumi.String("https://vaulturi"),
						KeyVersion:           pulumi.String("123"),
						UserAssignedIdentity: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1"),
					},
					Name: pulumi.String("westus"),
				},
			},
			ResourceGroupName: pulumi.String("adpClient"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
