package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewService(ctx, "service", &documentdb.ServiceArgs{
			AccountName:       pulumi.String("ddb1"),
			InstanceCount:     pulumi.Int(1),
			InstanceSize:      pulumi.String("Cosmos.D4s"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("GraphAPICompute"),
			ServiceType:       pulumi.String("GraphAPICompute"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
