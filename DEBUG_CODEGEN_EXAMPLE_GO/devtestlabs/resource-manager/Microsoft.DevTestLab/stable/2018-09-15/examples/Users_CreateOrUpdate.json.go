package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewUser(ctx, "user", &devtestlab.UserArgs{
			Identity: &devtestlab.UserIdentityArgs{
				AppId:         pulumi.String("{appId}"),
				ObjectId:      pulumi.String("{objectId}"),
				PrincipalId:   pulumi.String("{principalId}"),
				PrincipalName: pulumi.String("{principalName}"),
				TenantId:      pulumi.String("{tenantId}"),
			},
			LabName:           pulumi.String("{devtestlabName}"),
			Location:          pulumi.String("{location}"),
			Name:              pulumi.String("{userName}"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			SecretStore: &devtestlab.UserSecretStoreArgs{
				KeyVaultId:  pulumi.String("{keyVaultId}"),
				KeyVaultUri: pulumi.String("{keyVaultUri}"),
			},
			Tags: pulumi.StringMap{
				"tagName1": pulumi.String("tagValue1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
