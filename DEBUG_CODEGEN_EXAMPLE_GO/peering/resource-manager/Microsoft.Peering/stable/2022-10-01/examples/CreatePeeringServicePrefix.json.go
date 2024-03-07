package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/peering/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := peering.NewPrefix(ctx, "prefix", &peering.PrefixArgs{
			PeeringServiceName:      pulumi.String("peeringServiceName"),
			PeeringServicePrefixKey: pulumi.String("00000000-0000-0000-0000-000000000000"),
			Prefix:                  pulumi.String("192.168.1.0/24"),
			PrefixName:              pulumi.String("peeringServicePrefixName"),
			ResourceGroupName:       pulumi.String("rgName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
