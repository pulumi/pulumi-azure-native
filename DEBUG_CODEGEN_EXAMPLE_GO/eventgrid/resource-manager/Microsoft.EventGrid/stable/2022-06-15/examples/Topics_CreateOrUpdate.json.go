package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewTopic(ctx, "topic", &eventgrid.TopicArgs{
			InboundIpRules: eventgrid.InboundIpRuleArray{
				&eventgrid.InboundIpRuleArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("12.18.30.15"),
				},
				&eventgrid.InboundIpRuleArgs{
					Action: pulumi.String("Allow"),
					IpMask: pulumi.String("12.18.176.1"),
				},
			},
			Location:            pulumi.String("westus2"),
			PublicNetworkAccess: pulumi.String("Enabled"),
			ResourceGroupName:   pulumi.String("examplerg"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
			TopicName: pulumi.String("exampletopic1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
