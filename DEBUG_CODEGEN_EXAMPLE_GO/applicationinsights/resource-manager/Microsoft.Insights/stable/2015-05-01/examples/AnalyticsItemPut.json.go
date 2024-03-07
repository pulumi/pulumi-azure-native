package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewAnalyticsItem(ctx, "analyticsItem", &insights.AnalyticsItemArgs{
			Content: pulumi.String(`let newExceptionsTimeRange = 1d;
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
			Name:              pulumi.String("Exceptions - New in the last 24 hours"),
			ResourceGroupName: pulumi.String("my-resource-group"),
			ResourceName:      pulumi.String("my-component"),
			Scope:             pulumi.String("shared"),
			ScopePath:         pulumi.String("analyticsItems"),
			Type:              pulumi.String("query"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
