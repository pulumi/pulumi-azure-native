package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewRouteFilterRule(ctx, "routeFilterRule", &network.RouteFilterRuleArgs{
			Access: pulumi.String("Allow"),
			Communities: pulumi.StringArray{
				pulumi.String("12076:5030"),
				pulumi.String("12076:5040"),
			},
			ResourceGroupName:   pulumi.String("rg1"),
			RouteFilterName:     pulumi.String("filterName"),
			RouteFilterRuleType: pulumi.String("Community"),
			RuleName:            pulumi.String("ruleName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
