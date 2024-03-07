package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewGateway(ctx, "gateway", &appplatform.GatewayArgs{
			GatewayName: pulumi.String("default"),
			Properties: &appplatform.GatewayPropertiesArgs{
				Public: pulumi.Bool(true),
				ResourceRequests: &appplatform.GatewayResourceRequestsArgs{
					Cpu:    pulumi.String("1"),
					Memory: pulumi.String("1G"),
				},
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
