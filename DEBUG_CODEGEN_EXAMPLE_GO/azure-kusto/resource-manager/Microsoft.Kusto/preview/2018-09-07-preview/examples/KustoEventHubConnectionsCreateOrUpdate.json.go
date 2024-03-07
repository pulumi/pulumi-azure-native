package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kusto/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kusto.NewEventHubConnection(ctx, "eventHubConnection", &kusto.EventHubConnectionArgs{
			ClusterName:            pulumi.String("KustoClusterRPTest4"),
			ConsumerGroup:          pulumi.String("testConsumerGroup1"),
			DatabaseName:           pulumi.String("KustoDatabase8"),
			EventHubConnectionName: pulumi.String("kustoeventhubconnection1"),
			EventHubResourceId:     pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
			Location:               pulumi.String("westus"),
			ResourceGroupName:      pulumi.String("kustorptest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
