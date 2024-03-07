package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewArtifactStore(ctx, "artifactStore", &hybridnetwork.ArtifactStoreArgs{
			ArtifactStoreName: pulumi.String("TestArtifactStore"),
			Location:          pulumi.String("eastus"),
			Properties: hybridnetwork.ArtifactStorePropertiesFormatResponse{
				ManagedResourceGroupConfiguration: &hybridnetwork.ArtifactStorePropertiesFormatManagedResourceGroupConfigurationArgs{
					Location: pulumi.String("eastus"),
					Name:     pulumi.String("testRg"),
				},
				ReplicationStrategy: pulumi.String("SingleReplication"),
				StoreType:           pulumi.String("AzureContainerRegistry"),
			},
			PublisherName:     pulumi.String("TestPublisher"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
