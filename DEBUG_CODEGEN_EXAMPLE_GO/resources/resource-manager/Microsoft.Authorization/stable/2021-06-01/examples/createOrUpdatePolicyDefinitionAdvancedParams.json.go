package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewPolicyDefinition(ctx, "policyDefinition", &authorization.PolicyDefinitionArgs{
			Description: pulumi.String("Audit enabling of logs and retain them up to a year. This enables recreation of activity trails for investigation purposes when a security incident occurs or your network is compromised"),
			DisplayName: pulumi.String("Event Hubs should have diagnostic logging enabled"),
			Metadata: pulumi.Any{
				Category: "Event Hub",
			},
			Mode: pulumi.String("Indexed"),
			Parameters: authorization.ParameterDefinitionsValueMap{
				"requiredRetentionDays": &authorization.ParameterDefinitionsValueArgs{
					AllowedValues: pulumi.Array{
						pulumi.Any(0),
						pulumi.Any(30),
						pulumi.Any(90),
						pulumi.Any(180),
						pulumi.Any(365),
					},
					DefaultValue: pulumi.Any(365),
					Metadata: &authorization.ParameterDefinitionsValueMetadataArgs{
						Description: pulumi.String("The required diagnostic logs retention in days"),
						DisplayName: pulumi.String("Required retention (days)"),
					},
					Type: pulumi.String("Integer"),
				},
			},
			PolicyDefinitionName: pulumi.String("EventHubDiagnosticLogs"),
			PolicyRule: pulumi.Any{
				If: map[string]interface{}{
					"equals": "Microsoft.EventHub/namespaces",
					"field":  "type",
				},
				Then: map[string]interface{}{
					"details": map[string]interface{}{
						"existenceCondition": map[string]interface{}{
							"allOf": []map[string]interface{}{
								map[string]interface{}{
									"equals": "true",
									"field":  "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.enabled",
								},
								map[string]interface{}{
									"equals": "[parameters('requiredRetentionDays')]",
									"field":  "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.days",
								},
							},
						},
						"type": "Microsoft.Insights/diagnosticSettings",
					},
					"effect": "AuditIfNotExists",
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
