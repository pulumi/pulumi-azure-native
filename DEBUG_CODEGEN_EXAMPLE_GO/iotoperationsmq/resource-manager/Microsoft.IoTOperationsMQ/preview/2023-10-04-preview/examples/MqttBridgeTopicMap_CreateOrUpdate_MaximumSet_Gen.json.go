package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsmq/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsmq.NewMqttBridgeTopicMap(ctx, "mqttBridgeTopicMap", &iotoperationsmq.MqttBridgeTopicMapArgs{
			ExtendedLocation: &iotoperationsmq.ExtendedLocationPropertyArgs{
				Name: pulumi.String("an"),
				Type: pulumi.String("CustomLocation"),
			},
			Location:                pulumi.String("icfdftifk"),
			MqName:                  pulumi.String("52A1-D1-t--Q7O9-C-2S"),
			MqttBridgeConnectorName: pulumi.String("5sKfh6461-KDI8h-5"),
			MqttBridgeConnectorRef:  pulumi.String("aemmhvfdzmdtxwgimpaqephgo"),
			ResourceGroupName:       pulumi.String("rgiotoperationsmq"),
			Routes: iotoperationsmq.MqttBridgeRoutesArray{
				&iotoperationsmq.MqttBridgeRoutesArgs{
					Direction: pulumi.String("remote-to-local"),
					Name:      pulumi.String("u"),
					Qos:       pulumi.Int(2),
					SharedSubscription: &iotoperationsmq.MqttBridgeRouteSharedSubscriptionArgs{
						GroupMinimumShareNumber: pulumi.Int(129),
						GroupName:               pulumi.String("exnfgkdccpuvzqhxrg"),
					},
					Source: pulumi.String("xwnfgkkfezlgh"),
					Target: pulumi.String("mgwem"),
				},
			},
			Tags:         nil,
			TopicMapName: pulumi.String("5--CR4S47--UaxB4-"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
