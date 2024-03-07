package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewCertificate(ctx, "certificate", &automation.CertificateArgs{
			AutomationAccountName: pulumi.String("myAutomationAccount18"),
			Base64Value:           pulumi.String("base 64 value of cert"),
			CertificateName:       pulumi.String("testCert"),
			Description:           pulumi.String("Sample Cert"),
			IsExportable:          pulumi.Bool(false),
			Name:                  pulumi.String("testCert"),
			ResourceGroupName:     pulumi.String("rg"),
			Thumbprint:            pulumi.String("thumbprint of cert"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
