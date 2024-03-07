package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/deviceregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := deviceregistry.NewAssetEndpointProfile(ctx, "assetEndpointProfile", &deviceregistry.AssetEndpointProfileArgs{
			AssetEndpointProfileName: pulumi.String("my-assetendpointprofile"),
			ExtendedLocation: &deviceregistry.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
				Type: pulumi.String("CustomLocation"),
			},
			Location:          pulumi.String("West Europe"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Tags: pulumi.StringMap{
				"site": pulumi.String("building-1"),
			},
			TargetAddress: pulumi.String("https://www.example.com/myTargetAddress"),
			UserAuthentication: &deviceregistry.UserAuthenticationArgs{
				Mode: pulumi.String("Anonymous"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
