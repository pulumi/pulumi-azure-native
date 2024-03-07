package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualApplianceSite(ctx, "virtualApplianceSite", &network.VirtualApplianceSiteArgs{
			AddressPrefix:               pulumi.String("192.168.1.0/24"),
			NetworkVirtualApplianceName: pulumi.String("nva"),
			O365Policy: &network.Office365PolicyPropertiesArgs{
				BreakOutCategories: &network.BreakOutCategoryPoliciesArgs{
					Allow:    pulumi.Bool(true),
					Default:  pulumi.Bool(true),
					Optimize: pulumi.Bool(true),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			SiteName:          pulumi.String("site1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
