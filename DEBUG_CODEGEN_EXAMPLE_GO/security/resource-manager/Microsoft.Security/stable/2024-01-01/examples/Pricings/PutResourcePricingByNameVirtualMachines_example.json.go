package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewPricing(ctx, "pricing", &security.PricingArgs{
			PricingName: pulumi.String("virtualMachines"),
			PricingTier: pulumi.String("Standard"),
			ScopeId:     pulumi.String("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/DEMO/providers/Microsoft.Compute/virtualMachines/VM-1"),
			SubPlan:     pulumi.String("P1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
