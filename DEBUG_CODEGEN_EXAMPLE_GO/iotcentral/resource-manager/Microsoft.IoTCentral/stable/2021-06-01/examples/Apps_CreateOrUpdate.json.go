package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotcentral/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotcentral.NewApp(ctx, "app", &iotcentral.AppArgs{
			DisplayName: pulumi.String("My IoT Central App"),
			Identity: &iotcentral.SystemAssignedServiceIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("resRg"),
			ResourceName:      pulumi.String("myIoTCentralApp"),
			Sku: &iotcentral.AppSkuInfoArgs{
				Name: pulumi.String("ST2"),
			},
			Subdomain: pulumi.String("my-iot-central-app"),
			Template:  pulumi.String("iotc-pnp-preview@1.0.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
