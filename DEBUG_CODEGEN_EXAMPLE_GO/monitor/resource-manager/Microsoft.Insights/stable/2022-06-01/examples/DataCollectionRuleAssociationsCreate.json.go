package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewDataCollectionRuleAssociation(ctx, "dataCollectionRuleAssociation", &insights.DataCollectionRuleAssociationArgs{
			AssociationName:      pulumi.String("myAssociation"),
			DataCollectionRuleId: pulumi.String("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionRules/myCollectionRule"),
			ResourceUri:          pulumi.String("subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
