package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewWorkspaceSqlAadAdmin(ctx, "workspaceSqlAadAdmin", &synapse.WorkspaceSqlAadAdminArgs{
			AdministratorType: pulumi.String("ActiveDirectory"),
			Login:             pulumi.String("bob@contoso.com"),
			ResourceGroupName: pulumi.String("resourceGroup1"),
			Sid:               pulumi.String("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
			TenantId:          pulumi.String("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
			WorkspaceName:     pulumi.String("workspace1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
