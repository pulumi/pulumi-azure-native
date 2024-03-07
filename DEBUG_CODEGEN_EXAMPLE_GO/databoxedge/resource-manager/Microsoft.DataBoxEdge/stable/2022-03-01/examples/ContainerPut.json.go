package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databoxedge/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databoxedge.NewContainer(ctx, "container", &databoxedge.ContainerArgs{
			ContainerName:      pulumi.String("blobcontainer1"),
			DataFormat:         pulumi.String("BlockBlob"),
			DeviceName:         pulumi.String("testedgedevice"),
			ResourceGroupName:  pulumi.String("GroupForEdgeAutomation"),
			StorageAccountName: pulumi.String("storageaccount1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
