package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewMigrateProjectsControllerMigrateProject(ctx, "migrateProjectsControllerMigrateProject", &migrate.MigrateProjectsControllerMigrateProjectArgs{
			Location:           pulumi.String("eastus"),
			MigrateProjectName: pulumi.String("projTest1"),
			Properties: &migrate.MigrateProjectPropertiesArgs{
				PublicNetworkAccess: pulumi.String("Enabled"),
			},
			ResourceGroupName: pulumi.String("pajindTest1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
