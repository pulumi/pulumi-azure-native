package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewSolutionsControllerSolution(ctx, "solutionsControllerSolution", &migrate.SolutionsControllerSolutionArgs{
			MigrateProjectName: pulumi.String("1GQwlI-"),
			Properties: &migrate.SolutionPropertiesArgs{
				CleanupState: pulumi.String("None"),
				Details: &migrate.SolutionDetailsArgs{
					AssessmentCount: pulumi.Int(7),
					ExtendedDetails: pulumi.StringMap{
						"key9031": pulumi.String("ombnjq"),
					},
					GroupCount: pulumi.Int(30),
				},
				Goal:    pulumi.String("Servers"),
				Purpose: pulumi.String("Discovery"),
				Status:  pulumi.String("Inactive"),
				Tool:    pulumi.String("ServerDiscovery"),
			},
			ResourceGroupName: pulumi.String("rghubmigrate"),
			SolutionName:      pulumi.String("1zJY9v3KFENX698GSOl"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
