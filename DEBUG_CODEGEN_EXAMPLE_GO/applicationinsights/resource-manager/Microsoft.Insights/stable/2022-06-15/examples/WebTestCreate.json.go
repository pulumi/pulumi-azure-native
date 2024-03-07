package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewWebTest(ctx, "webTest", &insights.WebTestArgs{
			Configuration: &insights.WebTestPropertiesConfigurationArgs{
				WebTest: pulumi.String("<WebTest Name=\"my-webtest\" Id=\"678ddf96-1ab8-44c8-9274-123456789abc\" Enabled=\"True\" CssProjectStructure=\"\" CssIteration=\"\" Timeout=\"120\" WorkItemIds=\"\" xmlns=\"http://microsoft.com/schemas/VisualStudio/TeamTest/2010\" Description=\"\" CredentialUserName=\"\" CredentialPassword=\"\" PreAuthenticate=\"True\" Proxy=\"default\" StopOnError=\"False\" RecordedResultFile=\"\" ResultsLocale=\"\" ><Items><Request Method=\"GET\" Guid=\"a4162485-9114-fcfc-e086-123456789abc\" Version=\"1.1\" Url=\"http://my-component.azurewebsites.net\" ThinkTime=\"0\" Timeout=\"120\" ParseDependentRequests=\"True\" FollowRedirects=\"True\" RecordResult=\"True\" Cache=\"False\" ResponseTimeGoal=\"0\" Encoding=\"utf-8\" ExpectedHttpStatusCode=\"200\" ExpectedResponseUrl=\"\" ReportingName=\"\" IgnoreHttpStatusCode=\"False\" /></Items></WebTest>"),
			},
			Description: pulumi.String("Ping web test alert for mytestwebapp"),
			Enabled:     pulumi.Bool(true),
			Frequency:   pulumi.Int(900),
			Kind:        insights.WebTestKindPing,
			Location:    pulumi.String("South Central US"),
			Locations: insights.WebTestGeolocationArray{
				&insights.WebTestGeolocationArgs{
					Location: pulumi.String("us-fl-mia-edge"),
				},
			},
			ResourceGroupName:  pulumi.String("my-resource-group"),
			RetryEnabled:       pulumi.Bool(true),
			SyntheticMonitorId: pulumi.String("my-webtest-my-component"),
			Timeout:            pulumi.Int(120),
			WebTestKind:        insights.WebTestKindPing,
			WebTestName:        pulumi.String("my-webtest-my-component"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
