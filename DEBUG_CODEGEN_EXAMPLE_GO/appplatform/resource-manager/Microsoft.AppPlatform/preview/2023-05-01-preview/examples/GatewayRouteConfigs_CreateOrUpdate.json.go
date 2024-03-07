package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewGatewayRouteConfig(ctx, "gatewayRouteConfig", &appplatform.GatewayRouteConfigArgs{
			GatewayName: pulumi.String("default"),
			Properties: &appplatform.GatewayRouteConfigPropertiesArgs{
				AppResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myApp"),
				OpenApi: &appplatform.GatewayRouteConfigOpenApiPropertiesArgs{
					Uri: pulumi.String("https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore.json"),
				},
				Protocol: pulumi.String("HTTPS"),
				Routes: appplatform.GatewayApiRouteArray{
					&appplatform.GatewayApiRouteArgs{
						Filters: pulumi.StringArray{
							pulumi.String("StripPrefix=2"),
							pulumi.String("RateLimit=1,1s"),
						},
						Predicates: pulumi.StringArray{
							pulumi.String("Path=/api5/customer/**"),
						},
						SsoEnabled: pulumi.Bool(true),
						Title:      pulumi.String("myApp route config"),
					},
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			RouteConfigName:   pulumi.String("myRouteConfig"),
			ServiceName:       pulumi.String("myservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
