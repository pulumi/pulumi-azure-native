package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewSystemTopic(ctx, "systemTopic", &eventgrid.SystemTopicArgs{
			Location:          pulumi.String("westus2"),
			ResourceGroupName: pulumi.String("examplerg"),
			Source:            pulumi.String("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/azureeventgridrunnerrgcentraluseuap/providers/microsoft.storage/storageaccounts/pubstgrunnerb71cd29e"),
			SystemTopicName:   pulumi.String("exampleSystemTopic1"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
			TopicType: pulumi.String("microsoft.storage.storageaccounts"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
