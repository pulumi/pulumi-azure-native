package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devices.NewCertificate(ctx, "certificate", &devices.CertificateArgs{
			CertificateName: pulumi.String("cert"),
			Properties: &devices.CertificatePropertiesArgs{
				Certificate: pulumi.String("############################################"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ResourceName:      pulumi.String("iothub"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
