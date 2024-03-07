package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewIpGroup(ctx, "ipGroup", &network.IpGroupArgs{
			IpAddresses: pulumi.StringArray{
				pulumi.String("13.64.39.16/32"),
				pulumi.String("40.74.146.80/31"),
				pulumi.String("40.74.147.32/28"),
			},
			IpGroupsName:      pulumi.String("ipGroups1"),
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
