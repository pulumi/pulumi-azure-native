package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/batch/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.NewPool(ctx, "pool", &batch.PoolArgs{
			AccountName: pulumi.String("sampleacct"),
			DeploymentConfiguration: &batch.DeploymentConfigurationArgs{
				VirtualMachineConfiguration: &batch.VirtualMachineConfigurationArgs{
					ImageReference: &batch.ImageReferenceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/networking-group/providers/Microsoft.Compute/galleries/testgallery/images/testimagedef/versions/0.0.1"),
					},
					NodeAgentSkuId: pulumi.String("batch.node.ubuntu 18.04"),
				},
			},
			NetworkConfiguration: &batch.NetworkConfigurationArgs{
				PublicIPAddressConfiguration: &batch.PublicIPAddressConfigurationArgs{
					IpAddressIds: pulumi.StringArray{
						pulumi.String("/subscriptions/subid1/resourceGroups/rg13/providers/Microsoft.Network/publicIPAddresses/ip135"),
					},
					Provision: batch.IPAddressProvisioningTypeUserManaged,
				},
				SubnetId: pulumi.String("/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123"),
			},
			PoolName:          pulumi.String("testpool"),
			ResourceGroupName: pulumi.String("default-azurebatch-japaneast"),
			VmSize:            pulumi.String("STANDARD_D4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
