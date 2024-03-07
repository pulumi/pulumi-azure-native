package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/certificateregistration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := certificateregistration.NewAppServiceCertificateOrder(ctx, "appServiceCertificateOrder", &certificateregistration.AppServiceCertificateOrderArgs{
			AutoRenew:            pulumi.Bool(true),
			CertificateOrderName: pulumi.String("SampleCertificateOrderName"),
			Certificates: certificateregistration.AppServiceCertificateMap{
				"SampleCertName1": &certificateregistration.AppServiceCertificateArgs{
					KeyVaultId:         pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
					KeyVaultSecretName: pulumi.String("SampleSecretName1"),
				},
				"SampleCertName2": &certificateregistration.AppServiceCertificateArgs{
					KeyVaultId:         pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
					KeyVaultSecretName: pulumi.String("SampleSecretName2"),
				},
			},
			DistinguishedName: pulumi.String("CN=SampleCustomDomain.com"),
			KeySize:           pulumi.Int(2048),
			Location:          pulumi.String("Global"),
			ProductType:       certificateregistration.CertificateProductTypeStandardDomainValidatedSsl,
			ResourceGroupName: pulumi.String("testrg123"),
			ValidityInYears:   pulumi.Int(2),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
