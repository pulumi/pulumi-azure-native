package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewDevice(ctx, "device", &hybridnetwork.DeviceArgs{
			DeviceName:        pulumi.String("TestDevice"),
			DeviceType:        pulumi.String("AzureStackEdge"),
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
