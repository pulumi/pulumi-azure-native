package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/addons/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := addons.NewSupportPlanType(ctx, "supportPlanType", &addons.SupportPlanTypeArgs{
			PlanTypeName: pulumi.String("Standard"),
			ProviderName: pulumi.String("Canonical"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
