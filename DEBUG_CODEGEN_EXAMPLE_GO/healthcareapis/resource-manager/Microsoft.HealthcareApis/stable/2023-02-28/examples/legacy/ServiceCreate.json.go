package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/healthcareapis/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := healthcareapis.NewService(ctx, "service", &healthcareapis.ServiceArgs{
			Identity: &healthcareapis.ServicesResourceIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Kind:     healthcareapis.Kind_Fhir_R4,
			Location: pulumi.String("westus2"),
			Properties: healthcareapis.ServicesPropertiesResponse{
				AccessPolicies: healthcareapis.ServiceAccessPolicyEntryArray{
					&healthcareapis.ServiceAccessPolicyEntryArgs{
						ObjectId: pulumi.String("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
					},
					&healthcareapis.ServiceAccessPolicyEntryArgs{
						ObjectId: pulumi.String("5b307da8-43d4-492b-8b66-b0294ade872f"),
					},
				},
				AuthenticationConfiguration: &healthcareapis.ServiceAuthenticationConfigurationInfoArgs{
					Audience:          pulumi.String("https://azurehealthcareapis.com"),
					Authority:         pulumi.String("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
					SmartProxyEnabled: pulumi.Bool(true),
				},
				CorsConfiguration: &healthcareapis.ServiceCorsConfigurationInfoArgs{
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
				CosmosDbConfiguration: &healthcareapis.ServiceCosmosDbConfigurationInfoArgs{
					KeyVaultKeyUri:  pulumi.String("https://my-vault.vault.azure.net/keys/my-key"),
					OfferThroughput: pulumi.Int(1000),
				},
				ExportConfiguration: &healthcareapis.ServiceExportConfigurationInfoArgs{
					StorageAccountName: pulumi.String("existingStorageAccount"),
				},
				PrivateEndpointConnections: healthcareapis.PrivateEndpointConnectionTypeArray{},
				PublicNetworkAccess:        pulumi.String("Disabled"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ResourceName:      pulumi.String("service1"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
