package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewPacketCapture(ctx, "packetCapture", &network.PacketCaptureArgs{
			BytesToCapturePerPacket: pulumi.Float64(10000),
			Filters: []network.PacketCaptureFilterArgs{
				{
					LocalIPAddress: pulumi.String("10.0.0.4"),
					LocalPort:      pulumi.String("80"),
					Protocol:       pulumi.String("TCP"),
				},
			},
			NetworkWatcherName: pulumi.String("nw1"),
			PacketCaptureName:  pulumi.String("pc1"),
			ResourceGroupName:  pulumi.String("rg1"),
			StorageLocation: &network.PacketCaptureStorageLocationArgs{
				FilePath:    pulumi.String("D:\\capture\\pc1.cap"),
				StorageId:   pulumi.String("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Storage/storageAccounts/pcstore"),
				StoragePath: pulumi.String("https://mytestaccountname.blob.core.windows.net/capture/pc1.cap"),
			},
			Target:               pulumi.String("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
			TimeLimitInSeconds:   pulumi.Int(100),
			TotalBytesPerSession: pulumi.Float64(100000),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
