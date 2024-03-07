package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewWorkloadDeployment(ctx, "workloadDeployment", &migrate.WorkloadDeploymentArgs{
			ModernizeProjectName:   pulumi.String("tc"),
			ResourceGroupName:      pulumi.String("rgmigrateEngine"),
			WorkloadDeploymentName: pulumi.String("wo2rs4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
