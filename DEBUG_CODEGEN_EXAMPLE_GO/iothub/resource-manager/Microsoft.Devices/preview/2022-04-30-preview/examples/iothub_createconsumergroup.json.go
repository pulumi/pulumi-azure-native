package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devices.NewIotHubResourceEventHubConsumerGroup(ctx, "iotHubResourceEventHubConsumerGroup", &devices.IotHubResourceEventHubConsumerGroupArgs{
			EventHubEndpointName: pulumi.String("events"),
			Name:                 pulumi.String("test"),
			Properties: &devices.EventHubConsumerGroupNameArgs{
				Name: pulumi.String("test"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ResourceName:      pulumi.String("testHub"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
