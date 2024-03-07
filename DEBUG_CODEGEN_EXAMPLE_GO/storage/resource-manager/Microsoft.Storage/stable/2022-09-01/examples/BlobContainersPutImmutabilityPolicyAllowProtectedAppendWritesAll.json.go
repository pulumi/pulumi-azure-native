package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewBlobContainerImmutabilityPolicy(ctx, "blobContainerImmutabilityPolicy", &storage.BlobContainerImmutabilityPolicyArgs{
			AccountName:                           pulumi.String("sto7069"),
			AllowProtectedAppendWritesAll:         pulumi.Bool(true),
			ContainerName:                         pulumi.String("container6397"),
			ImmutabilityPeriodSinceCreationInDays: pulumi.Int(3),
			ImmutabilityPolicyName:                pulumi.String("default"),
			ResourceGroupName:                     pulumi.String("res1782"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
