package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewRestorePointCollection(ctx, "restorePointCollection", &compute.RestorePointCollectionArgs{
			Location:                   pulumi.String("norwayeast"),
			ResourceGroupName:          pulumi.String("myResourceGroup"),
			RestorePointCollectionName: pulumi.String("myRpc"),
			Source: &compute.RestorePointCollectionSourcePropertiesArgs{
				Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
			},
			Tags: pulumi.StringMap{
				"myTag1": pulumi.String("tagValue1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
