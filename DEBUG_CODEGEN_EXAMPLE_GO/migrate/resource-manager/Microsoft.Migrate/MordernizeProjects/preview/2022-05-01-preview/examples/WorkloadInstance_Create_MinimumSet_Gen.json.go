package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewWorkloadInstance(ctx, "workloadInstance", &migrate.WorkloadInstanceArgs{
			ModernizeProjectName: pulumi.String("tv39"),
			ResourceGroupName:    pulumi.String("rgmigrateEngine"),
			WorkloadInstanceName: pulumi.String("io"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
