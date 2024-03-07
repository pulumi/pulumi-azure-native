package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachineScaleSetExtension(ctx, "virtualMachineScaleSetExtension", &compute.VirtualMachineScaleSetExtensionArgs{
			ResourceGroupName: pulumi.String("rgcompute"),
			VmScaleSetName:    pulumi.String("aaaaaaaaaaa"),
			VmssExtensionName: pulumi.String("aaaaaaaaaaa"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
