package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewContentTemplate(ctx, "contentTemplate", &securityinsights.ContentTemplateArgs{
			Author: &securityinsights.MetadataAuthorArgs{
				Email: pulumi.String("support@microsoft.com"),
				Name:  pulumi.String("Microsoft"),
			},
			ContentId:   pulumi.String("8365ebfe-a381-45b7-ad08-7d818070e11f"),
			ContentKind: pulumi.String("AnalyticsRule"),
			DisplayName: pulumi.String("API Protection workbook template"),
			MainTemplate: pulumi.Any{
				Schema:         "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
				ContentVersion: "1.0.1",
				Resources: []interface{}{
					map[string]interface{}{
						"apiVersion": "2022-04-01-preview",
						"kind":       "Scheduled",
						"location":   "[parameters('workspace-location')]",
						"name":       "8365ebfe-a381-45b7-ad08-7d818070e11f",
						"properties": map[string]interface{}{
							"description":         "Creates an incident when a large number of Critical/High severity CrowdStrike Falcon sensor detections is triggered by a single user",
							"displayName":         "Critical or High Severity Detections by User",
							"enabled":             false,
							"query":               "...",
							"queryFrequency":      "PT1H",
							"queryPeriod":         "PT1H",
							"severity":            "High",
							"status":              "Available",
							"suppressionDuration": "PT1H",
							"suppressionEnabled":  false,
							"triggerOperator":     "GreaterThan",
							"triggerThreshold":    0,
						},
						"type": "Microsoft.SecurityInsights/AlertRuleTemplates",
					},
					map[string]interface{}{
						"apiVersion": "2022-01-01-preview",
						"name":       "[concat(parameters('workspace'),'/Microsoft.SecurityInsights/',concat('AnalyticsRule-', last(split([resourceId('Microsoft.SecurityInsights/AlertRuleTemplates', 8365ebfe-a381-45b7-ad08-7d818070e11f)],'/'))))]",
						"properties": map[string]interface{}{
							"author": map[string]interface{}{
								"email": "support@microsoft.com",
								"name":  "Microsoft",
							},
							"contentId":   "4465ebde-b381-45f7-ad08-7d818070a11c",
							"description": "CrowdStrike Falcon Endpoint Protection Analytics Rule 1",
							"kind":        "AnalyticsRule",
							"parentId":    "[resourceId('Microsoft.SecurityInsights/AlertRuleTemplates', 8365ebfe-a381-45b7-ad08-7d818070e11f)]",
							"source": map[string]interface{}{
								"kind":     "Solution",
								"name":     "str",
								"sourceId": "str.azure-sentinel-solution-str",
							},
							"support": map[string]interface{}{
								"email": "support@microsoft.com",
								"link":  "https://support.microsoft.com/",
								"name":  "Microsoft Corporation",
								"tier":  "Microsoft",
							},
							"version": "1.0.0",
						},
						"type": "Microsoft.OperationalInsights/workspaces/providers/metadata",
					},
				},
			},
			PackageId:         pulumi.String("str.azure-sentinel-solution-str"),
			PackageKind:       pulumi.String("Solution"),
			PackageName:       pulumi.String("str"),
			ResourceGroupName: pulumi.String("myRg"),
			Source: &securityinsights.MetadataSourceArgs{
				Kind:     pulumi.String("Solution"),
				Name:     pulumi.String("str"),
				SourceId: pulumi.String("str.azure-sentinel-solution-str"),
			},
			Support: &securityinsights.MetadataSupportArgs{
				Email: pulumi.String("support@microsoft.com"),
				Link:  pulumi.String("https://support.microsoft.com/"),
				Name:  pulumi.String("Microsoft Corporation"),
				Tier:  pulumi.String("Microsoft"),
			},
			TemplateId:    pulumi.String("str.azure-sentinel-solution-str"),
			Version:       pulumi.String("1.0.1"),
			WorkspaceName: pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
