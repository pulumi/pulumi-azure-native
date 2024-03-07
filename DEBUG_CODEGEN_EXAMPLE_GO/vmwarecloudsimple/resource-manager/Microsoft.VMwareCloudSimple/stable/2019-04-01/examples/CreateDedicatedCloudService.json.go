package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/vmwarecloudsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vmwarecloudsimple.NewDedicatedCloudService(ctx, "dedicatedCloudService", &vmwarecloudsimple.DedicatedCloudServiceArgs{
			DedicatedCloudServiceName: pulumi.String("myService"),
			GatewaySubnet:             pulumi.String("10.0.0.0"),
			Location:                  pulumi.String("westus"),
			ResourceGroupName:         pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
