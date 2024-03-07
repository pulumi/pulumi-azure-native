package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachineScaleSetExtension(ctx, "virtualMachineScaleSetExtension", &compute.VirtualMachineScaleSetExtensionArgs{
			AutoUpgradeMinorVersion: pulumi.Bool(true),
			EnableAutomaticUpgrade:  pulumi.Bool(true),
			ForceUpdateTag:          pulumi.String("aaaaaaaaa"),
			Name:                    pulumi.String("{extension-name}"),
			ProtectedSettings:       nil,
			ProvisionAfterExtensions: pulumi.StringArray{
				pulumi.String("aa"),
			},
			Publisher:          pulumi.String("{extension-Publisher}"),
			ResourceGroupName:  pulumi.String("rgcompute"),
			Settings:           nil,
			SuppressFailures:   pulumi.Bool(true),
			Type:               pulumi.String("{extension-Type}"),
			TypeHandlerVersion: pulumi.String("{handler-version}"),
			VmScaleSetName:     pulumi.String("aaaaaaa"),
			VmssExtensionName:  pulumi.String("aaaaaaaaaaaaaaaaaaaaa"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
