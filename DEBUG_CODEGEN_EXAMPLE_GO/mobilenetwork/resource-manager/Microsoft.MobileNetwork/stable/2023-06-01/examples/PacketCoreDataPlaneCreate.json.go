package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewPacketCoreDataPlane(ctx, "packetCoreDataPlane", &mobilenetwork.PacketCoreDataPlaneArgs{
			Location:                   pulumi.String("eastus"),
			PacketCoreControlPlaneName: pulumi.String("testPacketCoreCP"),
			PacketCoreDataPlaneName:    pulumi.String("testPacketCoreDP"),
			ResourceGroupName:          pulumi.String("rg1"),
			UserPlaneAccessInterface: &mobilenetwork.InterfacePropertiesArgs{
				Name: pulumi.String("N3"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
