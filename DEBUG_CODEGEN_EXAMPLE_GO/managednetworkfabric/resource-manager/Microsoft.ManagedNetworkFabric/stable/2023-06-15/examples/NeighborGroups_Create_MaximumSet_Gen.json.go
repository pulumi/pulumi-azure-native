package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewNeighborGroup(ctx, "neighborGroup", &managednetworkfabric.NeighborGroupArgs{
			Annotation: pulumi.String("annotation"),
			Destination: &managednetworkfabric.NeighborGroupDestinationArgs{
				Ipv4Addresses: pulumi.StringArray{
					pulumi.String("10.10.10.10"),
					pulumi.String("20.10.10.10"),
					pulumi.String("30.10.10.10"),
					pulumi.String("40.10.10.10"),
					pulumi.String("50.10.10.10"),
					pulumi.String("60.10.10.10"),
					pulumi.String("70.10.10.10"),
					pulumi.String("80.10.10.10"),
					pulumi.String("90.10.10.10"),
				},
				Ipv6Addresses: pulumi.StringArray{
					pulumi.String("2F::/100"),
				},
			},
			Location:          pulumi.String("eastus"),
			NeighborGroupName: pulumi.String("example-neighborGroup"),
			ResourceGroupName: pulumi.String("example-rg"),
			Tags: pulumi.StringMap{
				"key8107": pulumi.String("1234"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
