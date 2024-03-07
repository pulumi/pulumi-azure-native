package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewCaCertificate(ctx, "caCertificate", &eventgrid.CaCertificateArgs{
			CaCertificateName:  pulumi.String("exampleCACertificateName1"),
			Description:        pulumi.String("This is a test certificate"),
			EncodedCertificate: pulumi.String("base64EncodePemFormattedCertificateString"),
			NamespaceName:      pulumi.String("exampleNamespaceName1"),
			ResourceGroupName:  pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
