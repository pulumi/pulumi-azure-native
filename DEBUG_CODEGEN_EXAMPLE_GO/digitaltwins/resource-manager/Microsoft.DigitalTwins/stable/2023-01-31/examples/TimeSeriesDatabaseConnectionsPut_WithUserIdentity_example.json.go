package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/digitaltwins/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := digitaltwins.NewTimeSeriesDatabaseConnection(ctx, "timeSeriesDatabaseConnection", &digitaltwins.TimeSeriesDatabaseConnectionArgs{
			Properties: &digitaltwins.AzureDataExplorerConnectionPropertiesArgs{
				AdxDatabaseName:             pulumi.String("myDatabase"),
				AdxEndpointUri:              pulumi.String("https://mycluster.kusto.windows.net"),
				AdxResourceId:               pulumi.String("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.Kusto/clusters/mycluster"),
				AdxTableName:                pulumi.String("myTable"),
				ConnectionType:              pulumi.String("AzureDataExplorer"),
				EventHubEndpointUri:         pulumi.String("sb://myeh.servicebus.windows.net/"),
				EventHubEntityPath:          pulumi.String("myeh"),
				EventHubNamespaceResourceId: pulumi.String("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.EventHub/namespaces/myeh"),
				Identity: &digitaltwins.ManagedIdentityReferenceArgs{
					Type:                 pulumi.String("UserAssigned"),
					UserAssignedIdentity: pulumi.String("/subscriptions/50016170-c839-41ba-a724-51e9df440b9e/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity"),
				},
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
