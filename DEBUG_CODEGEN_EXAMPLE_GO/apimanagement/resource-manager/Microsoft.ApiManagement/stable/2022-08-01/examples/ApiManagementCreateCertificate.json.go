package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewCertificate(ctx, "certificate", &apimanagement.CertificateArgs{
			CertificateId:     pulumi.String("tempcert"),
			Data:              pulumi.String("****************Base 64 Encoded Certificate *******************************"),
			Password:          pulumi.String("****Certificate Password******"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
