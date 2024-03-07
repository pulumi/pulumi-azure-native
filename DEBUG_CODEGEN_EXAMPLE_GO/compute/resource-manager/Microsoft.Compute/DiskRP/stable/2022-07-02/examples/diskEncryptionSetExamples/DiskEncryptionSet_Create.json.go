package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewDiskEncryptionSet(ctx, "diskEncryptionSet", &compute.DiskEncryptionSetArgs{
			ActiveKey: compute.KeyForDiskEncryptionSetResponse{
				KeyUrl: pulumi.String("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
				SourceVault: &compute.SourceVaultArgs{
					Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
				},
			},
			DiskEncryptionSetName: pulumi.String("myDiskEncryptionSet"),
			EncryptionType:        pulumi.String("EncryptionAtRestWithCustomerKey"),
			Identity: &compute.EncryptionSetIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
