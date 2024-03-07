package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewAppServiceEnvironmentAseCustomDnsSuffixConfiguration(ctx, "appServiceEnvironmentAseCustomDnsSuffixConfiguration", &web.AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs{
			CertificateUrl:            pulumi.String("https://test-kv.vault.azure.net/secrets/contosocert"),
			DnsSuffix:                 pulumi.String("contoso.com"),
			KeyVaultReferenceIdentity: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/test-rg/providers/microsoft.managedidentity/userassignedidentities/test-user-mi"),
			Name:                      pulumi.String("test-ase"),
			ResourceGroupName:         pulumi.String("test-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
