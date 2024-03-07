package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewPacketCapture(ctx, "packetCapture", &mobilenetwork.PacketCaptureArgs{
			BytesToCapturePerPacket: pulumi.Float64(10000),
			NetworkInterfaces: pulumi.StringArray{
				pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP"),
				pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP"),
				pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestADN"),
			},
			PacketCaptureName:          pulumi.String("pc1"),
			PacketCoreControlPlaneName: pulumi.String("TestPacketCoreCP"),
			ResourceGroupName:          pulumi.String("rg1"),
			TimeLimitInSeconds:         pulumi.Int(100),
			TotalBytesPerSession:       pulumi.Float64(100000),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
