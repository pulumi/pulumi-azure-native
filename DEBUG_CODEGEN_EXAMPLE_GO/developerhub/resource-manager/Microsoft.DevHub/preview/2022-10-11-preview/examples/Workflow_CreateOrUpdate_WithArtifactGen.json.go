package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devhub.NewWorkflow(ctx, "workflow", &devhub.WorkflowArgs{
			Acr: &devhub.ACRArgs{
				AcrRegistryName:   pulumi.String("registry1"),
				AcrRepositoryName: pulumi.String("repo1"),
				AcrResourceGroup:  pulumi.String("resourceGroup1"),
				AcrSubscriptionId: pulumi.String("subscriptionId1"),
			},
			AksResourceId: pulumi.String("/subscriptions/subscriptionId1/resourcegroups/resourceGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1"),
			AppName:       pulumi.String("my-app"),
			BranchName:    pulumi.String("branch1"),
			DeploymentProperties: &devhub.DeploymentPropertiesArgs{
				KubeManifestLocations: pulumi.StringArray{
					pulumi.String("/src/manifests/"),
				},
				ManifestType: pulumi.String("kube"),
				Overrides: pulumi.StringMap{
					"key1": pulumi.String("value1"),
				},
			},
			DockerBuildContext:        pulumi.String("repo1/src/"),
			Dockerfile:                pulumi.String("repo1/images/Dockerfile"),
			DockerfileGenerationMode:  pulumi.String("enabled"),
			DockerfileOutputDirectory: pulumi.String("./"),
			GenerationLanguage:        pulumi.String("javascript"),
			ImageName:                 pulumi.String("myimage"),
			ImageTag:                  pulumi.String("latest"),
			LanguageVersion:           pulumi.String("14"),
			Location:                  pulumi.String("location1"),
			ManifestGenerationMode:    pulumi.String("enabled"),
			ManifestOutputDirectory:   pulumi.String("./"),
			ManifestType:              pulumi.String("kube"),
			OidcCredentials: &devhub.GitHubWorkflowProfileOidcCredentialsArgs{
				AzureClientId: pulumi.String("12345678-3456-7890-5678-012345678901"),
				AzureTenantId: pulumi.String("66666666-3456-7890-5678-012345678901"),
			},
			Port:              pulumi.String("80"),
			RepositoryName:    pulumi.String("repo1"),
			RepositoryOwner:   pulumi.String("owner1"),
			ResourceGroupName: pulumi.String("resourceGroup1"),
			Tags: pulumi.StringMap{
				"appname": pulumi.String("testApp"),
			},
			WorkflowName: pulumi.String("workflow1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
