package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewGraphResourceGraph(ctx, "graphResourceGraph", &documentdb.GraphResourceGraphArgs{
			AccountName: pulumi.String("ddb1"),
			GraphName:   pulumi.String("graphName"),
			Location:    pulumi.String("West US"),
			Options:     nil,
			Resource: &documentdb.GraphResourceArgs{
				Id: pulumi.String("graphName"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
