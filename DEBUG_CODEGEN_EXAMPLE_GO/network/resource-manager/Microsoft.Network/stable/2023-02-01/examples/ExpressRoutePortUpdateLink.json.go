package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewExpressRoutePort(ctx, "expressRoutePort", &network.ExpressRoutePortArgs{
			BandwidthInGbps:      pulumi.Int(100),
			BillingType:          pulumi.String("UnlimitedData"),
			Encapsulation:        pulumi.String("QinQ"),
			ExpressRoutePortName: pulumi.String("portName"),
			Links: []network.ExpressRouteLinkArgs{
				{
					AdminState: pulumi.String("Enabled"),
					Name:       pulumi.String("link1"),
				},
			},
			Location:          pulumi.String("westus"),
			PeeringLocation:   pulumi.String("peeringLocationName"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
