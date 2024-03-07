package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewWorkbook(ctx, "workbook", &insights.WorkbookArgs{
			Category:          pulumi.String("workbook"),
			Description:       pulumi.String("Sample workbook"),
			DisplayName:       pulumi.String("Sample workbook"),
			Kind:              pulumi.String("shared"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("my-resource-group"),
			ResourceName:      pulumi.String("deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
			SerializedData:    pulumi.String("{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}"),
			SourceId:          pulumi.String("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group"),
			Tags: pulumi.StringMap{
				"TagSample01": pulumi.String("sample01"),
				"TagSample02": pulumi.String("sample02"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
