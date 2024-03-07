package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datalakestore/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datalakestore.NewTrustedIdProvider(ctx, "trustedIdProvider", &datalakestore.TrustedIdProviderArgs{
			AccountName:           pulumi.String("contosoadla"),
			IdProvider:            pulumi.String("https://sts.windows.net/ea9ec534-a3e3-4e45-ad36-3afc5bb291c1"),
			ResourceGroupName:     pulumi.String("contosorg"),
			TrustedIdProviderName: pulumi.String("test_trusted_id_provider_name"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
