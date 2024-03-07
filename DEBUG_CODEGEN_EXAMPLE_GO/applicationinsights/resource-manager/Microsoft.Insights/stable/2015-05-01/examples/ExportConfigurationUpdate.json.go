package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewExportConfiguration(ctx, "exportConfiguration", &insights.ExportConfigurationArgs{
			DestinationAccountId:             pulumi.String("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.ClassicStorage/storageAccounts/mystorageblob"),
			DestinationAddress:               pulumi.String("https://mystorageblob.blob.core.windows.net/fchentest?sv=2015-04-05&sr=c&sig=token"),
			DestinationStorageLocationId:     pulumi.String("eastus"),
			DestinationStorageSubscriptionId: pulumi.String("subid"),
			DestinationType:                  pulumi.String("Blob"),
			ExportId:                         pulumi.String("uGOoki0jQsyEs3IdQ83Q4QsNr4="),
			IsEnabled:                        pulumi.String("true"),
			NotificationQueueEnabled:         pulumi.String("false"),
			NotificationQueueUri:             pulumi.String(""),
			RecordTypes:                      pulumi.String("Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability"),
			ResourceGroupName:                pulumi.String("my-resource-group"),
			ResourceName:                     pulumi.String("my-component"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
