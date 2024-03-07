package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewModernizeProject(ctx, "modernizeProject", &migrate.ModernizeProjectArgs{
			ModernizeProjectName: pulumi.String("j"),
			ResourceGroupName:    pulumi.String("rgmigrateEngine"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
