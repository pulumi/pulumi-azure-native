package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := appplatform.NewService(ctx, "service", &appplatform.ServiceArgs{
Location: pulumi.String("eastus"),
Properties: appplatform.ClusterResourcePropertiesResponse{
NetworkProfile: interface{}{
AppNetworkResourceGroup: pulumi.String("my-app-network-rg"),
AppSubnetId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/apps"),
IngressConfig: &appplatform.IngressConfigArgs{
ReadTimeoutInSeconds: pulumi.Int(300),
},
ServiceCidr: pulumi.String("10.8.0.0/16,10.244.0.0/16,10.245.0.1/16"),
ServiceRuntimeNetworkResourceGroup: pulumi.String("my-service-runtime-network-rg"),
ServiceRuntimeSubnetId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/serviceRuntime"),
},
VnetAddons: &appplatform.ServiceVNetAddonsArgs{
DataPlanePublicEndpoint: pulumi.Bool(true),
LogStreamPublicEndpoint: pulumi.Bool(true),
},
},
ResourceGroupName: pulumi.String("myResourceGroup"),
ServiceName: pulumi.String("myservice"),
Sku: &appplatform.SkuArgs{
Name: pulumi.String("S0"),
Tier: pulumi.String("Standard"),
},
Tags: pulumi.StringMap{
"key1": pulumi.String("value1"),
},
})
if err != nil {
return err
}
return nil
})
}
