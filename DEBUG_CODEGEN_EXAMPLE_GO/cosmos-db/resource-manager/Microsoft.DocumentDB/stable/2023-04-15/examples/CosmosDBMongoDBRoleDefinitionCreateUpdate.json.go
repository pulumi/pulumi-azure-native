package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewMongoDBResourceMongoRoleDefinition(ctx, "mongoDBResourceMongoRoleDefinition", &documentdb.MongoDBResourceMongoRoleDefinitionArgs{
			AccountName:           pulumi.String("myAccountName"),
			DatabaseName:          pulumi.String("sales"),
			MongoRoleDefinitionId: pulumi.String("myMongoRoleDefinitionId"),
			Privileges: documentdb.PrivilegeArray{
				&documentdb.PrivilegeArgs{
					Actions: pulumi.StringArray{
						pulumi.String("insert"),
						pulumi.String("find"),
					},
					Resource: &documentdb.PrivilegeResourceArgs{
						Collection: pulumi.String("sales"),
						Db:         pulumi.String("sales"),
					},
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroupName"),
			RoleName:          pulumi.String("myRoleName"),
			Roles: documentdb.RoleArray{
				&documentdb.RoleArgs{
					Db:   pulumi.String("sales"),
					Role: pulumi.String("myInheritedRole"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
