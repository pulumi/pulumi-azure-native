package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbformysql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbformysql.NewAzureADAdministrator(ctx, "azureADAdministrator", &dbformysql.AzureADAdministratorArgs{
			AdministratorName:  pulumi.String("ActiveDirectory"),
			AdministratorType:  pulumi.String("ActiveDirectory"),
			IdentityResourceId: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/test-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-umi"),
			Login:              pulumi.String("bob@contoso.com"),
			ResourceGroupName:  pulumi.String("testrg"),
			ServerName:         pulumi.String("mysqltestsvc4"),
			Sid:                pulumi.String("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
			TenantId:           pulumi.String("c12b7025-bfe2-46c1-b463-993b5e4cd467"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
