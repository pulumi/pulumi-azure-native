package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewSimPolicy(ctx, "simPolicy", &mobilenetwork.SimPolicyArgs{
			DefaultSlice: &mobilenetwork.SliceResourceIdArgs{
				Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
			},
			Location:          pulumi.String("eastus"),
			MobileNetworkName: pulumi.String("testMobileNetwork"),
			RegistrationTimer: pulumi.Int(3240),
			ResourceGroupName: pulumi.String("rg1"),
			SimPolicyName:     pulumi.String("testPolicy"),
			SliceConfigurations: []mobilenetwork.SliceConfigurationArgs{
				{
					DataNetworkConfigurations: mobilenetwork.DataNetworkConfigurationArray{
						{
							AdditionalAllowedSessionTypes:       pulumi.StringArray{},
							AllocationAndRetentionPriorityLevel: pulumi.Int(9),
							AllowedServices: mobilenetwork.ServiceResourceIdArray{
								{
									Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/services/testService"),
								},
							},
							DataNetwork: {
								Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork"),
							},
							DefaultSessionType:             pulumi.String("IPv4"),
							FiveQi:                         pulumi.Int(9),
							MaximumNumberOfBufferedPackets: pulumi.Int(200),
							PreemptionCapability:           pulumi.String("NotPreempt"),
							PreemptionVulnerability:        pulumi.String("Preemptable"),
							SessionAmbr: {
								Downlink: pulumi.String("1 Gbps"),
								Uplink:   pulumi.String("500 Mbps"),
							},
						},
					},
					DefaultDataNetwork: {
						Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork"),
					},
					Slice: {
						Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
					},
				},
			},
			UeAmbr: &mobilenetwork.AmbrArgs{
				Downlink: pulumi.String("1 Gbps"),
				Uplink:   pulumi.String("500 Mbps"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
