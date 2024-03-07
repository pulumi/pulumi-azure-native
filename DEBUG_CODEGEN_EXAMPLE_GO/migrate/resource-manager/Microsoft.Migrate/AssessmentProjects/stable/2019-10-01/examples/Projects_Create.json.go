package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewProject(ctx, "project", &migrate.ProjectArgs{
			ETag:        pulumi.String(""),
			Location:    pulumi.String("West Europe"),
			ProjectName: pulumi.String("abGoyalProject2"),
			Properties: &migrate.ProjectPropertiesArgs{
				AssessmentSolutionId: pulumi.String("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/abgoyal-westeurope/providers/microsoft.migrate/migrateprojects/abgoyalweselfhost/Solutions/Servers-Assessment-ServerAssessment"),
				ProjectStatus:        pulumi.String("Active"),
			},
			ResourceGroupName: pulumi.String("abgoyal-westEurope"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
