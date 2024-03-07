package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewIpExtendedCommunity(ctx, "ipExtendedCommunity", &managednetworkfabric.IpExtendedCommunityArgs{
			Action:                  pulumi.String("Permit"),
			Annotation:              pulumi.String("annotationValue"),
			IpExtendedCommunityName: pulumi.String("example_ipExtendedCommunity"),
			Location:                pulumi.String("EastUs"),
			ResourceGroupName:       pulumi.String("rgIpExtendedCommunityLists"),
			RouteTargets: pulumi.StringArray{
				pulumi.String("1234:5678"),
			},
			Tags: pulumi.StringMap{
				"key5054": pulumi.String("key"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
