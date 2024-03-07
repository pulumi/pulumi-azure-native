package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dataprotection/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := dataprotection.NewBackupInstance(ctx, "backupInstance", &dataprotection.BackupInstanceArgs{
BackupInstanceName: pulumi.String("testInstance1"),
Properties: &dataprotection.BackupInstanceTypeArgs{
DataSourceInfo: &dataprotection.DatasourceArgs{
DatasourceType: pulumi.String("Microsoft.DBforPostgreSQL/servers/databases"),
ObjectType: pulumi.String("Datasource"),
ResourceID: pulumi.String("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
ResourceLocation: pulumi.String(""),
ResourceName: pulumi.String("testdb"),
ResourceType: pulumi.String("Microsoft.DBforPostgreSQL/servers/databases"),
ResourceUri: pulumi.String(""),
},
DataSourceSetInfo: &dataprotection.DatasourceSetArgs{
DatasourceType: pulumi.String("Microsoft.DBforPostgreSQL/servers/databases"),
ObjectType: pulumi.String("DatasourceSet"),
ResourceID: pulumi.String("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
ResourceLocation: pulumi.String(""),
ResourceName: pulumi.String("viveksipgtest"),
ResourceType: pulumi.String("Microsoft.DBforPostgreSQL/servers"),
ResourceUri: pulumi.String(""),
},
DatasourceAuthCredentials: interface{}{
ObjectType: pulumi.String("SecretStoreBasedAuthCredentials"),
SecretStoreResource: &dataprotection.SecretStoreResourceArgs{
SecretStoreType: pulumi.String("AzureKeyVault"),
Uri: pulumi.String("https://samplevault.vault.azure.net/secrets/credentials"),
},
},
FriendlyName: pulumi.String("harshitbi2"),
ObjectType: pulumi.String("BackupInstance"),
PolicyInfo: &dataprotection.PolicyInfoArgs{
PolicyId: pulumi.String("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/Backupvaults/PratikPrivatePreviewVault1/backupPolicies/PratikPolicy1"),
PolicyParameters: &dataprotection.PolicyParametersArgs{
DataStoreParametersList: []dataprotection.AzureOperationalStoreParametersArgs{
{
DataStoreType: pulumi.String("OperationalStore"),
ObjectType: pulumi.String("AzureOperationalStoreParameters"),
ResourceGroupId: pulumi.String("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest"),
},
},
},
},
ValidationType: pulumi.String("ShallowValidation"),
},
ResourceGroupName: pulumi.String("000pikumar"),
Tags: pulumi.StringMap{
"key1": pulumi.String("val1"),
},
VaultName: pulumi.String("PratikPrivatePreviewVault1"),
})
if err != nil {
return err
}
return nil
})
}
