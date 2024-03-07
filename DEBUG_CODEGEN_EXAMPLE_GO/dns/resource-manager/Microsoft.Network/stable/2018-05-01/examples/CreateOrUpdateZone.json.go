package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewZone(ctx, "zone", &network.ZoneArgs{
			Location:          pulumi.String("Global"),
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			ZoneName: pulumi.String("zone1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
