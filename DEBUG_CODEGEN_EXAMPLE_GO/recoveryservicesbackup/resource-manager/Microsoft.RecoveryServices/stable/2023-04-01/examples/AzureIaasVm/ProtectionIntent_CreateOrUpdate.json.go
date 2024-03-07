package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewProtectionIntent(ctx, "protectionIntent", &recoveryservices.ProtectionIntentArgs{
			FabricName:       pulumi.String("Azure"),
			IntentObjectName: pulumi.String("vm;iaasvmcontainerv2;chamsrgtest;chamscandel"),
			Properties: recoveryservices.AzureResourceProtectionIntent{
				PolicyId:                 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupPolicies/myPolicy",
				ProtectionIntentItemType: "AzureResourceItem",
				SourceResourceId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/chamsrgtest/providers/Microsoft.Compute/virtualMachines/chamscandel",
			},
			ResourceGroupName: pulumi.String("myRG"),
			VaultName:         pulumi.String("myVault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
