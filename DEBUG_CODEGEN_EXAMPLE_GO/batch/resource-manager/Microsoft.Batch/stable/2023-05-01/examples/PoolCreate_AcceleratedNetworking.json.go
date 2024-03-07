package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/batch/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := batch.NewPool(ctx, "pool", &batch.PoolArgs{
AccountName: pulumi.String("sampleacct"),
DeploymentConfiguration: batch.DeploymentConfigurationResponse{
VirtualMachineConfiguration: interface{}{
ImageReference: &batch.ImageReferenceArgs{
Offer: pulumi.String("WindowsServer"),
Publisher: pulumi.String("MicrosoftWindowsServer"),
Sku: pulumi.String("2016-datacenter-smalldisk"),
Version: pulumi.String("latest"),
},
NodeAgentSkuId: pulumi.String("batch.node.windows amd64"),
},
},
NetworkConfiguration: &batch.NetworkConfigurationArgs{
EnableAcceleratedNetworking: pulumi.Bool(true),
SubnetId: pulumi.String("/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123"),
},
PoolName: pulumi.String("testpool"),
ResourceGroupName: pulumi.String("default-azurebatch-japaneast"),
ScaleSettings: batch.ScaleSettingsResponse{
FixedScale: &batch.FixedScaleSettingsArgs{
TargetDedicatedNodes: pulumi.Int(1),
TargetLowPriorityNodes: pulumi.Int(0),
},
},
VmSize: pulumi.String("STANDARD_D1_V2"),
})
if err != nil {
return err
}
return nil
})
}
