package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/certificateregistration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := certificateregistration.NewAppServiceCertificateOrderCertificate(ctx, "appServiceCertificateOrderCertificate", &certificateregistration.AppServiceCertificateOrderCertificateArgs{
			CertificateOrderName: pulumi.String("SampleCertificateOrderName"),
			KeyVaultId:           pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
			KeyVaultSecretName:   pulumi.String("SampleSecretName1"),
			Location:             pulumi.String("Global"),
			Name:                 pulumi.String("SampleCertName1"),
			ResourceGroupName:    pulumi.String("testrg123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
