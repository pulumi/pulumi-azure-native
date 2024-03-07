package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewWebAppAuthSettings(ctx, "webAppAuthSettings", &web.WebAppAuthSettingsArgs{
			AllowedExternalRedirectUrls: pulumi.StringArray{
				pulumi.String("sitef6141.customdomain.net"),
				pulumi.String("sitef6141.customdomain.info"),
			},
			ClientId:                    pulumi.String("42d795a9-8abb-4d06-8534-39528af40f8e.apps.googleusercontent.com"),
			DefaultProvider:             web.BuiltInAuthenticationProviderGoogle,
			Enabled:                     pulumi.Bool(true),
			Name:                        pulumi.String("sitef6141"),
			ResourceGroupName:           pulumi.String("testrg123"),
			RuntimeVersion:              pulumi.String("~1"),
			TokenRefreshExtensionHours:  pulumi.Float64(120),
			TokenStoreEnabled:           pulumi.Bool(true),
			UnauthenticatedClientAction: web.UnauthenticatedClientActionRedirectToLoginPage,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
