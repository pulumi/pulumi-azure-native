package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewSolution(ctx, "solution", &migrate.SolutionArgs{
			MigrateProjectName: pulumi.String("project01"),
			Properties: &migrate.SolutionPropertiesArgs{
				Goal:    pulumi.String("Databases"),
				Purpose: pulumi.String("Assessment"),
				Tool:    pulumi.String("DataMigrationAssistant"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SolutionName:      pulumi.String("dbsolution"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
