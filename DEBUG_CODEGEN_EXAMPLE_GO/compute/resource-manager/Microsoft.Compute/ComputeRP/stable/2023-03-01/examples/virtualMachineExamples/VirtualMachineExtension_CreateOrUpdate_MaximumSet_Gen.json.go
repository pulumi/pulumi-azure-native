package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachineExtension(ctx, "virtualMachineExtension", &compute.VirtualMachineExtensionArgs{
			AutoUpgradeMinorVersion: pulumi.Bool(true),
			EnableAutomaticUpgrade:  pulumi.Bool(true),
			ForceUpdateTag:          pulumi.String("a"),
			InstanceView: &compute.VirtualMachineExtensionInstanceViewArgs{
				Name: pulumi.String("aaaaaaaaaaaaaaaaa"),
				Statuses: compute.InstanceViewStatusArray{
					&compute.InstanceViewStatusArgs{
						Code:          pulumi.String("aaaaaaaaaaaaaaaaaaaaaaa"),
						DisplayStatus: pulumi.String("aaaaaa"),
						Level:         compute.StatusLevelTypesInfo,
						Message:       pulumi.String("a"),
						Time:          pulumi.String("2021-11-30T12:58:26.522Z"),
					},
				},
				Substatuses: compute.InstanceViewStatusArray{
					&compute.InstanceViewStatusArgs{
						Code:          pulumi.String("aaaaaaaaaaaaaaaaaaaaaaa"),
						DisplayStatus: pulumi.String("aaaaaa"),
						Level:         compute.StatusLevelTypesInfo,
						Message:       pulumi.String("a"),
						Time:          pulumi.String("2021-11-30T12:58:26.522Z"),
					},
				},
				Type:               pulumi.String("aaaaaaaaa"),
				TypeHandlerVersion: pulumi.String("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
			},
			Location:          pulumi.String("westus"),
			ProtectedSettings: nil,
			Publisher:         pulumi.String("extPublisher"),
			ResourceGroupName: pulumi.String("rgcompute"),
			Settings:          nil,
			SuppressFailures:  pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"key9183": pulumi.String("aa"),
			},
			Type:               pulumi.String("extType"),
			TypeHandlerVersion: pulumi.String("1.2"),
			VmExtensionName:    pulumi.String("aaaaaaaaaaaaa"),
			VmName:             pulumi.String("aaaaaaaaaaaaaaaaaaaaaaaa"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
