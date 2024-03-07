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
			DeleteRetentionPolicy: &storage.DeleteRetentionPolicyArgs{
				AllowPermanentDelete: pulumi.Bool(true),
				Days:                 pulumi.Int(300),
				Enabled:              pulumi.Bool(true),
			},
			IsVersioningEnabled: pulumi.Bool(true),
			ResourceGroupName:   pulumi.String("res4410"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
