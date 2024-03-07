package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewSqlResourceSqlContainer(ctx, "sqlResourceSqlContainer", &documentdb.SqlResourceSqlContainerArgs{
			AccountName:   pulumi.String("ddb1"),
			ContainerName: pulumi.String("containerName"),
			DatabaseName:  pulumi.String("databaseName"),
			Location:      pulumi.String("West US"),
			Options:       nil,
			Resource: &documentdb.SqlContainerResourceArgs{
				ClientEncryptionPolicy: &documentdb.ClientEncryptionPolicyArgs{
					IncludedPaths: documentdb.ClientEncryptionIncludedPathArray{
						&documentdb.ClientEncryptionIncludedPathArgs{
							ClientEncryptionKeyId: pulumi.String("keyId"),
							EncryptionAlgorithm:   pulumi.String("AEAD_AES_256_CBC_HMAC_SHA256"),
							EncryptionType:        pulumi.String("Deterministic"),
							Path:                  pulumi.String("/path"),
						},
					},
					PolicyFormatVersion: pulumi.Int(2),
				},
				ConflictResolutionPolicy: &documentdb.ConflictResolutionPolicyArgs{
					ConflictResolutionPath: pulumi.String("/path"),
					Mode:                   pulumi.String("LastWriterWins"),
				},
				DefaultTtl: pulumi.Int(100),
				Id:         pulumi.String("containerName"),
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
