package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mixedreality/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mixedreality.NewSpatialAnchorsAccount(ctx, "spatialAnchorsAccount", &mixedreality.SpatialAnchorsAccountArgs{
			AccountName:       pulumi.String("MyAccount"),
			Location:          pulumi.String("eastus2euap"),
			ResourceGroupName: pulumi.String("MyResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
