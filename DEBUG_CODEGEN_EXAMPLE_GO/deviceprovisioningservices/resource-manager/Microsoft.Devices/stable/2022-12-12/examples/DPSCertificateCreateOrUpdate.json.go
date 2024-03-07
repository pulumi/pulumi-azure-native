package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devices.NewDpsCertificate(ctx, "dpsCertificate", &devices.DpsCertificateArgs{
			CertificateName: pulumi.String("cert"),
			Properties: &devices.CertificatePropertiesArgs{
				Certificate: pulumi.String("MA=="),
			},
			ProvisioningServiceName: pulumi.String("myFirstProvisioningService"),
			ResourceGroupName:       pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
