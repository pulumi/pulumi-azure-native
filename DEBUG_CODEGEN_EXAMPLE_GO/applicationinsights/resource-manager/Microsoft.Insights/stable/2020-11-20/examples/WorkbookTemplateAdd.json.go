package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewWorkbookTemplate(ctx, "workbookTemplate", &insights.WorkbookTemplateArgs{
			Author: pulumi.String("Contoso"),
			Galleries: insights.WorkbookTemplateGalleryArray{
				&insights.WorkbookTemplateGalleryArgs{
					Category:     pulumi.String("Failures"),
					Name:         pulumi.String("Simple Template"),
					Order:        pulumi.Int(100),
					ResourceType: pulumi.String("microsoft.insights/components"),
					Type:         pulumi.String("tsg"),
				},
			},
			Location:          pulumi.String("west us"),
			Priority:          pulumi.Int(1),
			ResourceGroupName: pulumi.String("my-resource-group"),
			ResourceName:      pulumi.String("testtemplate2"),
			TemplateData: pulumi.Any{
				Schema: "https://github.com/Microsoft/Application-Insights-Workbooks/blob/master/schema/workbook.json",
				Items: []interface{}{
					map[string]interface{}{
						"content": map[string]interface{}{
							"json": "## New workbook\n---\n\nWelcome to your new workbook.  This area will display text formatted as markdown.\n\n\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.",
						},
						"name": "text - 2",
						"type": 1,
					},
					map[string]interface{}{
						"content": map[string]interface{}{
							"exportToExcelOptions": "visible",
							"query":                "union withsource=TableName *\n| summarize Count=count() by TableName\n| render barchart",
							"queryType":            0,
							"resourceType":         "microsoft.operationalinsights/workspaces",
							"size":                 1,
							"version":              "KqlItem/1.0",
						},
						"name": "query - 2",
						"type": 3,
					},
				},
				StyleSettings: nil,
				Version:       "Notebook/1.0",
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
