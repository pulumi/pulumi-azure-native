package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/chaos/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := chaos.NewCapability(ctx, "capability", &chaos.CapabilityArgs{
			CapabilityName:          pulumi.String("Shutdown-1.0"),
			ParentProviderNamespace: pulumi.String("Microsoft.Compute"),
			ParentResourceName:      pulumi.String("exampleVM"),
			ParentResourceType:      pulumi.String("virtualMachines"),
			ResourceGroupName:       pulumi.String("exampleRG"),
			TargetName:              pulumi.String("Microsoft-VirtualMachine"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
