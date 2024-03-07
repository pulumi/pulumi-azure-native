package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewEdgeDevice(ctx, "edgeDevice", &azurestackhci.EdgeDeviceArgs{
			DeviceConfiguration: &azurestackhci.DeviceConfigurationArgs{
				DeviceMetadata: pulumi.String(""),
				NicDetails: azurestackhci.NicDetailArray{
					&azurestackhci.NicDetailArgs{
						AdapterName:        pulumi.String("ethernet"),
						ComponentId:        pulumi.String("VMBUS{f8615163-df3e-46c5-913f-f2d2f965ed0g} "),
						DefaultGateway:     pulumi.String("10.10.10.1"),
						DefaultIsolationId: pulumi.String("0"),
						DnsServers: pulumi.StringArray{
							pulumi.String("100.10.10.1"),
						},
						DriverVersion:        pulumi.String("10.0.20348.1547 "),
						InterfaceDescription: pulumi.String("NDIS 6.70 "),
						Ip4Address:           pulumi.String("10.10.10.10"),
						SubnetMask:           pulumi.String("255.255.255.0"),
					},
				},
			},
			EdgeDeviceName: pulumi.String("default"),
			ResourceUri:    pulumi.String("subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
