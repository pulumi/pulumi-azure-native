package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := recoveryservices.NewReplicationMigrationItem(ctx, "replicationMigrationItem", &recoveryservices.ReplicationMigrationItemArgs{
FabricName: pulumi.String("vmwarefabric1"),
MigrationItemName: pulumi.String("virtualmachine1"),
Properties: recoveryservices.MigrationItemPropertiesResponse{
PolicyId: pulumi.String("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.RecoveryServices/vaults/migrationvault/replicationPolicies/vmwarepolicy1"),
ProviderSpecificDetails: interface{}{
DataMoverRunAsAccountId: pulumi.String("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.OffAzure/VMwareSites/vmwaresite1/runasaccounts/dataMoverRunAsAccount1"),
DisksToInclude: recoveryservices.VMwareCbtDiskInputArray{
&recoveryservices.VMwareCbtDiskInputArgs{
DiskId: pulumi.String("disk1"),
IsOSDisk: pulumi.String("true"),
LogStorageAccountId: pulumi.String("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.Storage/storageAccounts/logStorageAccount1"),
LogStorageAccountSasSecretName: pulumi.String("logStorageSas"),
},
},
InstanceType: pulumi.String("VMwareCbt"),
SnapshotRunAsAccountId: pulumi.String("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.OffAzure/VMwareSites/vmwaresite1/runasaccounts/snapshotRunAsAccount1"),
TargetNetworkId: pulumi.String("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.Network/virtualNetworks/virtualNetwork1"),
TargetResourceGroupId: pulumi.String("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1"),
VmwareMachineId: pulumi.String("/Subscriptions/cb53d0c3-bd59-4721-89bc-06916a9147ef/resourceGroups/resourcegroup1/providers/Microsoft.OffAzure/VMwareSites/vmwaresite1/machines/virtualmachine1"),
},
},
ProtectionContainerName: pulumi.String("vmwareContainer1"),
ResourceGroupName: pulumi.String("resourcegroup1"),
ResourceName: pulumi.String("migrationvault"),
})
if err != nil {
return err
}
return nil
})
}
