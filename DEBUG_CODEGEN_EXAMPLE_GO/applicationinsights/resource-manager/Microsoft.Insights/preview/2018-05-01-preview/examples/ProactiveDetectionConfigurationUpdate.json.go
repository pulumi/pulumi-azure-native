package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewProactiveDetectionConfiguration(ctx, "proactiveDetectionConfiguration", &insights.ProactiveDetectionConfigurationArgs{
			ConfigurationId: pulumi.String("slowpageloadtime"),
			CustomEmails: pulumi.StringArray{
				pulumi.String("foo@microsoft.com"),
				pulumi.String("foo2@microsoft.com"),
			},
			Enabled:           pulumi.Bool(true),
			Location:          pulumi.String("South Central US"),
			Name:              pulumi.String("slowpageloadtime"),
			ResourceGroupName: pulumi.String("my-resource-group"),
			ResourceName:      pulumi.String("my-component"),
			RuleDefinitions: &insights.ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs{
				Description:                pulumi.String("Smart Detection rules notify you of performance anomaly issues."),
				DisplayName:                pulumi.String("Slow page load time"),
				HelpUrl:                    pulumi.String("https://docs.microsoft.com/en-us/azure/application-insights/app-insights-proactive-performance-diagnostics"),
				IsEnabledByDefault:         pulumi.Bool(true),
				IsHidden:                   pulumi.Bool(false),
				IsInPreview:                pulumi.Bool(false),
				Name:                       pulumi.String("slowpageloadtime"),
				SupportsEmailNotifications: pulumi.Bool(true),
			},
			SendEmailsToSubscriptionOwners: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
