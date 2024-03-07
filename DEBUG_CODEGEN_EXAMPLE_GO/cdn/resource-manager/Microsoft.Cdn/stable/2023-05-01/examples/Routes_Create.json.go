package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewRoute(ctx, "route", &cdn.RouteArgs{
			CacheConfiguration: cdn.AfdRouteCacheConfigurationResponse{
				CompressionSettings: &cdn.CompressionSettingsArgs{
					ContentTypesToCompress: pulumi.StringArray{
						pulumi.String("text/html"),
						pulumi.String("application/octet-stream"),
					},
					IsCompressionEnabled: pulumi.Bool(true),
				},
				QueryParameters:            pulumi.String("querystring=test"),
				QueryStringCachingBehavior: pulumi.String("IgnoreSpecifiedQueryStrings"),
			},
			CustomDomains: []cdn.ActivatedResourceReferenceArgs{
				{
					Id: pulumi.String("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/customDomains/domain1"),
				},
			},
			EnabledState:        pulumi.String("Enabled"),
			EndpointName:        pulumi.String("endpoint1"),
			ForwardingProtocol:  pulumi.String("MatchRequest"),
			HttpsRedirect:       pulumi.String("Enabled"),
			LinkToDefaultDomain: pulumi.String("Enabled"),
			OriginGroup: &cdn.ResourceReferenceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1"),
			},
			PatternsToMatch: pulumi.StringArray{
				pulumi.String("/*"),
			},
			ProfileName:       pulumi.String("profile1"),
			ResourceGroupName: pulumi.String("RG"),
			RouteName:         pulumi.String("route1"),
			RuleSets: []cdn.ResourceReferenceArgs{
				{
					Id: pulumi.String("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1"),
				},
			},
			SupportedProtocols: pulumi.StringArray{
				pulumi.String("Https"),
				pulumi.String("Http"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
