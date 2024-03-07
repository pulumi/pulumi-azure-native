package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewMongoDBResourceMongoUserDefinition(ctx, "mongoDBResourceMongoUserDefinition", &documentdb.MongoDBResourceMongoUserDefinitionArgs{
			AccountName:           pulumi.String("myAccountName"),
			CustomData:            pulumi.String("My custom data"),
			DatabaseName:          pulumi.String("sales"),
			Mechanisms:            pulumi.String("SCRAM-SHA-256"),
			MongoUserDefinitionId: pulumi.String("myMongoUserDefinitionId"),
			Password:              pulumi.String("myPassword"),
			ResourceGroupName:     pulumi.String("myResourceGroupName"),
			Roles: documentdb.RoleArray{
				&documentdb.RoleArgs{
					Db:   pulumi.String("sales"),
					Role: pulumi.String("myReadRole"),
				},
			},
			UserName: pulumi.String("myUserName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
