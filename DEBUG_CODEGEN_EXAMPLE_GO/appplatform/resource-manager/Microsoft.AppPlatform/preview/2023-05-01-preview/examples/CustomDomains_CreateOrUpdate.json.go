package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewCustomDomain(ctx, "customDomain", &appplatform.CustomDomainArgs{
			AppName:    pulumi.String("myapp"),
			DomainName: pulumi.String("mydomain.com"),
			Properties: &appplatform.CustomDomainPropertiesArgs{
				CertName:   pulumi.String("mycert"),
				Thumbprint: pulumi.String("934367bf1c97033f877db0f15cb1b586957d3133"),
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
