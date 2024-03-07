package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/peering/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := peering.NewPeeringService(ctx, "peeringService", &peering.PeeringServiceArgs{
			Location:                       pulumi.String("eastus"),
			PeeringServiceLocation:         pulumi.String("state1"),
			PeeringServiceName:             pulumi.String("peeringServiceName"),
			PeeringServiceProvider:         pulumi.String("serviceProvider1"),
			ProviderBackupPeeringLocation:  pulumi.String("peeringLocation2"),
			ProviderPrimaryPeeringLocation: pulumi.String("peeringLocation1"),
			ResourceGroupName:              pulumi.String("rgName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
