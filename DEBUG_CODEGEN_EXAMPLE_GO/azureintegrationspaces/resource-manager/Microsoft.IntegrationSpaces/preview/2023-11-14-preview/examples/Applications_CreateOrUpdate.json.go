package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/integrationspaces/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := integrationspaces.NewApplication(ctx, "application", &integrationspaces.ApplicationArgs{
			ApplicationName:   pulumi.String("Application1"),
			Description:       pulumi.String("This is the user provided description of the application."),
			Location:          pulumi.String("CentralUS"),
			ResourceGroupName: pulumi.String("testrg"),
			SpaceName:         pulumi.String("Space1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("Value1"),
			},
			TrackingDataStores: integrationspaces.TrackingDataStoreMap{
				"dataStoreName1": &integrationspaces.TrackingDataStoreArgs{
					DataStoreIngestionUri: pulumi.String("https://ingest-someClusterName.someRegionName.kusto.windows.net"),
					DataStoreResourceId:   pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Kusto/Clusters/cluster1"),
					DataStoreUri:          pulumi.String("https://someClusterName.someRegionName.kusto.windows.net"),
					DatabaseName:          pulumi.String("testDatabase1"),
				},
				"dataStoreName2": &integrationspaces.TrackingDataStoreArgs{
					DataStoreIngestionUri: pulumi.String("https://ingest-someClusterName.someRegionName.kusto.windows.net"),
					DataStoreResourceId:   pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Kusto/Clusters/cluster1"),
					DataStoreUri:          pulumi.String("https://someClusterName.someRegionName.kusto.windows.net"),
					DatabaseName:          pulumi.String("testDatabase1"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
