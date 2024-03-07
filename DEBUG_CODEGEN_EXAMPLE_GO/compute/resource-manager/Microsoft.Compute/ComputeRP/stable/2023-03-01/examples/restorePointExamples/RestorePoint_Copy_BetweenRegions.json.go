package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewRestorePoint(ctx, "restorePoint", &compute.RestorePointArgs{
			ResourceGroupName:          pulumi.String("myResourceGroup"),
			RestorePointCollectionName: pulumi.String("rpcName"),
			RestorePointName:           pulumi.String("rpName"),
			SourceRestorePoint: &compute.ApiEntityReferenceArgs{
				Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/sourceRpcName/restorePoints/sourceRpName"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
