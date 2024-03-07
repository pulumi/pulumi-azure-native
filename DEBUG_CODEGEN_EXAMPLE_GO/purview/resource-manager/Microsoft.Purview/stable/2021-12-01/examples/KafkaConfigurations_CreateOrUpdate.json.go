package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/purview/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := purview.NewKafkaConfiguration(ctx, "kafkaConfiguration", &purview.KafkaConfigurationArgs{
			AccountName:   pulumi.String("account1"),
			ConsumerGroup: pulumi.String("consumerGroup"),
			Credentials: &purview.CredentialsArgs{
				IdentityId: pulumi.String("/subscriptions/47e8596d-ee73-4eb2-b6b4-cc13c2b87ssd/resourceGroups/testRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testId"),
				Type:       pulumi.String("UserAssigned"),
			},
			EventHubPartitionId:    pulumi.String("partitionId"),
			EventHubResourceId:     pulumi.String("/subscriptions/225be6fe-ec1c-4d51-a368-f69348d2e6c5/resourceGroups/testRG/providers/Microsoft.EventHub/namespaces/eventHubNameSpaceName"),
			EventHubType:           pulumi.String("Notification"),
			EventStreamingState:    pulumi.String("Enabled"),
			EventStreamingType:     pulumi.String("Azure"),
			KafkaConfigurationName: pulumi.String("kafkaConfigName"),
			ResourceGroupName:      pulumi.String("rgpurview"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
