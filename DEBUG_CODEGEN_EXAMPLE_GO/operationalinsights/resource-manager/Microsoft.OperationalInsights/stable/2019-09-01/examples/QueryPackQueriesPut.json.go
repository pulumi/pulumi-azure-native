package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewQuery(ctx, "query", &operationalinsights.QueryArgs{
			Body: pulumi.String(`let newExceptionsTimeRange = 1d;
let timeRangeToCheckBefore = 7d;
exceptions
| where timestamp < ago(timeRangeToCheckBefore)
| summarize count() by problemId
| join kind= rightanti (
exceptions
| where timestamp >= ago(newExceptionsTimeRange)
| extend stack = tostring(details[0].rawStack)
| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  
) on problemId 
| order by  count_ desc
`),
			Description:   pulumi.String("my description"),
			DisplayName:   pulumi.String("Exceptions - New in the last 24 hours"),
			Id:            pulumi.String("a449f8af-8e64-4b3a-9b16-5a7165ff98c4"),
			QueryPackName: pulumi.String("my-querypack"),
			Related: &operationalinsights.LogAnalyticsQueryPackQueryPropertiesRelatedArgs{
				Categories: pulumi.StringArray{
					pulumi.String("analytics"),
				},
			},
			ResourceGroupName: pulumi.String("my-resource-group"),
			Tags: pulumi.StringArrayMap{
				"my-label": pulumi.StringArray{
					pulumi.String("label1"),
				},
				"my-other-label": pulumi.StringArray{
					pulumi.String("label2"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
