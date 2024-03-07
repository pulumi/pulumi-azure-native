package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewBookmarkRelation(ctx, "bookmarkRelation", &securityinsights.BookmarkRelationArgs{
			BookmarkId:        pulumi.String("2216d0e1-91e3-4902-89fd-d2df8c535096"),
			RelatedResourceId: pulumi.String("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/afbd324f-6c48-459c-8710-8d1e1cd03812"),
			RelationName:      pulumi.String("4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014"),
			ResourceGroupName: pulumi.String("myRg"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
