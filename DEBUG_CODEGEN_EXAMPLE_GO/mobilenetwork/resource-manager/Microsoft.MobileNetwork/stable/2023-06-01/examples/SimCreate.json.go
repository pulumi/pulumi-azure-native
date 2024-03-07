package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewSim(ctx, "sim", &mobilenetwork.SimArgs{
			AuthenticationKey:                     pulumi.String("00000000000000000000000000000000"),
			DeviceType:                            pulumi.String("Video camera"),
			IntegratedCircuitCardIdentifier:       pulumi.String("8900000000000000000"),
			InternationalMobileSubscriberIdentity: pulumi.String("00000"),
			OperatorKeyCode:                       pulumi.String("00000000000000000000000000000000"),
			ResourceGroupName:                     pulumi.String("rg1"),
			SimGroupName:                          pulumi.String("testSimGroup"),
			SimName:                               pulumi.String("testSim"),
			SimPolicy: &mobilenetwork.SimPolicyResourceIdArgs{
				Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
			},
			StaticIpConfiguration: mobilenetwork.SimStaticIpPropertiesArray{
				&mobilenetwork.SimStaticIpPropertiesArgs{
					AttachedDataNetwork: &mobilenetwork.AttachedDataNetworkResourceIdArgs{
						Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
					},
					Slice: &mobilenetwork.SliceResourceIdArgs{
						Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
					},
					StaticIp: &mobilenetwork.SimStaticIpPropertiesStaticIpArgs{
						Ipv4Address: pulumi.String("2.4.0.1"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
