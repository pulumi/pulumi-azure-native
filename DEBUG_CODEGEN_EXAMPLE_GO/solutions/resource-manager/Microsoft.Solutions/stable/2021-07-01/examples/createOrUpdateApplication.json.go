package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/solutions/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := solutions.NewApplication(ctx, "application", &solutions.ApplicationArgs{
			ApplicationDefinitionId: pulumi.String("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef"),
			ApplicationName:         pulumi.String("myManagedApplication"),
			Kind:                    pulumi.String("ServiceCatalog"),
			ManagedResourceGroupId:  pulumi.String("/subscriptions/subid/resourceGroups/myManagedRG"),
			ResourceGroupName:       pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
