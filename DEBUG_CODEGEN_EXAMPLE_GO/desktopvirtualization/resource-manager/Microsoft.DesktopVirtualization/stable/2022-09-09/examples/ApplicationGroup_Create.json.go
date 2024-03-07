package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/desktopvirtualization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := desktopvirtualization.NewApplicationGroup(ctx, "applicationGroup", &desktopvirtualization.ApplicationGroupArgs{
			ApplicationGroupName: pulumi.String("applicationGroup1"),
			ApplicationGroupType: pulumi.String("RemoteApp"),
			Description:          pulumi.String("des1"),
			FriendlyName:         pulumi.String("friendly"),
			HostPoolArmPath:      pulumi.String("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1"),
			Location:             pulumi.String("centralus"),
			ResourceGroupName:    pulumi.String("resourceGroup1"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
