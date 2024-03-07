package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/digitaltwins/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := digitaltwins.NewTimeSeriesDatabaseConnection(ctx, "timeSeriesDatabaseConnection", &digitaltwins.TimeSeriesDatabaseConnectionArgs{
			Properties: &digitaltwins.AzureDataExplorerConnectionPropertiesArgs{
				AdxDatabaseName:                         pulumi.String("myDatabase"),
				AdxEndpointUri:                          pulumi.String("https://mycluster.kusto.windows.net"),
				AdxRelationshipLifecycleEventsTableName: pulumi.String("myRelationshipLifecycleEventsTable"),
				AdxResourceId:                           pulumi.String("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.Kusto/clusters/mycluster"),
				AdxTableName:                            pulumi.String("myPropertyUpdatesTable"),
				AdxTwinLifecycleEventsTableName:         pulumi.String("myTwinLifecycleEventsTable"),
				ConnectionType:                          pulumi.String("AzureDataExplorer"),
				EventHubEndpointUri:                     pulumi.String("sb://myeh.servicebus.windows.net/"),
				EventHubEntityPath:                      pulumi.String("myeh"),
				EventHubNamespaceResourceId:             pulumi.String("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.EventHub/namespaces/myeh"),
				RecordPropertyAndItemRemovals:           pulumi.String("true"),
			},
			ResourceGroupName:                pulumi.String("resRg"),
			ResourceName:                     pulumi.String("myDigitalTwinsService"),
			TimeSeriesDatabaseConnectionName: pulumi.String("myConnection"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
