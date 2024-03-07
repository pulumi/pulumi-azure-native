package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customproviders/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customproviders.NewAssociation(ctx, "association", &customproviders.AssociationArgs{
			AssociationName:  pulumi.String("associationName"),
			Scope:            pulumi.String("scope"),
			TargetResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/appRG/providers/Microsoft.Solutions/applications/applicationName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
