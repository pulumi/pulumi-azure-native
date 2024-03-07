package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiManagementService(ctx, "apiManagementService", &apimanagement.ApiManagementServiceArgs{
			AdditionalLocations: apimanagement.AdditionalLocationArray{
				&apimanagement.AdditionalLocationArgs{
					DisableGateway: pulumi.Bool(true),
					Location:       pulumi.String("East US"),
					Sku: &apimanagement.ApiManagementServiceSkuPropertiesArgs{
						Capacity: pulumi.Int(1),
						Name:     pulumi.String("Premium"),
					},
				},
			},
			ApiVersionConstraint: &apimanagement.ApiVersionConstraintArgs{
				MinApiVersion: pulumi.String("2019-01-01"),
			},
			HostnameConfigurations: apimanagement.HostnameConfigurationArray{
				&apimanagement.HostnameConfigurationArgs{
					CertificatePassword: pulumi.String("Password"),
					DefaultSslBinding:   pulumi.Bool(true),
					EncodedCertificate:  pulumi.String("****** Base 64 Encoded Certificate ************"),
					HostName:            pulumi.String("gateway1.msitesting.net"),
					Type:                pulumi.String("Proxy"),
				},
				&apimanagement.HostnameConfigurationArgs{
					CertificatePassword: pulumi.String("Password"),
					EncodedCertificate:  pulumi.String("****** Base 64 Encoded Certificate ************"),
					HostName:            pulumi.String("mgmt.msitesting.net"),
					Type:                pulumi.String("Management"),
				},
				&apimanagement.HostnameConfigurationArgs{
					CertificatePassword: pulumi.String("Password"),
					EncodedCertificate:  pulumi.String("****** Base 64 Encoded Certificate ************"),
					HostName:            pulumi.String("portal1.msitesting.net"),
					Type:                pulumi.String("Portal"),
				},
			},
			Location:          pulumi.String("West US"),
			PublisherEmail:    pulumi.String("apim@autorestsdk.com"),
			PublisherName:     pulumi.String("autorestsdk"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Sku: &apimanagement.ApiManagementServiceSkuPropertiesArgs{
				Capacity: pulumi.Int(1),
				Name:     pulumi.String("Premium"),
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
				"tag3": pulumi.String("value3"),
			},
			VirtualNetworkType: pulumi.String("None"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
