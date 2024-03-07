package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewCloudService(ctx, "cloudService", &compute.CloudServiceArgs{
CloudServiceName: pulumi.String("{cs-name}"),
Location: pulumi.String("westus"),
Properties: compute.CloudServicePropertiesResponse{
Configuration: pulumi.String("{ServiceConfiguration}"),
NetworkProfile: interface{}{
LoadBalancerConfigurations: compute.LoadBalancerConfigurationArray{
interface{}{
Name: pulumi.String("contosolb"),
Properties: interface{}{
FrontendIpConfigurations: compute.LoadBalancerFrontendIpConfigurationArray{
interface{}{
Name: pulumi.String("contosofe"),
Properties: interface{}{
PublicIPAddress: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Network/publicIPAddresses/contosopublicip"),
},
},
},
},
},
},
},
},
OsProfile: interface{}{
Secrets: compute.CloudServiceVaultSecretGroupArray{
interface{}{
SourceVault: &compute.SubResourceArgs{
Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.KeyVault/vaults/{keyvault-name}"),
},
VaultCertificates: compute.CloudServiceVaultCertificateArray{
&compute.CloudServiceVaultCertificateArgs{
CertificateUrl: pulumi.String("https://{keyvault-name}.vault.azure.net:443/secrets/ContosoCertificate/{secret-id}"),
},
},
},
},
},
PackageUrl: pulumi.String("{PackageUrl}"),
RoleProfile: interface{}{
Roles: compute.CloudServiceRoleProfilePropertiesArray{
interface{}{
Name: pulumi.String("ContosoFrontend"),
Sku: &compute.CloudServiceRoleSkuArgs{
Capacity: pulumi.Float64(1),
Name: pulumi.String("Standard_D1_v2"),
Tier: pulumi.String("Standard"),
},
},
},
},
UpgradeMode: pulumi.String("Auto"),
},
ResourceGroupName: pulumi.String("ConstosoRG"),
})
if err != nil {
return err
}
return nil
})
}
