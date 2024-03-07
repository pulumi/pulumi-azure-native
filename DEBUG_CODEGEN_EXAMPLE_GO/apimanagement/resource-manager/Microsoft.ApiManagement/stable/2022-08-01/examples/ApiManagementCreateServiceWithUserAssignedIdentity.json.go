package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiManagementService(ctx, "apiManagementService", &apimanagement.ApiManagementServiceArgs{
			Identity: &apimanagement.ApiManagementServiceIdentityArgs{
				Type: pulumi.String("UserAssigned"),
				UserAssignedIdentities: apimanagement.UserIdentityPropertiesMap{
					"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/apimService1": nil,
				},
			},
			Location:          pulumi.String("West US"),
			PublisherEmail:    pulumi.String("apim@autorestsdk.com"),
			PublisherName:     pulumi.String("autorestsdk"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Sku: &apimanagement.ApiManagementServiceSkuPropertiesArgs{
				Capacity: pulumi.Int(0),
				Name:     pulumi.String("Consumption"),
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
				"tag3": pulumi.String("value3"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
