package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databricks/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := databricks.NewWorkspace(ctx, "workspace", &databricks.WorkspaceArgs{
Location: pulumi.String("westus"),
ManagedResourceGroupId: pulumi.String("/subscriptions/subid/resourceGroups/myManagedRG"),
Parameters: databricks.WorkspaceCustomParametersResponse{
Encryption: interface{}{
Value: &databricks.EncryptionArgs{
KeyName: pulumi.String("myKeyName"),
KeySource: pulumi.String("Microsoft.Keyvault"),
KeyVaultUri: pulumi.String("https://myKeyVault.vault.azure.net/"),
KeyVersion: pulumi.String("00000000000000000000000000000000"),
},
},
PrepareEncryption: &databricks.WorkspaceCustomBooleanParameterArgs{
Value: pulumi.Bool(true),
},
},
ResourceGroupName: pulumi.String("rg"),
WorkspaceName: pulumi.String("myWorkspace"),
})
if err != nil {
return err
}
return nil
})
}
