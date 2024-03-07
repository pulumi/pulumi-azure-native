package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/peering/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := peering.NewRegisteredAsn(ctx, "registeredAsn", &peering.RegisteredAsnArgs{
			Asn:               pulumi.Int(65000),
			PeeringName:       pulumi.String("peeringName"),
			RegisteredAsnName: pulumi.String("registeredAsnName"),
			ResourceGroupName: pulumi.String("rgName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
