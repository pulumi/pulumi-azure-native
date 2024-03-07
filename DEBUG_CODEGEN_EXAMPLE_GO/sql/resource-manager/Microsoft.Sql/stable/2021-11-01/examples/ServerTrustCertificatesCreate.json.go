package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewServerTrustCertificate(ctx, "serverTrustCertificate", &sql.ServerTrustCertificateArgs{
			CertificateName:     pulumi.String("customerCertificateName"),
			ManagedInstanceName: pulumi.String("testcl"),
			PublicBlob:          pulumi.String("308203AE30820296A0030201020210"),
			ResourceGroupName:   pulumi.String("testrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
