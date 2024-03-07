package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewDisk(ctx, "disk", &compute.DiskArgs{
			CreationData: &compute.CreationDataArgs{
				CreateOption: pulumi.String("Empty"),
			},
			DiskAccessId:        pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/{existing-diskAccess-name}"),
			DiskName:            pulumi.String("myDisk"),
			DiskSizeGB:          pulumi.Int(200),
			Location:            pulumi.String("West US"),
			NetworkAccessPolicy: pulumi.String("AllowPrivate"),
			ResourceGroupName:   pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
