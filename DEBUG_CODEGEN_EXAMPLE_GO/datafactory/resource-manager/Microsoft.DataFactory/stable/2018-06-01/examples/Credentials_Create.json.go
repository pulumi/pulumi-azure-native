package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datafactory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datafactory.NewCredentialOperation(ctx, "credentialOperation", &datafactory.CredentialOperationArgs{
			CredentialName: pulumi.String("exampleCredential"),
			FactoryName:    pulumi.String("exampleFactoryName"),
			Properties: &datafactory.ManagedIdentityCredentialArgs{
				ResourceId: pulumi.String("/subscriptions/12345678-1234-1234-1234-12345678abc/resourcegroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleUami"),
				Type:       pulumi.String("ManagedIdentity"),
			},
			ResourceGroupName: pulumi.String("exampleResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
