package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cognitiveservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := cognitiveservices.NewAccount(ctx, "account", &cognitiveservices.AccountArgs{
AccountName: pulumi.String("testCreate1"),
Identity: &cognitiveservices.IdentityArgs{
Type: cognitiveservices.ResourceIdentityTypeSystemAssigned,
},
Kind: pulumi.String("Emotion"),
Location: pulumi.String("West US"),
Properties: cognitiveservices.AccountPropertiesResponse{
Encryption: interface{}{
KeySource: pulumi.String("Microsoft.KeyVault"),
KeyVaultProperties: &cognitiveservices.KeyVaultPropertiesArgs{
KeyName: pulumi.String("KeyName"),
KeyVaultUri: pulumi.String("https://pltfrmscrts-use-pc-dev.vault.azure.net/"),
KeyVersion: pulumi.String("891CF236-D241-4738-9462-D506AF493DFA"),
},
},
UserOwnedStorage: cognitiveservices.UserOwnedStorageArray{
&cognitiveservices.UserOwnedStorageArgs{
ResourceId: pulumi.String("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
},
},
},
ResourceGroupName: pulumi.String("myResourceGroup"),
Sku: &cognitiveservices.SkuArgs{
Name: pulumi.String("S0"),
},
})
if err != nil {
return err
}
return nil
})
}
