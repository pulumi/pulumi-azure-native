package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datalakestore/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datalakestore.NewAccount(ctx, "account", &datalakestore.AccountArgs{
			AccountName:  pulumi.String("contosoadla"),
			DefaultGroup: pulumi.String("test_default_group"),
			EncryptionConfig: datalakestore.EncryptionConfigResponse{
				KeyVaultMetaInfo: &datalakestore.KeyVaultMetaInfoArgs{
					EncryptionKeyName:    pulumi.String("test_encryption_key_name"),
					EncryptionKeyVersion: pulumi.String("encryption_key_version"),
					KeyVaultResourceId:   pulumi.String("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
				},
				Type: datalakestore.EncryptionConfigTypeUserManaged,
			},
			EncryptionState:       datalakestore.EncryptionStateEnabled,
			FirewallAllowAzureIps: datalakestore.FirewallAllowAzureIpsStateEnabled,
			FirewallRules: []datalakestore.CreateFirewallRuleWithAccountParametersArgs{
				{
					EndIpAddress:   pulumi.String("2.2.2.2"),
					Name:           pulumi.String("test_rule"),
					StartIpAddress: pulumi.String("1.1.1.1"),
				},
			},
			FirewallState: datalakestore.FirewallStateEnabled,
			Identity: &datalakestore.EncryptionIdentityArgs{
				Type: datalakestore.EncryptionIdentityTypeSystemAssigned,
			},
			Location:          pulumi.String("eastus2"),
			NewTier:           datalakestore.TierTypeConsumption,
			ResourceGroupName: pulumi.String("contosorg"),
			Tags: pulumi.StringMap{
				"test_key": pulumi.String("test_value"),
			},
			TrustedIdProviderState: datalakestore.TrustedIdProviderStateEnabled,
			TrustedIdProviders: []datalakestore.CreateTrustedIdProviderWithAccountParametersArgs{
				{
					IdProvider: pulumi.String("https://sts.windows.net/ea9ec534-a3e3-4e45-ad36-3afc5bb291c1"),
					Name:       pulumi.String("test_trusted_id_provider_name"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
