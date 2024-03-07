package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewWebTest(ctx, "webTest", &insights.WebTestArgs{
			Description: pulumi.String("Ping web test alert for mytestwebapp"),
			Enabled:     pulumi.Bool(true),
			Frequency:   pulumi.Int(900),
			Location:    pulumi.String("South Central US"),
			Locations: []insights.WebTestGeolocationArgs{
				{
					Location: pulumi.String("us-fl-mia-edge"),
				},
			},
			Request: insights.WebTestPropertiesResponseRequest{
				Headers: insights.HeaderFieldArray{
					&insights.HeaderFieldArgs{
						HeaderFieldName:  pulumi.String("Content-Language"),
						HeaderFieldValue: pulumi.String("de-DE"),
					},
					&insights.HeaderFieldArgs{
						HeaderFieldName:  pulumi.String("Accept-Language"),
						HeaderFieldValue: pulumi.String("de-DE"),
					},
				},
				HttpVerb:    pulumi.String("POST"),
				RequestBody: pulumi.String("SGVsbG8gd29ybGQ="),
				RequestUrl:  pulumi.String("https://bing.com"),
			},
			ResourceGroupName:  pulumi.String("my-resource-group"),
			RetryEnabled:       pulumi.Bool(true),
			SyntheticMonitorId: pulumi.String("my-webtest-my-component"),
			Timeout:            pulumi.Int(120),
			ValidationRules: &insights.WebTestPropertiesValidationRulesArgs{
				SSLCertRemainingLifetimeCheck: pulumi.Int(100),
				SSLCheck:                      pulumi.Bool(true),
			},
			WebTestKind: insights.WebTestKindStandard,
			WebTestName: pulumi.String("my-webtest-my-component"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
