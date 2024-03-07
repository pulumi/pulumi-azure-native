package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewCertificate(ctx, "certificate", &web.CertificateArgs{
			HostNames: pulumi.StringArray{
				pulumi.String("ServerCert"),
			},
			Location:          pulumi.String("East US"),
			Name:              pulumi.String("testc6282"),
			Password:          pulumi.String("<password>"),
			ResourceGroupName: pulumi.String("testrg123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
