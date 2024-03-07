package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewAttachedDataNetwork(ctx, "attachedDataNetwork", &mobilenetwork.AttachedDataNetworkArgs{
			AttachedDataNetworkName: pulumi.String("TestAttachedDataNetwork"),
			DnsAddresses: pulumi.StringArray{
				pulumi.String("1.1.1.1"),
			},
			Location: pulumi.String("eastus"),
			NaptConfiguration: mobilenetwork.NaptConfigurationResponse{
				Enabled:       pulumi.String("Enabled"),
				PinholeLimits: pulumi.Int(65536),
				PinholeTimeouts: &mobilenetwork.PinholeTimeoutsArgs{
					Icmp: pulumi.Int(30),
					Tcp:  pulumi.Int(180),
					Udp:  pulumi.Int(30),
				},
				PortRange: &mobilenetwork.PortRangeArgs{
					MaxPort: pulumi.Int(49999),
					MinPort: pulumi.Int(1024),
				},
				PortReuseHoldTime: &mobilenetwork.PortReuseHoldTimesArgs{
					Tcp: pulumi.Int(120),
					Udp: pulumi.Int(60),
				},
			},
			PacketCoreControlPlaneName: pulumi.String("TestPacketCoreCP"),
			PacketCoreDataPlaneName:    pulumi.String("TestPacketCoreDP"),
			ResourceGroupName:          pulumi.String("rg1"),
			UserEquipmentAddressPoolPrefix: pulumi.StringArray{
				pulumi.String("2.2.0.0/16"),
			},
			UserEquipmentStaticAddressPoolPrefix: pulumi.StringArray{
				pulumi.String("2.4.0.0/16"),
			},
			UserPlaneDataInterface: &mobilenetwork.InterfacePropertiesArgs{
				Name: pulumi.String("N6"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
