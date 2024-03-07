package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/extendedlocation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := extendedlocation.NewResourceSyncRule(ctx, "resourceSyncRule", &extendedlocation.ResourceSyncRuleArgs{
			ChildResourceName: pulumi.String("resourceSyncRule01"),
			Location:          pulumi.String("West US"),
			Priority:          pulumi.Int(999),
			ResourceGroupName: pulumi.String("testresourcegroup"),
			ResourceName:      pulumi.String("customLocation01"),
			Selector: &extendedlocation.ResourceSyncRulePropertiesSelectorArgs{
				MatchLabels: pulumi.StringMap{
					"key1": pulumi.String("value1"),
				},
			},
			TargetResourceGroup: pulumi.String("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
