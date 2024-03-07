package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kubernetesconfiguration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kubernetesconfiguration.NewFluxConfiguration(ctx, "fluxConfiguration", &kubernetesconfiguration.FluxConfigurationArgs{
			Bucket: &kubernetesconfiguration.BucketDefinitionArgs{
				AccessKey:             pulumi.String("fluxminiotest"),
				BucketName:            pulumi.String("flux"),
				SyncIntervalInSeconds: pulumi.Float64(1000),
				TimeoutInSeconds:      pulumi.Float64(1000),
				Url:                   pulumi.String("https://fluxminiotest.az.minio.io"),
			},
			ClusterName:           pulumi.String("clusterName1"),
			ClusterResourceName:   pulumi.String("connectedClusters"),
			ClusterRp:             pulumi.String("Microsoft.Kubernetes"),
			FluxConfigurationName: pulumi.String("srs-fluxconfig"),
			Kustomizations: kubernetesconfiguration.KustomizationDefinitionMap{
				"srs-kustomization1": &kubernetesconfiguration.KustomizationDefinitionArgs{
					DependsOn:             pulumi.StringArray{},
					Path:                  pulumi.String("./test/path"),
					SyncIntervalInSeconds: pulumi.Float64(600),
					TimeoutInSeconds:      pulumi.Float64(600),
				},
				"srs-kustomization2": &kubernetesconfiguration.KustomizationDefinitionArgs{
					DependsOn: pulumi.StringArray{
						pulumi.String("srs-kustomization1"),
					},
					Path:                   pulumi.String("./other/test/path"),
					Prune:                  pulumi.Bool(false),
					RetryIntervalInSeconds: pulumi.Float64(600),
					SyncIntervalInSeconds:  pulumi.Float64(600),
					TimeoutInSeconds:       pulumi.Float64(600),
				},
			},
			Namespace:         pulumi.String("srs-namespace"),
			ResourceGroupName: pulumi.String("rg1"),
			Scope:             pulumi.String("cluster"),
			SourceKind:        pulumi.String("Bucket"),
			Suspend:           pulumi.Bool(false),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
