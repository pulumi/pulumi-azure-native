package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiManagementService(ctx, "apiManagementService", &apimanagement.ApiManagementServiceArgs{
			ApiVersionConstraint: &apimanagement.ApiVersionConstraintArgs{
				MinApiVersion: pulumi.String("2019-01-01"),
			},
			HostnameConfigurations: apimanagement.HostnameConfigurationArray{
				&apimanagement.HostnameConfigurationArgs{
					DefaultSslBinding: pulumi.Bool(true),
					HostName:          pulumi.String("gateway1.msitesting.net"),
					IdentityClientId:  pulumi.String("329419bc-adec-4dce-9568-25a6d486e468"),
					KeyVaultId:        pulumi.String("https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert"),
					Type:              pulumi.String("Proxy"),
				},
				&apimanagement.HostnameConfigurationArgs{
					HostName:         pulumi.String("mgmt.msitesting.net"),
					IdentityClientId: pulumi.String("329419bc-adec-4dce-9568-25a6d486e468"),
					KeyVaultId:       pulumi.String("https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert"),
					Type:             pulumi.String("Management"),
				},
				&apimanagement.HostnameConfigurationArgs{
					HostName:         pulumi.String("portal1.msitesting.net"),
					IdentityClientId: pulumi.String("329419bc-adec-4dce-9568-25a6d486e468"),
					KeyVaultId:       pulumi.String("https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert"),
					Type:             pulumi.String("Portal"),
				},
			},
			Identity: &apimanagement.ApiManagementServiceIdentityArgs{
				Type: pulumi.String("UserAssigned"),
				UserAssignedIdentities: apimanagement.UserIdentityPropertiesMap{
					"/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": nil,
				},
			},
			Location:          pulumi.String("North Europe"),
			PublisherEmail:    pulumi.String("apim@autorestsdk.com"),
			PublisherName:     pulumi.String("autorestsdk"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Sku: &apimanagement.ApiManagementServiceSkuPropertiesArgs{
				Capacity: pulumi.Int(1),
				Name:     pulumi.String("Premium"),
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
				"tag3": pulumi.String("value3"),
			},
			VirtualNetworkType: pulumi.String("None"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
