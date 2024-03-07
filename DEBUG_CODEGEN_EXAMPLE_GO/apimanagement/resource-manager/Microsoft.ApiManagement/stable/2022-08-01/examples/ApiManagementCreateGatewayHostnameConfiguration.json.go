package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewGatewayHostnameConfiguration(ctx, "gatewayHostnameConfiguration", &apimanagement.GatewayHostnameConfigurationArgs{
			CertificateId:              pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
			GatewayId:                  pulumi.String("gw1"),
			HcId:                       pulumi.String("default"),
			Hostname:                   pulumi.String("*"),
			Http2Enabled:               pulumi.Bool(true),
			NegotiateClientCertificate: pulumi.Bool(false),
			ResourceGroupName:          pulumi.String("rg1"),
			ServiceName:                pulumi.String("apimService1"),
			Tls10Enabled:               pulumi.Bool(false),
			Tls11Enabled:               pulumi.Bool(false),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
