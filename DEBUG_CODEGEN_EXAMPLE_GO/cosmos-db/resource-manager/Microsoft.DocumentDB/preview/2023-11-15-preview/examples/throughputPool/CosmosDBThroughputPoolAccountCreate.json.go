package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewThroughputPoolAccount(ctx, "throughputPoolAccount", &documentdb.ThroughputPoolAccountArgs{
			AccountLocation:           pulumi.String("West US"),
			AccountResourceIdentifier: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/providers/Microsoft.DocumentDB/resourceGroup/rg1/databaseAccounts/db1/"),
			ResourceGroupName:         pulumi.String("rg1"),
			ThroughputPoolAccountName: pulumi.String("db1"),
			ThroughputPoolName:        pulumi.String("tp1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
