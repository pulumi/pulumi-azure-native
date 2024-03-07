package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apicenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apicenter.NewEnvironment(ctx, "environment", &apicenter.EnvironmentArgs{
			Description:     pulumi.String("The primary Azure API Management service for the European division of Contoso."),
			EnvironmentName: pulumi.String("public"),
			Kind:            pulumi.String("production"),
			Onboarding: &apicenter.OnboardingArgs{
				DeveloperPortalUri: pulumi.StringArray{
					pulumi.String("https://developer.contoso.com"),
				},
				Instructions: pulumi.String("Sign in or sign up in the specified developer portal to request API access. You must complete the internal privacy training for your account to be approved."),
			},
			ResourceGroupName: pulumi.String("contoso-resources"),
			Server: &apicenter.EnvironmentServerArgs{
				ManagementPortalUri: pulumi.StringArray{
					pulumi.String("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ApiManagement/service/contoso"),
				},
				Type: pulumi.String("Azure API Management"),
			},
			ServiceName:   pulumi.String("contoso"),
			Title:         pulumi.String("Contoso Europe Azure API Management"),
			WorkspaceName: pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
