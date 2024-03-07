package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewThroughputPool(ctx, "throughputPool", &documentdb.ThroughputPoolArgs{
			Location:           pulumi.String("westus2"),
			MaxThroughput:      pulumi.Int(10000),
			ResourceGroupName:  pulumi.String("rg1"),
			Tags:               nil,
			ThroughputPoolName: pulumi.String("tp1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
