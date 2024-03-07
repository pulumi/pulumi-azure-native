package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationsmanagement.NewSolution(ctx, "solution", &operationsmanagement.SolutionArgs{
			Location: pulumi.String("East US"),
			Plan: &operationsmanagement.SolutionPlanArgs{
				Name:          pulumi.String("name1"),
				Product:       pulumi.String("product1"),
				PromotionCode: pulumi.String("promocode1"),
				Publisher:     pulumi.String("publisher1"),
			},
			Properties: &operationsmanagement.SolutionPropertiesArgs{
				ContainedResources: pulumi.StringArray{
					pulumi.String("/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource1"),
					pulumi.String("/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource2"),
				},
				ReferencedResources: pulumi.StringArray{
					pulumi.String("/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource2"),
					pulumi.String("/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource3"),
				},
				WorkspaceResourceId: pulumi.String("/subscriptions/sub2/resourceGroups/rg2/providers/Microsoft.OperationalInsights/workspaces/ws1"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			SolutionName:      pulumi.String("solution1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
