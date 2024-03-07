package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewDomainTopic(ctx, "domainTopic", &eventgrid.DomainTopicArgs{
			DomainName:        pulumi.String("exampledomain1"),
			DomainTopicName:   pulumi.String("exampledomaintopic1"),
			ResourceGroupName: pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
