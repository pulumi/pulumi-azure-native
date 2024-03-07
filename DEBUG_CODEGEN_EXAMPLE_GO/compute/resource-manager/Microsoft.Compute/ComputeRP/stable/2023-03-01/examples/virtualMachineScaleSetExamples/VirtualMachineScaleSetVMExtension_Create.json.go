package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachineScaleSetVMExtension(ctx, "virtualMachineScaleSetVMExtension", &compute.VirtualMachineScaleSetVMExtensionArgs{
			AutoUpgradeMinorVersion: pulumi.Bool(true),
			InstanceId:              pulumi.String("0"),
			Publisher:               pulumi.String("extPublisher"),
			ResourceGroupName:       pulumi.String("myResourceGroup"),
			Settings: pulumi.Any{
				UserName: "xyz@microsoft.com",
			},
			Type:               pulumi.String("extType"),
			TypeHandlerVersion: pulumi.String("1.2"),
			VmExtensionName:    pulumi.String("myVMExtension"),
			VmScaleSetName:     pulumi.String("myvmScaleSet"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
