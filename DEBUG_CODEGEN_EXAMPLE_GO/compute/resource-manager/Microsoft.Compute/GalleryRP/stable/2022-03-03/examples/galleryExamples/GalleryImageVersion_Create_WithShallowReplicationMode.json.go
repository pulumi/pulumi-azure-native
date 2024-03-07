package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewGalleryImageVersion(ctx, "galleryImageVersion", &compute.GalleryImageVersionArgs{
			GalleryImageName:        pulumi.String("myGalleryImageName"),
			GalleryImageVersionName: pulumi.String("1.0.0"),
			GalleryName:             pulumi.String("myGalleryName"),
			Location:                pulumi.String("West US"),
			PublishingProfile: compute.GalleryImageVersionPublishingProfileResponse{
				ReplicationMode: pulumi.String("Shallow"),
				TargetRegions: compute.TargetRegionArray{
					&compute.TargetRegionArgs{
						ExcludeFromLatest:    pulumi.Bool(false),
						Name:                 pulumi.String("West US"),
						RegionalReplicaCount: pulumi.Int(1),
					},
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SafetyProfile: &compute.GalleryImageVersionSafetyProfileArgs{
				AllowDeletionOfReplicatedLocations: pulumi.Bool(false),
			},
			StorageProfile: compute.GalleryImageVersionStorageProfileResponse{
				Source: &compute.GalleryArtifactVersionFullSourceArgs{
					Id: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
