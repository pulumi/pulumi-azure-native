package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewBlobServiceProperties(ctx, "blobServiceProperties", &storage.BlobServicePropertiesArgs{
			AccountName:      pulumi.String("sto8607"),
			BlobServicesName: pulumi.String("default"),
			LastAccessTimeTrackingPolicy: &storage.LastAccessTimeTrackingPolicyArgs{
				BlobType: pulumi.StringArray{
					pulumi.String("blockBlob"),
				},
				Enable:                    pulumi.Bool(true),
				Name:                      pulumi.String("AccessTimeTracking"),
				TrackingGranularityInDays: pulumi.Int(1),
			},
			ResourceGroupName: pulumi.String("res4410"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
