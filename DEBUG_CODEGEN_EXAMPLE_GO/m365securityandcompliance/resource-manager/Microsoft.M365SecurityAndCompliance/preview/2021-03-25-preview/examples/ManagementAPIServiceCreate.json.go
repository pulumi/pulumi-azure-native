package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/m365securityandcompliance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := m365securityandcompliance.NewPrivateLinkServicesForO365ManagementActivityAPI(ctx, "privateLinkServicesForO365ManagementActivityAPI", &m365securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIArgs{
			Identity: &m365securityandcompliance.ServicesResourceIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Kind:     m365securityandcompliance.Kind_Fhir_R4,
			Location: pulumi.String("westus2"),
			Properties: m365securityandcompliance.ServicesPropertiesResponse{
				AccessPolicies: m365securityandcompliance.ServiceAccessPolicyEntryArray{
					&m365securityandcompliance.ServiceAccessPolicyEntryArgs{
						ObjectId: pulumi.String("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
					},
					&m365securityandcompliance.ServiceAccessPolicyEntryArgs{
						ObjectId: pulumi.String("5b307da8-43d4-492b-8b66-b0294ade872f"),
					},
				},
				AuthenticationConfiguration: &m365securityandcompliance.ServiceAuthenticationConfigurationInfoArgs{
					Audience:          pulumi.String("https://azurehealthcareapis.com"),
					Authority:         pulumi.String("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
					SmartProxyEnabled: pulumi.Bool(true),
				},
				CorsConfiguration: &m365securityandcompliance.ServiceCorsConfigurationInfoArgs{
					AllowCredentials: pulumi.Bool(false),
					Headers: pulumi.StringArray{
						pulumi.String("*"),
					},
					MaxAge: pulumi.Float64(1440),
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
				CosmosDbConfiguration: &m365securityandcompliance.ServiceCosmosDbConfigurationInfoArgs{
					KeyVaultKeyUri:  pulumi.String("https://my-vault.vault.azure.net/keys/my-key"),
					OfferThroughput: pulumi.Float64(1000),
				},
				ExportConfiguration: &m365securityandcompliance.ServiceExportConfigurationInfoArgs{
					StorageAccountName: pulumi.String("existingStorageAccount"),
				},
				PrivateEndpointConnections: m365securityandcompliance.PrivateEndpointConnectionArray{},
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
