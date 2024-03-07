package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewTopicSpace(ctx, "topicSpace", &eventgrid.TopicSpaceArgs{
			NamespaceName:     pulumi.String("exampleNamespaceName1"),
			ResourceGroupName: pulumi.String("examplerg"),
			TopicSpaceName:    pulumi.String("exampleTopicSpaceName1"),
			TopicTemplates: pulumi.StringArray{
				pulumi.String("filter1"),
				pulumi.String("filter2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
