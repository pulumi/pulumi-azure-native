package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewProfile(ctx, "profile", &cdn.ProfileArgs{
			Location:          pulumi.String("global"),
			ProfileName:       pulumi.String("profile1"),
			ResourceGroupName: pulumi.String("RG"),
			Sku: &cdn.SkuArgs{
				Name: pulumi.String("Premium_AzureFrontDoor"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
