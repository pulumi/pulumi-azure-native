package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/healthcareapis/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := healthcareapis.NewFhirService(ctx, "fhirService", &healthcareapis.FhirServiceArgs{
			AccessPolicies: []healthcareapis.FhirServiceAccessPolicyEntryArgs{
				{
					ObjectId: pulumi.String("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
				},
				{
					ObjectId: pulumi.String("5b307da8-43d4-492b-8b66-b0294ade872f"),
				},
			},
			AcrConfiguration: &healthcareapis.FhirServiceAcrConfigurationArgs{
				LoginServers: pulumi.StringArray{
					pulumi.String("test1.azurecr.io"),
				},
			},
			AuthenticationConfiguration: &healthcareapis.FhirServiceAuthenticationConfigurationArgs{
				Audience:          pulumi.String("https://azurehealthcareapis.com"),
				Authority:         pulumi.String("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
				SmartProxyEnabled: pulumi.Bool(true),
			},
			CorsConfiguration: &healthcareapis.FhirServiceCorsConfigurationArgs{
				AllowCredentials: pulumi.Bool(false),
				Headers: pulumi.StringArray{
					pulumi.String("*"),
				},
				MaxAge: pulumi.Int(1440),
				Methods: pulumi.StringArray{
					pulumi.String("DELETE"),
					pulumi.String("GET"),
					pulumi.String("OPTIONS"),
					pulumi.String("PATCH"),
					pulumi.String("POST"),
					pulumi.String("PUT"),
				},
				Origins: pulumi.StringArray{
					pulumi.String("*"),
				},
			},
			ExportConfiguration: &healthcareapis.FhirServiceExportConfigurationArgs{
				StorageAccountName: pulumi.String("existingStorageAccount"),
			},
			FhirServiceName: pulumi.String("fhirservice1"),
			Identity: &healthcareapis.ServiceManagedIdentityIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			ImplementationGuidesConfiguration: &healthcareapis.ImplementationGuidesConfigurationArgs{
				UsCoreMissingData: pulumi.Bool(false),
			},
			ImportConfiguration: &healthcareapis.FhirServiceImportConfigurationArgs{
				Enabled:              pulumi.Bool(false),
				InitialImportMode:    pulumi.Bool(false),
				IntegrationDataStore: pulumi.String("existingStorageAccount"),
			},
			Kind:              pulumi.String("fhir-R4"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("testRG"),
			Tags: pulumi.StringMap{
				"additionalProp1": pulumi.String("string"),
				"additionalProp2": pulumi.String("string"),
				"additionalProp3": pulumi.String("string"),
			},
			WorkspaceName: pulumi.String("workspace1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
