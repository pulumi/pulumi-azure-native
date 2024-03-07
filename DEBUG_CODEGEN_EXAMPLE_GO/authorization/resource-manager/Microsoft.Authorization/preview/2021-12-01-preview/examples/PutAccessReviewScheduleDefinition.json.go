package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewScopeAccessReviewScheduleDefinitionById(ctx, "scopeAccessReviewScheduleDefinitionById", &authorization.ScopeAccessReviewScheduleDefinitionByIdArgs{
			ScheduleDefinitionId: pulumi.String("fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
			Scope:                pulumi.String("subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
