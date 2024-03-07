package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customerinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerinsights.NewHub(ctx, "hub", &customerinsights.HubArgs{
			HubBillingInfo: &customerinsights.HubBillingInfoFormatArgs{
				MaxUnits: pulumi.Int(5),
				MinUnits: pulumi.Int(1),
				SkuName:  pulumi.String("B0"),
			},
			HubName:           pulumi.String("sdkTestHub"),
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("TestHubRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
