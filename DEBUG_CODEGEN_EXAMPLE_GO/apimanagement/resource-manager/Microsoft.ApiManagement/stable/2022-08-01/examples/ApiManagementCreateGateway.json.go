package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewGateway(ctx, "gateway", &apimanagement.GatewayArgs{
			Description: pulumi.String("my gateway 1"),
			GatewayId:   pulumi.String("gw1"),
			LocationData: &apimanagement.ResourceLocationDataContractArgs{
				Name: pulumi.String("my location"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
