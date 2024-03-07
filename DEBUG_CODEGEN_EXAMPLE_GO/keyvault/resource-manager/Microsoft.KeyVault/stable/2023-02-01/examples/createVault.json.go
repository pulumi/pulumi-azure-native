package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/keyvault/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := keyvault.NewVault(ctx, "vault", &keyvault.VaultArgs{
Location: pulumi.String("westus"),
Properties: keyvault.VaultPropertiesResponse{
AccessPolicies: keyvault.AccessPolicyEntryArray{
interface{}{
ObjectId: pulumi.String("00000000-0000-0000-0000-000000000000"),
Permissions: &keyvault.PermissionsArgs{
Certificates: pulumi.StringArray{
pulumi.String("get"),
pulumi.String("list"),
pulumi.String("delete"),
pulumi.String("create"),
pulumi.String("import"),
pulumi.String("update"),
pulumi.String("managecontacts"),
pulumi.String("getissuers"),
pulumi.String("listissuers"),
pulumi.String("setissuers"),
pulumi.String("deleteissuers"),
pulumi.String("manageissuers"),
pulumi.String("recover"),
pulumi.String("purge"),
},
Keys: pulumi.StringArray{
pulumi.String("encrypt"),
pulumi.String("decrypt"),
pulumi.String("wrapKey"),
pulumi.String("unwrapKey"),
pulumi.String("sign"),
pulumi.String("verify"),
pulumi.String("get"),
pulumi.String("list"),
pulumi.String("create"),
pulumi.String("update"),
pulumi.String("import"),
pulumi.String("delete"),
pulumi.String("backup"),
pulumi.String("restore"),
pulumi.String("recover"),
pulumi.String("purge"),
},
Secrets: pulumi.StringArray{
pulumi.String("get"),
pulumi.String("list"),
pulumi.String("set"),
pulumi.String("delete"),
pulumi.String("backup"),
pulumi.String("restore"),
pulumi.String("recover"),
pulumi.String("purge"),
},
},
TenantId: pulumi.String("00000000-0000-0000-0000-000000000000"),
},
},
EnabledForDeployment: pulumi.Bool(true),
EnabledForDiskEncryption: pulumi.Bool(true),
EnabledForTemplateDeployment: pulumi.Bool(true),
PublicNetworkAccess: pulumi.String("Enabled"),
Sku: &keyvault.SkuArgs{
Family: pulumi.String("A"),
Name: keyvault.SkuNameStandard,
},
TenantId: pulumi.String("00000000-0000-0000-0000-000000000000"),
},
ResourceGroupName: pulumi.String("sample-resource-group"),
VaultName: pulumi.String("sample-vault"),
})
if err != nil {
return err
}
return nil
})
}
