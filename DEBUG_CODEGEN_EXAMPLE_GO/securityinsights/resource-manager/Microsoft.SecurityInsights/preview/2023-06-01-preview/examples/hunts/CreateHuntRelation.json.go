package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewHuntRelation(ctx, "huntRelation", &securityinsights.HuntRelationArgs{
			HuntId:         pulumi.String("163e7b2a-a2ec-4041-aaba-d878a38f265f"),
			HuntRelationId: pulumi.String("2216d0e1-91e3-4902-89fd-d2df8c535096"),
			Labels: pulumi.StringArray{
				pulumi.String("Test Label"),
			},
			RelatedResourceId: pulumi.String("/subscriptions/bd794837-4d29-4647-9105-6339bfdb4e6a/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/workspaces/avdvirint/providers/Microsoft.SecurityInsights/Bookmarks/2216d0e1-91e3-4902-89fd-d2df8c535096"),
			ResourceGroupName: pulumi.String("myRg"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
