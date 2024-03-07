package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databoxedge/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databoxedge.NewIoTAddon(ctx, "ioTAddon", &databoxedge.IoTAddonArgs{
			AddonName:         pulumi.String("arcName"),
			DeviceName:        pulumi.String("testedgedevice"),
			ResourceGroupName: pulumi.String("GroupForEdgeAutomation"),
			RoleName:          pulumi.String("KubernetesRole"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
