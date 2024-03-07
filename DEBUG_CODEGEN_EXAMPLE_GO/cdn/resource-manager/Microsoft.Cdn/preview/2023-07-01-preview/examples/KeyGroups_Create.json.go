package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewKeyGroup(ctx, "keyGroup", &cdn.KeyGroupArgs{
			KeyGroupName: pulumi.String("kg1"),
			KeyReferences: []cdn.ResourceReferenceArgs{
				{
					Id: pulumi.String("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/secrets/secret1"),
				},
				{
					Id: pulumi.String("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/secrets/secret2"),
				},
				{
					Id: pulumi.String("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/secrets/secret3"),
				},
			},
			ProfileName:       pulumi.String("profile1"),
			ResourceGroupName: pulumi.String("RG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
