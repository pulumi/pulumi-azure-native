package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewArtifactManifest(ctx, "artifactManifest", &hybridnetwork.ArtifactManifestArgs{
			ArtifactManifestName: pulumi.String("TestManifest"),
			ArtifactStoreName:    pulumi.String("TestArtifactStore"),
			Location:             pulumi.String("eastus"),
			Properties: &hybridnetwork.ArtifactManifestPropertiesFormatArgs{
				Artifacts: hybridnetwork.ManifestArtifactFormatArray{
					&hybridnetwork.ManifestArtifactFormatArgs{
						ArtifactName:    pulumi.String("fed-rbac"),
						ArtifactType:    pulumi.String("OCIArtifact"),
						ArtifactVersion: pulumi.String("1.0.0"),
					},
					&hybridnetwork.ManifestArtifactFormatArgs{
						ArtifactName:    pulumi.String("nginx"),
						ArtifactType:    pulumi.String("OCIArtifact"),
						ArtifactVersion: pulumi.String("v1"),
					},
				},
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
