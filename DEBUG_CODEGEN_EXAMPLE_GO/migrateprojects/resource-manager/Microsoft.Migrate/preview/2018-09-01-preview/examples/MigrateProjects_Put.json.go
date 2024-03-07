package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewMigrateProject(ctx, "migrateProject", &migrate.MigrateProjectArgs{
			ETag:               pulumi.String("\"b701c73a-0000-0000-0000-59c12ff00000\""),
			Location:           pulumi.String("Southeast Asia"),
			MigrateProjectName: pulumi.String("project01"),
			Properties:         nil,
			ResourceGroupName:  pulumi.String("myResourceGroup"),
			Tags:               nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
