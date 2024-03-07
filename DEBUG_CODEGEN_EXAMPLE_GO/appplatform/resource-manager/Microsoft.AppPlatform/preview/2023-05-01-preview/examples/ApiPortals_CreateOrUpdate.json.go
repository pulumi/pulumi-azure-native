package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewApiPortal(ctx, "apiPortal", &appplatform.ApiPortalArgs{
			ApiPortalName: pulumi.String("default"),
			Properties: &appplatform.ApiPortalPropertiesArgs{
				GatewayIds: pulumi.StringArray{
					pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/gateways/default"),
				},
				Public: pulumi.Bool(true),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
			Sku: &appplatform.SkuArgs{
				Capacity: pulumi.Int(2),
				Name:     pulumi.String("E0"),
				Tier:     pulumi.String("Enterprise"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
