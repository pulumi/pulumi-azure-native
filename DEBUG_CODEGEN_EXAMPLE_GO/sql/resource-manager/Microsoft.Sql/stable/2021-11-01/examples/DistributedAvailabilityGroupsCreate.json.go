package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDistributedAvailabilityGroup(ctx, "distributedAvailabilityGroup", &sql.DistributedAvailabilityGroupArgs{
			DistributedAvailabilityGroupName: pulumi.String("dag"),
			ManagedInstanceName:              pulumi.String("testcl"),
			PrimaryAvailabilityGroupName:     pulumi.String("BoxLocalAg1"),
			ResourceGroupName:                pulumi.String("testrg"),
			SecondaryAvailabilityGroupName:   pulumi.String("testcl"),
			SourceEndpoint:                   pulumi.String("TCP://SERVER:7022"),
			TargetDatabase:                   pulumi.String("testdb"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
