package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewCustomDomain(ctx, "customDomain", &cdn.CustomDomainArgs{
			CustomDomainName:  pulumi.String("www-someDomain-net"),
			EndpointName:      pulumi.String("endpoint1"),
			HostName:          pulumi.String("www.someDomain.net"),
			ProfileName:       pulumi.String("profile1"),
			ResourceGroupName: pulumi.String("RG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
