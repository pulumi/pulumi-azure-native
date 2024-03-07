package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewStaticSiteCustomDomain(ctx, "staticSiteCustomDomain", &web.StaticSiteCustomDomainArgs{
			DomainName:        pulumi.String("custom.domain.net"),
			Name:              pulumi.String("testStaticSite0"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
