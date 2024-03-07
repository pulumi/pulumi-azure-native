package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/orbital/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := orbital.NewL2Connection(ctx, "l2Connection", &orbital.L2ConnectionArgs{
			EdgeSite: &orbital.L2ConnectionsPropertiesEdgeSiteArgs{
				Id: pulumi.String("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/rg1/providers/Microsoft.Orbital/edgeSites/es1"),
			},
			GroundStation: &orbital.L2ConnectionsPropertiesGroundStationArgs{
				Id: pulumi.String("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/rg1/providers/Microsoft.Orbital/groundStations/gs1"),
			},
			L2ConnectionName:  pulumi.String("connection1"),
			Location:          pulumi.String("westus"),
			Name:              pulumi.String("customerName-SiteName-01"),
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
			VlanId: pulumi.Int(200),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
