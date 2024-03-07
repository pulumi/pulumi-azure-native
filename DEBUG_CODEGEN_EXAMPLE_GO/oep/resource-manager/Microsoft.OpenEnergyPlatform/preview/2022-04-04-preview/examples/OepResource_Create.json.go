package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/openenergyplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := openenergyplatform.NewEnergyService(ctx, "energyService", &openenergyplatform.EnergyServiceArgs{
			ResourceGroupName: pulumi.String("DummyResourceGroupName"),
			ResourceName:      pulumi.String("DummyResourceName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
