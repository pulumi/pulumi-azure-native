package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewNamedValue(ctx, "namedValue", &apimanagement.NamedValueArgs{
			DisplayName: pulumi.String("prop6namekv"),
			KeyVault: &apimanagement.KeyVaultContractCreatePropertiesArgs{
				IdentityClientId: pulumi.String("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0"),
				SecretIdentifier: pulumi.String("https://contoso.vault.azure.net/secrets/aadSecret"),
			},
			NamedValueId:      pulumi.String("testprop6"),
			ResourceGroupName: pulumi.String("rg1"),
			Secret:            pulumi.Bool(true),
			ServiceName:       pulumi.String("apimService1"),
			Tags: pulumi.StringArray{
				pulumi.String("foo"),
				pulumi.String("bar"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
