package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewRouteFilter(ctx, "routeFilter", &network.RouteFilterArgs{
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("rg1"),
			RouteFilterName:   pulumi.String("filterName"),
			Rules: network.RouteFilterRuleTypeArray{
				&network.RouteFilterRuleTypeArgs{
					Access: pulumi.String("Allow"),
					Communities: pulumi.StringArray{
						pulumi.String("12076:5030"),
						pulumi.String("12076:5040"),
					},
					Name:                pulumi.String("ruleName"),
					RouteFilterRuleType: pulumi.String("Community"),
				},
			},
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
