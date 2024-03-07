package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewApiPortalCustomDomain(ctx, "apiPortalCustomDomain", &appplatform.ApiPortalCustomDomainArgs{
			ApiPortalName: pulumi.String("default"),
			DomainName:    pulumi.String("myDomainName"),
			Properties: &appplatform.ApiPortalCustomDomainPropertiesArgs{
				Thumbprint: pulumi.String("*"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
