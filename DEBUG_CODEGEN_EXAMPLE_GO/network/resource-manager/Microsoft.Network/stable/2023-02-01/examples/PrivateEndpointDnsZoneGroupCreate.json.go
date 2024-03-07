package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewPrivateDnsZoneGroup(ctx, "privateDnsZoneGroup", &network.PrivateDnsZoneGroupArgs{
			PrivateDnsZoneConfigs: []network.PrivateDnsZoneConfigArgs{
				{
					PrivateDnsZoneId: pulumi.String("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateDnsZones/zone1.com"),
				},
			},
			PrivateDnsZoneGroupName: pulumi.String("testPdnsgroup"),
			PrivateEndpointName:     pulumi.String("testPe"),
			ResourceGroupName:       pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
