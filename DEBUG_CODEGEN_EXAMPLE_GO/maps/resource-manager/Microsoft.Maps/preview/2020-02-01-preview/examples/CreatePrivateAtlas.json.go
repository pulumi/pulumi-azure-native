package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/maps/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := maps.NewPrivateAtlase(ctx, "privateAtlase", &maps.PrivateAtlaseArgs{
			AccountName:       pulumi.String("myMapsAccount"),
			Location:          pulumi.String("unitedstates"),
			PrivateAtlasName:  pulumi.String("myPrivateAtlas"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Tags: pulumi.StringMap{
				"test": pulumi.String("true"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
