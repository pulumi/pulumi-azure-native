package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewMongoDBResourceMongoDBCollection(ctx, "mongoDBResourceMongoDBCollection", &documentdb.MongoDBResourceMongoDBCollectionArgs{
			AccountName:    pulumi.String("ddb1"),
			CollectionName: pulumi.String("collectionName"),
			DatabaseName:   pulumi.String("databaseName"),
			Location:       pulumi.String("West US"),
			Options:        nil,
			Resource: &documentdb.MongoDBCollectionResourceArgs{
				Id: pulumi.String("collectionName"),
				Indexes: documentdb.MongoIndexArray{
					&documentdb.MongoIndexArgs{
						Key: &documentdb.MongoIndexKeysArgs{
							Keys: pulumi.StringArray{
								pulumi.String("_ts"),
							},
						},
						Options: &documentdb.MongoIndexOptionsArgs{
							ExpireAfterSeconds: pulumi.Int(100),
							Unique:             pulumi.Bool(true),
						},
					},
					&documentdb.MongoIndexArgs{
						Key: &documentdb.MongoIndexKeysArgs{
							Keys: pulumi.StringArray{
								pulumi.String("_id"),
							},
						},
					},
				},
				ShardKey: pulumi.StringMap{
					"testKey": pulumi.String("Hash"),
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
