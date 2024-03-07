package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/peering/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := peering.NewRegisteredPrefix(ctx, "registeredPrefix", &peering.RegisteredPrefixArgs{
			PeeringName:          pulumi.String("peeringName"),
			Prefix:               pulumi.String("10.22.20.0/24"),
			RegisteredPrefixName: pulumi.String("registeredPrefixName"),
			ResourceGroupName:    pulumi.String("rgName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
