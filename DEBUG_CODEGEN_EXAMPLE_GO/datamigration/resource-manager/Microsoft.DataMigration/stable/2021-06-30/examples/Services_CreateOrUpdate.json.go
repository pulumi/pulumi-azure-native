package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datamigration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datamigration.NewService(ctx, "service", &datamigration.ServiceArgs{
			GroupName:   pulumi.String("DmsSdkRg"),
			Location:    pulumi.String("southcentralus"),
			ServiceName: pulumi.String("DmsSdkService"),
			Sku: &datamigration.ServiceSkuArgs{
				Name: pulumi.String("Basic_1vCore"),
			},
			VirtualSubnetId: pulumi.String("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkTestNetwork/providers/Microsoft.Network/virtualNetworks/DmsSdkTestNetwork/subnets/default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
