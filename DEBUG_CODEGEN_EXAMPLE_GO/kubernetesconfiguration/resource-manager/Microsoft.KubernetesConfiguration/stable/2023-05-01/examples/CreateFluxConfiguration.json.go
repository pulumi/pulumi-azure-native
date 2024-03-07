package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kubernetesconfiguration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kubernetesconfiguration.NewFluxConfiguration(ctx, "fluxConfiguration", &kubernetesconfiguration.FluxConfigurationArgs{
			ClusterName:           pulumi.String("clusterName1"),
			ClusterResourceName:   pulumi.String("connectedClusters"),
			ClusterRp:             pulumi.String("Microsoft.Kubernetes"),
			FluxConfigurationName: pulumi.String("srs-fluxconfig"),
			GitRepository: &kubernetesconfiguration.GitRepositoryDefinitionArgs{
				HttpsCACert: pulumi.String("ZXhhbXBsZWNlcnRpZmljYXRl"),
				RepositoryRef: &kubernetesconfiguration.RepositoryRefDefinitionArgs{
					Branch: pulumi.String("master"),
				},
				SyncIntervalInSeconds: pulumi.Float64(600),
				TimeoutInSeconds:      pulumi.Float64(600),
				Url:                   pulumi.String("https://github.com/Azure/arc-k8s-demo"),
			},
			Kustomizations: kubernetesconfiguration.KustomizationDefinitionMap{
				"srs-kustomization1": &kubernetesconfiguration.KustomizationDefinitionArgs{
					DependsOn: pulumi.StringArray{},
					Path:      pulumi.String("./test/path"),
					PostBuild: &kubernetesconfiguration.PostBuildDefinitionArgs{
						Substitute: pulumi.StringMap{
							"cluster_env":   pulumi.String("prod"),
							"replica_count": pulumi.String("2"),
						},
						SubstituteFrom: kubernetesconfiguration.SubstituteFromDefinitionArray{
							&kubernetesconfiguration.SubstituteFromDefinitionArgs{
								Kind:     pulumi.String("ConfigMap"),
								Name:     pulumi.String("cluster-test"),
								Optional: pulumi.Bool(true),
							},
						},
					},
					SyncIntervalInSeconds: pulumi.Float64(600),
					TimeoutInSeconds:      pulumi.Float64(600),
					Wait:                  pulumi.Bool(true),
				},
				"srs-kustomization2": &kubernetesconfiguration.KustomizationDefinitionArgs{
					DependsOn: pulumi.StringArray{
						pulumi.String("srs-kustomization1"),
					},
					Path: pulumi.String("./other/test/path"),
					PostBuild: &kubernetesconfiguration.PostBuildDefinitionArgs{
						SubstituteFrom: kubernetesconfiguration.SubstituteFromDefinitionArray{
							&kubernetesconfiguration.SubstituteFromDefinitionArgs{
								Kind:     pulumi.String("ConfigMap"),
								Name:     pulumi.String("cluster-values"),
								Optional: pulumi.Bool(true),
							},
							&kubernetesconfiguration.SubstituteFromDefinitionArgs{
								Kind:     pulumi.String("Secret"),
								Name:     pulumi.String("secret-name"),
								Optional: pulumi.Bool(false),
							},
						},
					},
					Prune:                  pulumi.Bool(false),
					RetryIntervalInSeconds: pulumi.Float64(600),
					SyncIntervalInSeconds:  pulumi.Float64(600),
					TimeoutInSeconds:       pulumi.Float64(600),
					Wait:                   pulumi.Bool(false),
				},
			},
			Namespace:                  pulumi.String("srs-namespace"),
			ReconciliationWaitDuration: pulumi.String("PT30M"),
			ResourceGroupName:          pulumi.String("rg1"),
			Scope:                      pulumi.String("cluster"),
			SourceKind:                 pulumi.String("GitRepository"),
			Suspend:                    pulumi.Bool(false),
			WaitForReconciliation:      pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
