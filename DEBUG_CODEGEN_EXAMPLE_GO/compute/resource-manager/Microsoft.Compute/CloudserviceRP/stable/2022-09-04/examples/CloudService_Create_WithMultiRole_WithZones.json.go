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
interface{}{
Name: pulumi.String("ContosoBackend"),
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
Zones: pulumi.StringArray{
pulumi.String("1"),
},
})
if err != nil {
return err
}
return nil
})
}
