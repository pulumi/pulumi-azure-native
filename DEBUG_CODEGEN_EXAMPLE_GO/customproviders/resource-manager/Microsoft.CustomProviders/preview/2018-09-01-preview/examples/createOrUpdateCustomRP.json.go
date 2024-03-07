package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customproviders/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customproviders.NewCustomResourceProvider(ctx, "customResourceProvider", &customproviders.CustomResourceProviderArgs{
			Actions: customproviders.CustomRPActionRouteDefinitionArray{
				&customproviders.CustomRPActionRouteDefinitionArgs{
					Endpoint:    pulumi.String("https://mytestendpoint/"),
					Name:        pulumi.String("TestAction"),
					RoutingType: pulumi.String("Proxy"),
				},
			},
			Location:             pulumi.String("eastus"),
			ResourceGroupName:    pulumi.String("testRG"),
			ResourceProviderName: pulumi.String("newrp"),
			ResourceTypes: customproviders.CustomRPResourceTypeRouteDefinitionArray{
				&customproviders.CustomRPResourceTypeRouteDefinitionArgs{
					Endpoint:    pulumi.String("https://mytestendpoint2/"),
					Name:        pulumi.String("TestResource"),
					RoutingType: pulumi.String("Proxy,Cache"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
