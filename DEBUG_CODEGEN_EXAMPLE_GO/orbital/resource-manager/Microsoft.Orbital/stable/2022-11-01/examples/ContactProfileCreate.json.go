package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/orbital/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := orbital.NewContactProfile(ctx, "contactProfile", &orbital.ContactProfileArgs{
			AutoTrackingConfiguration: orbital.AutoTrackingConfigurationDisabled,
			ContactProfileName:        pulumi.String("CONTOSO-CP"),
			EventHubUri:               pulumi.String("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.EventHub/namespaces/contosoHub/eventhubs/contosoHub"),
			Links: []orbital.ContactProfileLinkArgs{
				{
					Channels: orbital.ContactProfileLinkChannelArray{
						{
							BandwidthMHz:       pulumi.Float64(2),
							CenterFrequencyMHz: pulumi.Float64(2250),
							EndPoint: {
								EndPointName: pulumi.String("ContosoTest_Uplink"),
								IpAddress:    pulumi.String("10.1.0.4"),
								Port:         pulumi.String("50000"),
								Protocol:     pulumi.String("TCP"),
							},
							Name: pulumi.String("contoso-uplink-channel"),
						},
					},
					Direction:           pulumi.String("Uplink"),
					EirpdBW:             pulumi.Float64(45),
					GainOverTemperature: pulumi.Float64(0),
					Name:                pulumi.String("contoso-uplink"),
					Polarization:        pulumi.String("LHCP"),
				},
				{
					Channels: orbital.ContactProfileLinkChannelArray{
						{
							BandwidthMHz:       pulumi.Float64(15),
							CenterFrequencyMHz: pulumi.Float64(8160),
							EndPoint: {
								EndPointName: pulumi.String("ContosoTest_Downlink"),
								IpAddress:    pulumi.String("10.1.0.5"),
								Port:         pulumi.String("50001"),
								Protocol:     pulumi.String("UDP"),
							},
							Name: pulumi.String("contoso-downlink-channel"),
						},
					},
					Direction:           pulumi.String("Downlink"),
					EirpdBW:             pulumi.Float64(0),
					GainOverTemperature: pulumi.Float64(25),
					Name:                pulumi.String("contoso-downlink"),
					Polarization:        pulumi.String("RHCP"),
				},
			},
			Location:                     pulumi.String("eastus2"),
			MinimumElevationDegrees:      pulumi.Float64(5),
			MinimumViableContactDuration: pulumi.String("PT1M"),
			NetworkConfiguration: &orbital.ContactProfilesPropertiesNetworkConfigurationArgs{
				SubnetId: pulumi.String("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Network/virtualNetworks/contoso-vnet/subnets/orbital-delegated-subnet"),
			},
			ResourceGroupName: pulumi.String("contoso-Rgp"),
			ThirdPartyConfigurations: []orbital.ContactProfileThirdPartyConfigurationArgs{
				{
					MissionConfiguration: pulumi.String("Ksat_MissionConfiguration"),
					ProviderName:         pulumi.String("KSAT"),
				},
				{
					MissionConfiguration: pulumi.String("Viasat_Configuration"),
					ProviderName:         pulumi.String("VIASAT"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
