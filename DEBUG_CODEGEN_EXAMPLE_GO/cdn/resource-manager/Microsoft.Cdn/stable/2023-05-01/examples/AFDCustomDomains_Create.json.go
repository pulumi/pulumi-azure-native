package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewAFDCustomDomain(ctx, "afdCustomDomain", &cdn.AFDCustomDomainArgs{
			AzureDnsZone: &cdn.ResourceReferenceArgs{
				Id: pulumi.String(""),
			},
			CustomDomainName:  pulumi.String("domain1"),
			HostName:          pulumi.String("www.someDomain.net"),
			ProfileName:       pulumi.String("profile1"),
			ResourceGroupName: pulumi.String("RG"),
			TlsSettings: &cdn.AFDDomainHttpsParametersArgs{
				CertificateType:   pulumi.String("ManagedCertificate"),
				MinimumTlsVersion: cdn.AfdMinimumTlsVersionTLS12,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
