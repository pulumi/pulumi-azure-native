package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewGremlinResourceGremlinGraph(ctx, "gremlinResourceGremlinGraph", &documentdb.GremlinResourceGremlinGraphArgs{
			AccountName:  pulumi.String("ddb1"),
			DatabaseName: pulumi.String("databaseName"),
			GraphName:    pulumi.String("graphName"),
			Location:     pulumi.String("West US"),
			Options:      nil,
			Resource: &documentdb.GremlinGraphResourceArgs{
				ConflictResolutionPolicy: &documentdb.ConflictResolutionPolicyArgs{
					ConflictResolutionPath: pulumi.String("/path"),
					Mode:                   pulumi.String("LastWriterWins"),
				},
				DefaultTtl: pulumi.Int(100),
				Id:         pulumi.String("graphName"),
				IndexingPolicy: &documentdb.IndexingPolicyArgs{
					Automatic:     pulumi.Bool(true),
					ExcludedPaths: documentdb.ExcludedPathArray{},
					IncludedPaths: documentdb.IncludedPathArray{
						&documentdb.IncludedPathArgs{
							Indexes: documentdb.IndexesArray{
								&documentdb.IndexesArgs{
									DataType:  pulumi.String("String"),
									Kind:      pulumi.String("Range"),
									Precision: -1,
								},
								&documentdb.IndexesArgs{
									DataType:  pulumi.String("Number"),
									Kind:      pulumi.String("Range"),
									Precision: -1,
								},
							},
							Path: pulumi.String("/*"),
						},
					},
					IndexingMode: pulumi.String("consistent"),
				},
				PartitionKey: &documentdb.ContainerPartitionKeyArgs{
					Kind: pulumi.String("Hash"),
					Paths: pulumi.StringArray{
						pulumi.String("/AccountNumber"),
					},
				},
				UniqueKeyPolicy: &documentdb.UniqueKeyPolicyArgs{
					UniqueKeys: documentdb.UniqueKeyArray{
						&documentdb.UniqueKeyArgs{
							Paths: pulumi.StringArray{
								pulumi.String("/testPath"),
							},
						},
					},
				},
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
