package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewScopeAccessReviewHistoryDefinitionById(ctx, "scopeAccessReviewHistoryDefinitionById", &authorization.ScopeAccessReviewHistoryDefinitionByIdArgs{
			HistoryDefinitionId: pulumi.String("44724910-d7a5-4c29-b28f-db73e717165a"),
			Scope:               pulumi.String("subscriptions/129a304b-4aea-4b86-a9f7-ba7e2b23737a"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
