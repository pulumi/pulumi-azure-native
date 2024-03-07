package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logic.NewIntegrationAccountCertificate(ctx, "integrationAccountCertificate", &logic.IntegrationAccountCertificateArgs{
			CertificateName:        pulumi.String("testCertificate"),
			IntegrationAccountName: pulumi.String("testIntegrationAccount"),
			Key: logic.KeyVaultKeyReferenceResponse{
				KeyName: pulumi.String("<keyName>"),
				KeyVault: &logic.KeyVaultKeyReferenceKeyVaultArgs{
					Id: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testResourceGroup/providers/microsoft.keyvault/vaults/<keyVaultName>"),
				},
				KeyVersion: pulumi.String("87d9764197604449b9b8eb7bd8710868"),
			},
			Location:          pulumi.String("brazilsouth"),
			PublicCertificate: pulumi.String("<publicCertificateValue>"),
			ResourceGroupName: pulumi.String("testResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
