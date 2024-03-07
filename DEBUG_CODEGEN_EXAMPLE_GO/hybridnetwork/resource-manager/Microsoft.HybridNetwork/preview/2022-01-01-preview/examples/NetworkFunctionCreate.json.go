package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewNetworkFunction(ctx, "networkFunction", &hybridnetwork.NetworkFunctionArgs{
			Device: &hybridnetwork.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/devices/testDevice"),
			},
			Location:                     pulumi.String("eastus"),
			ManagedApplicationParameters: nil,
			NetworkFunctionName:          pulumi.String("testNf"),
			NetworkFunctionUserConfigurations: hybridnetwork.NetworkFunctionUserConfigurationArray{
				&hybridnetwork.NetworkFunctionUserConfigurationArgs{
					NetworkInterfaces: hybridnetwork.NetworkInterfaceArray{
						&hybridnetwork.NetworkInterfaceArgs{
							IpConfigurations: hybridnetwork.NetworkInterfaceIPConfigurationArray{
								&hybridnetwork.NetworkInterfaceIPConfigurationArgs{
									Gateway:            pulumi.String(""),
									IpAddress:          pulumi.String(""),
									IpAllocationMethod: pulumi.String("Dynamic"),
									IpVersion:          pulumi.String("IPv4"),
									Subnet:             pulumi.String(""),
								},
							},
							MacAddress:           pulumi.String(""),
							NetworkInterfaceName: pulumi.String("nic1"),
							VmSwitchType:         pulumi.String("Management"),
						},
						&hybridnetwork.NetworkInterfaceArgs{
							IpConfigurations: hybridnetwork.NetworkInterfaceIPConfigurationArray{
								&hybridnetwork.NetworkInterfaceIPConfigurationArgs{
									Gateway:            pulumi.String(""),
									IpAddress:          pulumi.String(""),
									IpAllocationMethod: pulumi.String("Dynamic"),
									IpVersion:          pulumi.String("IPv4"),
									Subnet:             pulumi.String(""),
								},
							},
							MacAddress:           pulumi.String("DC-97-F8-79-16-7D"),
							NetworkInterfaceName: pulumi.String("nic2"),
							VmSwitchType:         pulumi.String("Wan"),
						},
					},
					RoleName:           pulumi.String("testRole"),
					UserDataParameters: nil,
				},
			},
			ResourceGroupName: pulumi.String("rg"),
			SkuName:           pulumi.String("testSku"),
			VendorName:        pulumi.String("testVendor"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
