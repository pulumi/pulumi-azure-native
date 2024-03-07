package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/softwareplan/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := softwareplan.NewHybridUseBenefit(ctx, "hybridUseBenefit", &softwareplan.HybridUseBenefitArgs{
			PlanId: pulumi.String("94f46eda-45f8-493a-8425-251921463a89"),
			Scope:  pulumi.String("subscriptions/{sub-id}/resourceGroups/{rg-name}/providers/Microsoft.Compute/HostGroups/{host-group-name}/hosts/{host-name}"),
			Sku: &softwareplan.SkuArgs{
				Name: pulumi.String("SQL_Server_Perpetual"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
