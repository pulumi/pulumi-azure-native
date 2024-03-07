package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewAnomalySecurityMLAnalyticsSettings(ctx, "anomalySecurityMLAnalyticsSettings", &securityinsights.AnomalySecurityMLAnalyticsSettingsArgs{
			AnomalySettingsVersion: pulumi.Int(0),
			AnomalyVersion:         pulumi.String("1.0.5"),
			CustomizableObservations: pulumi.Any{
				MultiSelectObservations:       nil,
				PrioritizeExcludeObservations: nil,
				SingleSelectObservations: []map[string]interface{}{
					map[string]interface{}{
						"description":    "Select device vendor of network connection logs from CommonSecurityLog",
						"name":           "Device vendor",
						"rerun":          "RerunAlways",
						"sequenceNumber": 1,
						"supportedValues": []string{
							"Palo Alto Networks",
							"Fortinet",
							"Check Point",
						},
						"supportedValuesKql": nil,
						"value": []string{
							"Palo Alto Networks",
						},
						"valuesKql": nil,
					},
				},
				SingleValueObservations: nil,
				ThresholdObservations: []interface{}{
					map[string]interface{}{
						"description":    "Suppress anomalies when daily data transfered (in MB) per hour is less than the chosen value",
						"maximum":        "100",
						"minimum":        "1",
						"name":           "Daily data transfer threshold in MB",
						"rerun":          "RerunAlways",
						"sequenceNumber": 1,
						"value":          "25",
					},
					map[string]interface{}{
						"description":    "Triggers anomalies when number of standard deviations is greater than the chosen value",
						"maximum":        "10",
						"minimum":        "2",
						"name":           "Number of standard deviations",
						"rerun":          "RerunAlways",
						"sequenceNumber": 2,
						"value":          "3",
					},
				},
			},
			Description:       pulumi.String("When account logs from a source region that has rarely been logged in from during the last 14 days, an anomaly is triggered."),
			DisplayName:       pulumi.String("Login from unusual region"),
			Enabled:           pulumi.Bool(true),
			Frequency:         pulumi.String("PT1H"),
			IsDefaultSettings: pulumi.Bool(true),
			Kind:              pulumi.String("Anomaly"),
			RequiredDataConnectors: securityinsights.SecurityMLAnalyticsSettingsDataSourceArray{
				&securityinsights.SecurityMLAnalyticsSettingsDataSourceArgs{
					ConnectorId: pulumi.String("AWS"),
					DataTypes: pulumi.StringArray{
						pulumi.String("AWSCloudTrail"),
					},
				},
			},
			ResourceGroupName:    pulumi.String("myRg"),
			SettingsDefinitionId: pulumi.String("f209187f-1d17-4431-94af-c141bf5f23db"),
			SettingsResourceName: pulumi.String("f209187f-1d17-4431-94af-c141bf5f23db"),
			SettingsStatus:       pulumi.String("Production"),
			Tactics: pulumi.StringArray{
				pulumi.String("Exfiltration"),
				pulumi.String("CommandAndControl"),
			},
			Techniques: pulumi.StringArray{
				pulumi.String("T1037"),
				pulumi.String("T1021"),
			},
			WorkspaceName: pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
