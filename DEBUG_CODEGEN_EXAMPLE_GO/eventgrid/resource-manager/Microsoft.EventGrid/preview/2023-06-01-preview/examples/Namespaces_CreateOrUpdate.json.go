package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewNamespace(ctx, "namespace", &eventgrid.NamespaceArgs{
			Location:          pulumi.String("westus"),
			NamespaceName:     pulumi.String("exampleNamespaceName1"),
			ResourceGroupName: pulumi.String("examplerg"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value11"),
				"tag2": pulumi.String("value22"),
			},
			TopicSpacesConfiguration: &eventgrid.TopicSpacesConfigurationArgs{
				RouteTopicResourceId: pulumi.String("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampleTopic1"),
				State:                pulumi.String("Enabled"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
