package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := storage.NewStorageAccount(ctx, "storageAccount", &storage.StorageAccountArgs{
AccountName: pulumi.String("sto4445"),
AllowBlobPublicAccess: pulumi.Bool(false),
AllowSharedKeyAccess: pulumi.Bool(true),
DefaultToOAuthAuthentication: pulumi.Bool(false),
Encryption: storage.EncryptionResponse{
KeySource: pulumi.String("Microsoft.Storage"),
RequireInfrastructureEncryption: pulumi.Bool(false),
Services: interface{}{
Blob: &storage.EncryptionServiceArgs{
Enabled: pulumi.Bool(true),
KeyType: pulumi.String("Account"),
},
File: &storage.EncryptionServiceArgs{
Enabled: pulumi.Bool(true),
KeyType: pulumi.String("Account"),
},
},
},
ExtendedLocation: &storage.ExtendedLocationArgs{
Name: pulumi.String("losangeles001"),
Type: pulumi.String("EdgeZone"),
},
IsHnsEnabled: pulumi.Bool(true),
IsSftpEnabled: pulumi.Bool(true),
KeyPolicy: &storage.KeyPolicyArgs{
KeyExpirationPeriodInDays: pulumi.Int(20),
},
Kind: pulumi.String("Storage"),
Location: pulumi.String("eastus"),
MinimumTlsVersion: pulumi.String("TLS1_2"),
ResourceGroupName: pulumi.String("res9101"),
RoutingPreference: &storage.RoutingPreferenceArgs{
PublishInternetEndpoints: pulumi.Bool(true),
PublishMicrosoftEndpoints: pulumi.Bool(true),
RoutingChoice: pulumi.String("MicrosoftRouting"),
},
SasPolicy: &storage.SasPolicyArgs{
ExpirationAction: pulumi.String("Log"),
SasExpirationPeriod: pulumi.String("1.15:59:59"),
},
Sku: &storage.SkuArgs{
Name: pulumi.String("Standard_GRS"),
},
Tags: pulumi.StringMap{
"key1": pulumi.String("value1"),
"key2": pulumi.String("value2"),
},
})
if err != nil {
return err
}
return nil
})
}
