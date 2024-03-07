package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewCredentialSet(ctx, "credentialSet", &containerregistry.CredentialSetArgs{
			AuthCredentials: containerregistry.AuthCredentialArray{
				&containerregistry.AuthCredentialArgs{
					Name:                     pulumi.String("Credential1"),
					PasswordSecretIdentifier: pulumi.String("https://myvault.vault.azure.net/secrets/password"),
					UsernameSecretIdentifier: pulumi.String("https://myvault.vault.azure.net/secrets/username"),
				},
			},
			CredentialSetName: pulumi.String("myCredentialSet"),
			Identity: &containerregistry.IdentityPropertiesArgs{
				Type: containerregistry.ResourceIdentityTypeSystemAssigned,
			},
			LoginServer:       pulumi.String("docker.io"),
			RegistryName:      pulumi.String("myRegistry"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
