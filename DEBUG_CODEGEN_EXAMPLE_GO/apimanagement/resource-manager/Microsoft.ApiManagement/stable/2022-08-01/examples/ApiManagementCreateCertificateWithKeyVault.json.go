package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewCertificate(ctx, "certificate", &apimanagement.CertificateArgs{
			CertificateId: pulumi.String("templateCertkv"),
			KeyVault: &apimanagement.KeyVaultContractCreatePropertiesArgs{
				IdentityClientId: pulumi.String("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0"),
				SecretIdentifier: pulumi.String("https://rpbvtkeyvaultintegration.vault-int.azure-int.net/secrets/msitestingCert"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
