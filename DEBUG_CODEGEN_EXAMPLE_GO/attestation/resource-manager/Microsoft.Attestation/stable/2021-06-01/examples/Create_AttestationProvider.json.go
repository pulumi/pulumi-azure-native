package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/attestation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := attestation.NewAttestationProvider(ctx, "attestationProvider", &attestation.AttestationProviderArgs{
			Location: pulumi.String("East US"),
			Properties: &attestation.AttestationServiceCreationSpecificParamsArgs{
				PublicNetworkAccess:          pulumi.String("Enabled"),
				TpmAttestationAuthentication: pulumi.String("Enabled"),
			},
			ProviderName:      pulumi.String("myattestationprovider"),
			ResourceGroupName: pulumi.String("MyResourceGroup"),
			Tags: pulumi.StringMap{
				"Property1": pulumi.String("Value1"),
				"Property2": pulumi.String("Value2"),
				"Property3": pulumi.String("Value3"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
