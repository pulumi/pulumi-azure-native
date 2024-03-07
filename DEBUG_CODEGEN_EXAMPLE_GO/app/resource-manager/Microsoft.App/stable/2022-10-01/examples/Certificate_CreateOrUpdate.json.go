package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewCertificate(ctx, "certificate", &app.CertificateArgs{
			CertificateName: pulumi.String("certificate-firendly-name"),
			EnvironmentName: pulumi.String("testcontainerenv"),
			Location:        pulumi.String("East US"),
			Properties: &app.CertificatePropertiesArgs{
				Password: pulumi.String("private key password"),
				Value:    pulumi.String("Y2VydA=="),
			},
			ResourceGroupName: pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
