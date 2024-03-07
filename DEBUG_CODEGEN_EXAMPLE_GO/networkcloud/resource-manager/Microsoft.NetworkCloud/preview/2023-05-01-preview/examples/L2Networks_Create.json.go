package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkcloud.NewL2Network(ctx, "l2Network", &networkcloud.L2NetworkArgs{
			ExtendedLocation: &networkcloud.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
				Type: pulumi.String("CustomLocation"),
			},
			HybridAksPluginType: pulumi.String("DPDK"),
			InterfaceName:       pulumi.String("eth0"),
			L2IsolationDomainId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/l2IsolationDomainName"),
			L2NetworkName:       pulumi.String("l2NetworkName"),
			Location:            pulumi.String("location"),
			ResourceGroupName:   pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("myvalue1"),
				"key2": pulumi.String("myvalue2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
