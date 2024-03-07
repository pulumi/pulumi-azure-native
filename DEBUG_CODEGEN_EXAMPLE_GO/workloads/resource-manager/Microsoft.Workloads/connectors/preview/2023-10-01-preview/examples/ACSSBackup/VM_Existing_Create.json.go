package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewACSSBackupConnection(ctx, "acssBackupConnection", &workloads.ACSSBackupConnectionArgs{
			BackupData: workloads.VMBackupData{
				BackupPolicy: workloads.VMBackupPolicyProperties{
					Name: "defaultVmPolicy",
				},
				BackupType: "VM",
				DiskExclusionProperties: workloads.DiskExclusionProperties{
					DiskLunList:     []interface{}{},
					IsInclusionList: true,
				},
				RecoveryServicesVault: workloads.ExistingRecoveryServicesVault{
					Id:        "/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.RecoveryServices/vaults/test-vault",
					VaultType: "Existing",
				},
			},
			BackupName:        pulumi.String("vmBackup"),
			ConnectorName:     pulumi.String("C1"),
			Location:          pulumi.String("westcentralus"),
			ResourceGroupName: pulumi.String("test-rg"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
