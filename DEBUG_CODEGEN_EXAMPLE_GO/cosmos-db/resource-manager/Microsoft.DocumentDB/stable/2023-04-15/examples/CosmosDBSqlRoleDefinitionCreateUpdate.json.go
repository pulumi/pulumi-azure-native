package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewSqlResourceSqlRoleDefinition(ctx, "sqlResourceSqlRoleDefinition", &documentdb.SqlResourceSqlRoleDefinitionArgs{
			AccountName: pulumi.String("myAccountName"),
			AssignableScopes: pulumi.StringArray{
				pulumi.String("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/sales"),
				pulumi.String("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases"),
			},
			Permissions: []documentdb.PermissionArgs{
				{
					DataActions: pulumi.StringArray{
						pulumi.String("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/create"),
						pulumi.String("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/read"),
					},
					NotDataActions: pulumi.StringArray{},
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroupName"),
			RoleDefinitionId:  pulumi.String("myRoleDefinitionId"),
			RoleName:          pulumi.String("myRoleName"),
			Type:              documentdb.RoleDefinitionTypeCustomRole,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
