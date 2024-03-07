package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityandcompliance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityandcompliance.NewPrivateLinkServicesForSCCPowershell(ctx, "privateLinkServicesForSCCPowershell", &securityandcompliance.PrivateLinkServicesForSCCPowershellArgs{
			Identity: &securityandcompliance.ServicesResourceIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Kind:     securityandcompliance.Kind_Fhir_R4,
			Location: pulumi.String("westus2"),
			Properties: &securityandcompliance.ServicesPropertiesArgs{
				AccessPolicies: securityandcompliance.ServiceAccessPolicyEntryArray{
					&securityandcompliance.ServiceAccessPolicyEntryArgs{
						ObjectId: pulumi.String("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
					},
					&securityandcompliance.ServiceAccessPolicyEntryArgs{
						ObjectId: pulumi.String("5b307da8-43d4-492b-8b66-b0294ade872f"),
					},
				},
				AuthenticationConfiguration: &securityandcompliance.ServiceAuthenticationConfigurationInfoArgs{
					Audience:          pulumi.String("https://azurehealthcareapis.com"),
					Authority:         pulumi.String("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
					SmartProxyEnabled: pulumi.Bool(true),
				},
				CorsConfiguration: &securityandcompliance.ServiceCorsConfigurationInfoArgs{
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
				CosmosDbConfiguration: &securityandcompliance.ServiceCosmosDbConfigurationInfoArgs{
					KeyVaultKeyUri:  pulumi.String("https://my-vault.vault.azure.net/keys/my-key"),
					OfferThroughput: pulumi.Float64(1000),
				},
				ExportConfiguration: &securityandcompliance.ServiceExportConfigurationInfoArgs{
					StorageAccountName: pulumi.String("existingStorageAccount"),
				},
				PrivateEndpointConnections: securityandcompliance.PrivateEndpointConnectionArray{},
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
