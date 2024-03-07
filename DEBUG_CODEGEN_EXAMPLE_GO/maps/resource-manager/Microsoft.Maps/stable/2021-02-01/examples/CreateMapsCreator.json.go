package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/maps/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := maps.NewCreator(ctx, "creator", &maps.CreatorArgs{
			AccountName: pulumi.String("myMapsAccount"),
			CreatorName: pulumi.String("myCreator"),
			Location:    pulumi.String("eastus2"),
			Properties: &maps.CreatorPropertiesArgs{
				StorageUnits: pulumi.Int(5),
			},
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
