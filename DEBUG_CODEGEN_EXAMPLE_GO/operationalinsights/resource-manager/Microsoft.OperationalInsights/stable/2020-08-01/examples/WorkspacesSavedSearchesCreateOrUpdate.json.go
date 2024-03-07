package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewSavedSearch(ctx, "savedSearch", &operationalinsights.SavedSearchArgs{
			Category:           pulumi.String("Saved Search Test Category"),
			DisplayName:        pulumi.String("Create or Update Saved Search Test"),
			FunctionAlias:      pulumi.String("heartbeat_func"),
			FunctionParameters: pulumi.String("a:int=1"),
			Query:              pulumi.String("Heartbeat | summarize Count() by Computer | take a"),
			ResourceGroupName:  pulumi.String("TestRG"),
			SavedSearchId:      pulumi.String("00000000-0000-0000-0000-00000000000"),
			Tags: operationalinsights.TagArray{
				&operationalinsights.TagArgs{
					Name:  pulumi.String("Group"),
					Value: pulumi.String("Computer"),
				},
			},
			Version:       pulumi.Float64(2),
			WorkspaceName: pulumi.String("TestWS"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
