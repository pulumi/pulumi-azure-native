package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewNetworkFunctionDefinitionVersion(ctx, "networkFunctionDefinitionVersion", &hybridnetwork.NetworkFunctionDefinitionVersionArgs{
			Location:                             pulumi.String("eastus"),
			NetworkFunctionDefinitionGroupName:   pulumi.String("TestNetworkFunctionDefinitionGroupName"),
			NetworkFunctionDefinitionVersionName: pulumi.String("1.0.0"),
			Properties: hybridnetwork.ContainerizedNetworkFunctionDefinitionVersion{
				DeployParameters: "{\"type\":\"object\",\"properties\":{\"releaseName\":{\"type\":\"string\"},\"namespace\":{\"type\":\"string\"}}}",
				NetworkFunctionTemplate: hybridnetwork.AzureArcKubernetesNetworkFunctionTemplate{
					NetworkFunctionApplications: []hybridnetwork.AzureArcKubernetesHelmApplication{
						{
							ArtifactProfile: {
								ArtifactStore: {
									Id: "/subscriptions/subid/resourcegroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/testArtifactStore",
								},
								HelmArtifactProfile: {
									HelmPackageName:         "fed-rbac",
									HelmPackageVersionRange: "~2.1.3",
									ImagePullSecretsValuesPaths: []string{
										"global.imagePullSecrets",
									},
									RegistryValuesPaths: []string{
										"global.registry.docker.repoPath",
									},
								},
							},
							ArtifactType: "HelmPackage",
							DependsOnProfile: {
								InstallDependsOn:   []interface{}{},
								UninstallDependsOn: []interface{}{},
								UpdateDependsOn:    []interface{}{},
							},
							DeployParametersMappingRuleProfile: {
								ApplicationEnablement: "Enabled",
								HelmMappingRuleProfile: {
									HelmPackageVersion: "2.1.3",
									Options: {
										InstallOptions: {
											Atomic:  "true",
											Timeout: "30",
											Wait:    "waitValue",
										},
										UpgradeOptions: {
											Atomic:  "true",
											Timeout: "30",
											Wait:    "waitValue",
										},
									},
									ReleaseName:      "{deployParameters.releaseName}",
									ReleaseNamespace: "{deployParameters.namesapce}",
									Values:           "",
								},
							},
							Name: "fedrbac",
						},
					},
					NfviType: "AzureArcKubernetes",
				},
				NetworkFunctionType: "ContainerizedNetworkFunction",
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
