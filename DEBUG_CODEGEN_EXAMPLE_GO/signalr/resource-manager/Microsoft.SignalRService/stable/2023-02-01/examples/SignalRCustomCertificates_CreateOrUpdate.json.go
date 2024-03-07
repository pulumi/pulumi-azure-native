package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/signalrservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := signalrservice.NewSignalRCustomCertificate(ctx, "signalRCustomCertificate", &signalrservice.SignalRCustomCertificateArgs{
			CertificateName:       pulumi.String("myCert"),
			KeyVaultBaseUri:       pulumi.String("https://myvault.keyvault.azure.net/"),
			KeyVaultSecretName:    pulumi.String("mycert"),
			KeyVaultSecretVersion: pulumi.String("bb6a44b2743f47f68dad0d6cc9756432"),
			ResourceGroupName:     pulumi.String("myResourceGroup"),
			ResourceName:          pulumi.String("mySignalRService"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
