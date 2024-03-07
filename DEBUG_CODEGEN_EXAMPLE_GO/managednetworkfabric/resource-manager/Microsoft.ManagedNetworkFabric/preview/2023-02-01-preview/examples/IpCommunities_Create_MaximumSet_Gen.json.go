package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewIpCommunity(ctx, "ipCommunity", &managednetworkfabric.IpCommunityArgs{
			Action:     pulumi.String("Permit"),
			Annotation: pulumi.String("annotationValue"),
			CommunityMembers: pulumi.StringArray{
				pulumi.String("1234:5678"),
			},
			IpCommunityName:   pulumi.String("example-ipCommunity"),
			Location:          pulumi.String("EastUS"),
			ResourceGroupName: pulumi.String("rgIpCommunityLists"),
			Tags: pulumi.StringMap{
				"key2814": pulumi.String(""),
			},
			WellKnownCommunities: pulumi.StringArray{
				pulumi.String("Internet"),
				pulumi.String("LocalAS"),
				pulumi.String("NoExport"),
				pulumi.String("GShut"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
