package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewWorkloadGroup(ctx, "workloadGroup", &sql.WorkloadGroupArgs{
			DatabaseName:                 pulumi.String("testdb"),
			Importance:                   pulumi.String("normal"),
			MaxResourcePercent:           pulumi.Int(100),
			MaxResourcePercentPerRequest: pulumi.Float64(3),
			MinResourcePercent:           pulumi.Int(0),
			MinResourcePercentPerRequest: pulumi.Float64(3),
			QueryExecutionTimeout:        pulumi.Int(0),
			ResourceGroupName:            pulumi.String("Default-SQL-SouthEastAsia"),
			ServerName:                   pulumi.String("testsvr"),
			WorkloadGroupName:            pulumi.String("smallrc"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
