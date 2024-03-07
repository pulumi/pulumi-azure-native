package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/providerhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := providerhub.NewResourceTypeRegistration(ctx, "resourceTypeRegistration", &providerhub.ResourceTypeRegistrationArgs{
			Properties: &providerhub.ResourceTypeRegistrationPropertiesArgs{
				Endpoints: providerhub.ResourceTypeEndpointArray{
					&providerhub.ResourceTypeEndpointArgs{
						ApiVersions: pulumi.StringArray{
							pulumi.String("2020-06-01-preview"),
						},
						Locations: pulumi.StringArray{
							pulumi.String("West US"),
							pulumi.String("East US"),
							pulumi.String("North Europe"),
						},
						RequiredFeatures: pulumi.StringArray{
							pulumi.String("<feature flag>"),
						},
					},
				},
				Management: &providerhub.ResourceTypeRegistrationPropertiesManagementArgs{
					IncidentContactEmail:   pulumi.String("helpme@contoso.com"),
					IncidentRoutingService: pulumi.String(""),
					IncidentRoutingTeam:    pulumi.String(""),
					ManifestOwners: pulumi.StringArray{
						pulumi.String("SPARTA-PlatformServiceAdministrator"),
					},
					ResourceAccessPolicy: pulumi.String("NotSpecified"),
					ServiceTreeInfos: providerhub.ServiceTreeInfoArray{
						&providerhub.ServiceTreeInfoArgs{
							ComponentId: pulumi.String("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
							Readiness:   pulumi.String("InDevelopment"),
							ServiceId:   pulumi.String("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
						},
					},
				},
				OpenApiConfiguration: &providerhub.OpenApiConfigurationArgs{
					Validation: &providerhub.OpenApiValidationArgs{
						AllowNoncompliantCollectionResponse: pulumi.Bool(true),
					},
				},
				Regionality: pulumi.String("Regional"),
				ResourceConcurrencyControlOptions: providerhub.ResourceConcurrencyControlOptionMap{
					"patch": &providerhub.ResourceConcurrencyControlOptionArgs{
						Policy: pulumi.String("SynchronizeBeginExtension"),
					},
					"post": &providerhub.ResourceConcurrencyControlOptionArgs{
						Policy: pulumi.String("SynchronizeBeginExtension"),
					},
					"put": &providerhub.ResourceConcurrencyControlOptionArgs{
						Policy: pulumi.String("SynchronizeBeginExtension"),
					},
				},
				ResourceGraphConfiguration: &providerhub.ResourceTypeRegistrationPropertiesResourceGraphConfigurationArgs{
					ApiVersion: pulumi.String("2019-01-01"),
					Enabled:    pulumi.Bool(true),
				},
				RoutingType: pulumi.String("Default"),
				SwaggerSpecifications: providerhub.SwaggerSpecificationArray{
					&providerhub.SwaggerSpecificationArgs{
						ApiVersions: pulumi.StringArray{
							pulumi.String("2020-06-01-preview"),
						},
						SwaggerSpecFolderUri: pulumi.String("https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/"),
					},
				},
			},
			ProviderNamespace: pulumi.String("Microsoft.Contoso"),
			ResourceType:      pulumi.String("employees"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
