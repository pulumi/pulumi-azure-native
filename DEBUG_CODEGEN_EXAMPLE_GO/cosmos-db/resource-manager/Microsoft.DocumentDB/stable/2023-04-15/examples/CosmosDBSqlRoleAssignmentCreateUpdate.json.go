package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewSqlResourceSqlRoleAssignment(ctx, "sqlResourceSqlRoleAssignment", &documentdb.SqlResourceSqlRoleAssignmentArgs{
			AccountName:       pulumi.String("myAccountName"),
			PrincipalId:       pulumi.String("myPrincipalId"),
			ResourceGroupName: pulumi.String("myResourceGroupName"),
			RoleAssignmentId:  pulumi.String("myRoleAssignmentId"),
			RoleDefinitionId:  pulumi.String("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleDefinitions/myRoleDefinitionId"),
			Scope:             pulumi.String("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
