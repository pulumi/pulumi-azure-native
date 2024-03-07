package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/costmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costmanagement.NewCostAllocationRule(ctx, "costAllocationRule", &costmanagement.CostAllocationRuleArgs{
			BillingAccountId: pulumi.String("100"),
			Properties: &costmanagement.CostAllocationRulePropertiesArgs{
				Description: pulumi.String("This is a testRule"),
				Details: &costmanagement.CostAllocationRuleDetailsArgs{
					SourceResources: costmanagement.SourceCostAllocationResourceArray{
						&costmanagement.SourceCostAllocationResourceArgs{
							Name:         pulumi.String("ResourceGroupName"),
							ResourceType: pulumi.String("Dimension"),
							Values: pulumi.StringArray{
								pulumi.String("sampleRG"),
								pulumi.String("secondRG"),
							},
						},
					},
					TargetResources: costmanagement.TargetCostAllocationResourceArray{
						&costmanagement.TargetCostAllocationResourceArgs{
							Name:         pulumi.String("ResourceGroupName"),
							PolicyType:   pulumi.String("FixedProportion"),
							ResourceType: pulumi.String("Dimension"),
							Values: costmanagement.CostAllocationProportionArray{
								&costmanagement.CostAllocationProportionArgs{
									Name:       pulumi.String("destinationRG"),
									Percentage: pulumi.Float64(45),
								},
								&costmanagement.CostAllocationProportionArgs{
									Name:       pulumi.String("destinationRG2"),
									Percentage: pulumi.Float64(54),
								},
							},
						},
					},
				},
				Status: pulumi.String("Active"),
			},
			RuleName: pulumi.String("testRule"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
