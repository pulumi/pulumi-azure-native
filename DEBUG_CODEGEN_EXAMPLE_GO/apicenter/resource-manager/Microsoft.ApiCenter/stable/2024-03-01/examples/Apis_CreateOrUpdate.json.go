package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apicenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apicenter.NewApi(ctx, "api", &apicenter.ApiArgs{
			ApiName: pulumi.String("echo-api"),
			CustomProperties: pulumi.Any{
				Author: "John Doe",
			},
			Description: pulumi.String("A simple HTTP request/response service."),
			ExternalDocumentation: []apicenter.ExternalDocumentationArgs{
				{
					Title: pulumi.String("Onboarding docs"),
					Url:   pulumi.String("https://docs.contoso.com"),
				},
			},
			Kind: pulumi.String("rest"),
			License: &apicenter.LicenseArgs{
				Url: pulumi.String("https://contoso.com/license"),
			},
			ResourceGroupName: pulumi.String("contoso-resources"),
			ServiceName:       pulumi.String("contoso"),
			TermsOfService: &apicenter.TermsOfServiceArgs{
				Url: pulumi.String("https://contoso.com/terms-of-service"),
			},
			Title:         pulumi.String("Echo API"),
			WorkspaceName: pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
