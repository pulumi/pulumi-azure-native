package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachineScaleSetVM(ctx, "virtualMachineScaleSetVM", &compute.VirtualMachineScaleSetVMArgs{
			InstanceId:        pulumi.String("aaaaaaaaaaaaaaaaaaaa"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("rgcompute"),
			VmScaleSetName:    pulumi.String("aaaaaaaaaaaaaaaaaa"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
