package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devcenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devcenter.NewDevBoxDefinition(ctx, "devBoxDefinition", &devcenter.DevBoxDefinitionArgs{
			DevBoxDefinitionName: pulumi.String("WebDevBox"),
			DevCenterName:        pulumi.String("Contoso"),
			HibernateSupport:     pulumi.String("Enabled"),
			ImageReference: &devcenter.ImageReferenceArgs{
				Id: pulumi.String("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/Example/providers/Microsoft.DevCenter/devcenters/Contoso/galleries/contosogallery/images/exampleImage/version/1.0.0"),
			},
			Location:          pulumi.String("centralus"),
			ResourceGroupName: pulumi.String("rg1"),
			Sku: &devcenter.SkuArgs{
				Name: pulumi.String("Preview"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
