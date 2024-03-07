package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewCertificate(ctx, "certificate", &appplatform.CertificateArgs{
			CertificateName: pulumi.String("mycertificate"),
			Properties: appplatform.KeyVaultCertificateProperties{
				CertVersion:      "08a219d06d874795a96db47e06fbb01e",
				KeyVaultCertName: "mycert",
				Type:             "KeyVaultCertificate",
				VaultUri:         "https://myvault.vault.azure.net",
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
