package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databricks/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := databricks.NewWorkspace(ctx, "workspace", &databricks.WorkspaceArgs{
Encryption: databricks.WorkspacePropertiesResponseEncryption{
Entities: interface{}{
ManagedDisk: interface{}{
KeySource: pulumi.String("Microsoft.Keyvault"),
KeyVaultProperties: &databricks.ManagedDiskEncryptionKeyVaultPropertiesArgs{
KeyName: pulumi.String("test-cmk-key"),
KeyVaultUri: pulumi.String("https://test-vault-name.vault.azure.net/"),
KeyVersion: pulumi.String("00000000000000000000000000000000"),
},
RotationToLatestKeyVersionEnabled: pulumi.Bool(true),
},
},
},
Location: pulumi.String("westus"),
ManagedResourceGroupId: pulumi.String("/subscriptions/subid/resourceGroups/myManagedRG"),
ResourceGroupName: pulumi.String("rg"),
Tags: pulumi.StringMap{
"mytag1": pulumi.String("myvalue1"),
},
WorkspaceName: pulumi.String("myWorkspace"),
})
if err != nil {
return err
}
return nil
})
}
