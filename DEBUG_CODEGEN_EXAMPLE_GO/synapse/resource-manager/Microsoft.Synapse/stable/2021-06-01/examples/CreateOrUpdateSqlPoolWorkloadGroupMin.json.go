package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewSqlPoolWorkloadGroup(ctx, "sqlPoolWorkloadGroup", &synapse.SqlPoolWorkloadGroupArgs{
			MaxResourcePercent:           pulumi.Int(100),
			MinResourcePercent:           pulumi.Int(0),
			MinResourcePercentPerRequest: pulumi.Float64(3),
			ResourceGroupName:            pulumi.String("sqlcrudtest-6852"),
			SqlPoolName:                  pulumi.String("sqlcrudtest-9187"),
			WorkloadGroupName:            pulumi.String("smallrc"),
			WorkspaceName:                pulumi.String("sqlcrudtest-2080"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
