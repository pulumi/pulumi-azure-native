package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewRouteMap(ctx, "routeMap", &network.RouteMapArgs{
			AssociatedInboundConnections: pulumi.StringArray{
				pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGateway1/expressRouteConnections/exrConn1"),
			},
			AssociatedOutboundConnections: pulumi.StringArray{},
			ResourceGroupName:             pulumi.String("rg1"),
			RouteMapName:                  pulumi.String("routeMap1"),
			Rules: network.RouteMapRuleArray{
				&network.RouteMapRuleArgs{
					Actions: network.ActionArray{
						&network.ActionArgs{
							Parameters: network.ParameterArray{
								&network.ParameterArgs{
									AsPath: pulumi.StringArray{
										pulumi.String("22334"),
									},
									Community:   pulumi.StringArray{},
									RoutePrefix: pulumi.StringArray{},
								},
							},
							Type: pulumi.String("Add"),
						},
					},
					MatchCriteria: network.CriterionArray{
						&network.CriterionArgs{
							AsPath:         pulumi.StringArray{},
							Community:      pulumi.StringArray{},
							MatchCondition: pulumi.String("Contains"),
							RoutePrefix: pulumi.StringArray{
								pulumi.String("10.0.0.0/8"),
							},
						},
					},
					Name:              pulumi.String("rule1"),
					NextStepIfMatched: pulumi.String("Continue"),
				},
			},
			VirtualHubName: pulumi.String("virtualHub1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
