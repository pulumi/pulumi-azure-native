package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewEndpoint(ctx, "endpoint", &cdn.EndpointArgs{
			ContentTypesToCompress: pulumi.StringArray{
				pulumi.String("text/html"),
				pulumi.String("application/octet-stream"),
			},
			DefaultOriginGroup: &cdn.ResourceReferenceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/originGroups/originGroup1"),
			},
			DeliveryPolicy: &cdn.EndpointPropertiesUpdateParametersDeliveryPolicyArgs{
				Description: pulumi.String("Test description for a policy."),
				Rules: []cdn.DeliveryRuleArgs{
					{
						Actions: pulumi.Array{
							{
								Name: "CacheExpiration",
								Parameters: {
									CacheBehavior: "Override",
									CacheDuration: "10:10:09",
									CacheType:     "All",
									TypeName:      "DeliveryRuleCacheExpirationActionParameters",
								},
							},
							{
								Name: "ModifyResponseHeader",
								Parameters: {
									HeaderAction: "Overwrite",
									HeaderName:   "Access-Control-Allow-Origin",
									TypeName:     "DeliveryRuleHeaderActionParameters",
									Value:        "*",
								},
							},
							{
								Name: "ModifyRequestHeader",
								Parameters: {
									HeaderAction: "Overwrite",
									HeaderName:   "Accept-Encoding",
									TypeName:     "DeliveryRuleHeaderActionParameters",
									Value:        "gzip",
								},
							},
						},
						Conditions: pulumi.Array{
							{
								Name: "RemoteAddress",
								Parameters: {
									MatchValues: []string{
										"192.168.1.0/24",
										"10.0.0.0/24",
									},
									NegateCondition: true,
									Operator:        "IPMatch",
									TypeName:        "DeliveryRuleRemoteAddressConditionParameters",
								},
							},
						},
						Name:  pulumi.String("rule1"),
						Order: pulumi.Int(1),
					},
				},
			},
			EndpointName:         pulumi.String("endpoint1"),
			IsCompressionEnabled: pulumi.Bool(true),
			IsHttpAllowed:        pulumi.Bool(true),
			IsHttpsAllowed:       pulumi.Bool(true),
			Location:             pulumi.String("WestUs"),
			OriginGroups: cdn.DeepCreatedOriginGroupArray{
				&cdn.DeepCreatedOriginGroupArgs{
					HealthProbeSettings: &cdn.HealthProbeParametersArgs{
						ProbeIntervalInSeconds: pulumi.Int(120),
						ProbePath:              pulumi.String("/health.aspx"),
						ProbeProtocol:          cdn.ProbeProtocolHttp,
						ProbeRequestType:       cdn.HealthProbeRequestTypeGET,
					},
					Name: pulumi.String("originGroup1"),
					Origins: cdn.ResourceReferenceArray{
						&cdn.ResourceReferenceArgs{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin1"),
						},
						&cdn.ResourceReferenceArgs{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin2"),
						},
					},
					ResponseBasedOriginErrorDetectionSettings: &cdn.ResponseBasedOriginErrorDetectionParametersArgs{
						ResponseBasedDetectedErrorTypes:          cdn.ResponseBasedDetectedErrorTypesTcpErrorsOnly,
						ResponseBasedFailoverThresholdPercentage: pulumi.Int(10),
					},
				},
			},
			OriginHostHeader: pulumi.String("www.bing.com"),
			OriginPath:       pulumi.String("/photos"),
			Origins: cdn.DeepCreatedOriginArray{
				&cdn.DeepCreatedOriginArgs{
					Enabled:          pulumi.Bool(true),
					HostName:         pulumi.String("www.someDomain1.net"),
					HttpPort:         pulumi.Int(80),
					HttpsPort:        pulumi.Int(443),
					Name:             pulumi.String("origin1"),
					OriginHostHeader: pulumi.String("www.someDomain1.net"),
					Priority:         pulumi.Int(1),
					Weight:           pulumi.Int(50),
				},
				&cdn.DeepCreatedOriginArgs{
					Enabled:          pulumi.Bool(true),
					HostName:         pulumi.String("www.someDomain2.net"),
					HttpPort:         pulumi.Int(80),
					HttpsPort:        pulumi.Int(443),
					Name:             pulumi.String("origin2"),
					OriginHostHeader: pulumi.String("www.someDomain2.net"),
					Priority:         pulumi.Int(2),
					Weight:           pulumi.Int(50),
				},
			},
			ProfileName:                pulumi.String("profile1"),
			QueryStringCachingBehavior: cdn.QueryStringCachingBehaviorBypassCaching,
			ResourceGroupName:          pulumi.String("RG"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
