package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkcloud.NewL3Network(ctx, "l3Network", &networkcloud.L3NetworkArgs{
			ExtendedLocation: &networkcloud.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
				Type: pulumi.String("CustomLocation"),
			},
			InterfaceName:       pulumi.String("eth0"),
			IpAllocationType:    pulumi.String("DualStack"),
			Ipv4ConnectedPrefix: pulumi.String("198.51.100.0/24"),
			Ipv6ConnectedPrefix: pulumi.String("2001:db8::/64"),
			L3IsolationDomainId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/l3IsolationDomainName"),
			L3NetworkName:       pulumi.String("l3NetworkName"),
			Location:            pulumi.String("location"),
			ResourceGroupName:   pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("myvalue1"),
				"key2": pulumi.String("myvalue2"),
			},
			Vlan: pulumi.Float64(12),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
