package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewReplicationStorageClassificationMapping(ctx, "replicationStorageClassificationMapping", &recoveryservices.ReplicationStorageClassificationMappingArgs{
			FabricName: pulumi.String("2a48e3770ac08aa2be8bfbd94fcfb1cbf2dcc487b78fb9d3bd778304441b06a0"),
			Properties: &recoveryservices.StorageMappingInputPropertiesArgs{
				TargetStorageClassificationId: pulumi.String("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/2a48e3770ac08aa2be8bfbd94fcfb1cbf2dcc487b78fb9d3bd778304441b06a0/replicationStorageClassifications/8891569e-aaef-4a46-a4a0-78c14f2d7b09"),
			},
			ResourceGroupName:                pulumi.String("resourceGroupPS1"),
			ResourceName:                     pulumi.String("vault1"),
			StorageClassificationMappingName: pulumi.String("testStorageMapping"),
			StorageClassificationName:        pulumi.String("8891569e-aaef-4a46-a4a0-78c14f2d7b09"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
