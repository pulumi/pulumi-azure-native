package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/vmwarecloudsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vmwarecloudsimple.NewVirtualMachine(ctx, "virtualMachine", &vmwarecloudsimple.VirtualMachineArgs{
			AmountOfRam: pulumi.Int(4096),
			Disks: []vmwarecloudsimple.VirtualDiskArgs{
				{
					ControllerId:     pulumi.String("1000"),
					IndependenceMode: vmwarecloudsimple.DiskIndependenceModePersistent,
					TotalSize:        pulumi.Int(10485760),
					VirtualDiskId:    pulumi.String("2000"),
				},
			},
			Location: pulumi.String("westus2"),
			Nics: []vmwarecloudsimple.VirtualNicArgs{
				{
					Network: {
						Id: pulumi.String("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
					},
					NicType:      vmwarecloudsimple.NICTypeE1000,
					PowerOnBoot:  pulumi.Bool(true),
					VirtualNicId: pulumi.String("4000"),
				},
			},
			NumberOfCores:     pulumi.Int(2),
			PrivateCloudId:    pulumi.String("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ResourcePool: &vmwarecloudsimple.ResourcePoolArgs{
				Id: pulumi.String("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/resourcePools/resgroup-26"),
			},
			TemplateId:         pulumi.String("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualMachineTemplates/vm-34"),
			VirtualMachineName: pulumi.String("myVirtualMachine"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
