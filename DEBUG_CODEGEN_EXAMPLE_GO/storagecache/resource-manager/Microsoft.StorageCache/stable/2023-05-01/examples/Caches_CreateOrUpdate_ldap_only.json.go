package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagecache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := storagecache.NewCache(ctx, "cache", &storagecache.CacheArgs{
CacheName: pulumi.String("sc1"),
CacheSizeGB: pulumi.Int(3072),
DirectoryServicesSettings: storagecache.CacheDirectorySettingsResponse{
UsernameDownload: interface{}{
Credentials: &storagecache.CacheUsernameDownloadSettingsCredentialsArgs{
BindDn: pulumi.String("cn=ldapadmin,dc=contosoad,dc=contoso,dc=local"),
BindPassword: pulumi.String("<bindPassword>"),
},
ExtendedGroups: pulumi.Bool(true),
LdapBaseDN: pulumi.String("dc=contosoad,dc=contoso,dc=local"),
LdapServer: pulumi.String("192.0.2.12"),
UsernameSource: pulumi.String("LDAP"),
},
},
EncryptionSettings: storagecache.CacheEncryptionSettingsResponse{
KeyEncryptionKey: interface{}{
KeyUrl: pulumi.String("https://keyvault-cmk.vault.azure.net/keys/key2048/test"),
SourceVault: &storagecache.KeyVaultKeyReferenceSourceVaultArgs{
Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"),
},
},
},
Location: pulumi.String("westus"),
ResourceGroupName: pulumi.String("scgroup"),
SecuritySettings: storagecache.CacheSecuritySettingsResponse{
AccessPolicies: storagecache.NfsAccessPolicyArray{
interface{}{
AccessRules: storagecache.NfsAccessRuleArray{
&storagecache.NfsAccessRuleArgs{
Access: pulumi.String("rw"),
RootSquash: pulumi.Bool(false),
Scope: pulumi.String("default"),
SubmountAccess: pulumi.Bool(true),
Suid: pulumi.Bool(false),
},
},
Name: pulumi.String("default"),
},
},
},
Sku: &storagecache.CacheSkuArgs{
Name: pulumi.String("Standard_2G"),
},
Subnet: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/sub1"),
Tags: pulumi.StringMap{
"Dept": pulumi.String("Contoso"),
},
UpgradeSettings: &storagecache.CacheUpgradeSettingsArgs{
ScheduledTime: pulumi.String("2022-04-26T18:25:43.511Z"),
UpgradeScheduleEnabled: pulumi.Bool(true),
},
})
if err != nil {
return err
}
return nil
})
}
