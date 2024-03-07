package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewScheduledAlertRule(ctx, "scheduledAlertRule", &securityinsights.ScheduledAlertRuleArgs{
			AlertDetailsOverride: &securityinsights.AlertDetailsOverrideArgs{
				AlertDescriptionFormat: pulumi.String("Suspicious activity was made by {{ComputerIP}}"),
				AlertDisplayNameFormat: pulumi.String("Alert from {{Computer}}"),
				AlertDynamicProperties: securityinsights.AlertPropertyMappingArray{
					&securityinsights.AlertPropertyMappingArgs{
						AlertProperty: pulumi.String("ProductComponentName"),
						Value:         pulumi.String("ProductComponentNameCustomColumn"),
					},
					&securityinsights.AlertPropertyMappingArgs{
						AlertProperty: pulumi.String("ProductName"),
						Value:         pulumi.String("ProductNameCustomColumn"),
					},
					&securityinsights.AlertPropertyMappingArgs{
						AlertProperty: pulumi.String("AlertLink"),
						Value:         pulumi.String("Link"),
					},
				},
			},
			CustomDetails: pulumi.StringMap{
				"OperatingSystemName": pulumi.String("OSName"),
				"OperatingSystemType": pulumi.String("OSType"),
			},
			Description: pulumi.String("An example for a scheduled rule"),
			DisplayName: pulumi.String("My scheduled rule"),
			Enabled:     pulumi.Bool(true),
			EntityMappings: securityinsights.EntityMappingArray{
				&securityinsights.EntityMappingArgs{
					EntityType: pulumi.String("Host"),
					FieldMappings: securityinsights.FieldMappingArray{
						&securityinsights.FieldMappingArgs{
							ColumnName: pulumi.String("Computer"),
							Identifier: pulumi.String("FullName"),
						},
					},
				},
				&securityinsights.EntityMappingArgs{
					EntityType: pulumi.String("IP"),
					FieldMappings: securityinsights.FieldMappingArray{
						&securityinsights.FieldMappingArgs{
							ColumnName: pulumi.String("ComputerIP"),
							Identifier: pulumi.String("Address"),
						},
					},
				},
			},
			EventGroupingSettings: &securityinsights.EventGroupingSettingsArgs{
				AggregationKind: pulumi.String("AlertPerResult"),
			},
			IncidentConfiguration: &securityinsights.IncidentConfigurationArgs{
				CreateIncident: pulumi.Bool(true),
				GroupingConfiguration: &securityinsights.GroupingConfigurationArgs{
					Enabled: pulumi.Bool(true),
					GroupByAlertDetails: pulumi.StringArray{
						pulumi.String("DisplayName"),
					},
					GroupByCustomDetails: pulumi.StringArray{
						pulumi.String("OperatingSystemType"),
						pulumi.String("OperatingSystemName"),
					},
					GroupByEntities: pulumi.StringArray{
						pulumi.String("Host"),
					},
					LookbackDuration:     pulumi.String("PT5H"),
					MatchingMethod:       pulumi.String("Selected"),
					ReopenClosedIncident: pulumi.Bool(false),
				},
			},
			Kind:                pulumi.String("Scheduled"),
			Query:               pulumi.String("Heartbeat"),
			QueryFrequency:      pulumi.String("PT1H"),
			QueryPeriod:         pulumi.String("P2DT1H30M"),
			ResourceGroupName:   pulumi.String("myRg"),
			RuleId:              pulumi.String("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
			Severity:            pulumi.String("High"),
			SuppressionDuration: pulumi.String("PT1H"),
			SuppressionEnabled:  pulumi.Bool(false),
			Tactics: pulumi.StringArray{
				pulumi.String("Persistence"),
				pulumi.String("LateralMovement"),
			},
			TriggerOperator:  securityinsights.TriggerOperatorGreaterThan,
			TriggerThreshold: pulumi.Int(0),
			WorkspaceName:    pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
