package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewSecret(ctx, "secret", &cdn.SecretArgs{
			Parameters: cdn.CustomerCertificateParameters{
				SecretSource: cdn.ResourceReference{
					Id: "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/secrets/certificatename",
				},
				SecretVersion:    "abcdef1234578900abcdef1234567890",
				Type:             "CustomerCertificate",
				UseLatestVersion: false,
			},
			ProfileName:       pulumi.String("profile1"),
			ResourceGroupName: pulumi.String("RG"),
			SecretName:        pulumi.String("secret1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
