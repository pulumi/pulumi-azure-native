package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsmq/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsmq.NewKafkaConnectorTopicMap(ctx, "kafkaConnectorTopicMap", &iotoperationsmq.KafkaConnectorTopicMapArgs{
			Batching: &iotoperationsmq.KafkaTopicMapBatchingArgs{
				Enabled:     pulumi.Bool(true),
				LatencyMs:   pulumi.Int(9110),
				MaxBytes:    pulumi.Float64(732052221),
				MaxMessages: pulumi.Float64(373078076),
			},
			Compression:        pulumi.String("none"),
			CopyMqttProperties: pulumi.String("efpqgkycuawnzyubdyt"),
			ExtendedLocation: &iotoperationsmq.ExtendedLocationPropertyArgs{
				Name: pulumi.String("an"),
				Type: pulumi.String("CustomLocation"),
			},
			KafkaConnectorName:   pulumi.String("216VN"),
			KafkaConnectorRef:    pulumi.String("icivjwerdspx"),
			Location:             pulumi.String("pavphpzfsgdudpyvufyebqh"),
			MqName:               pulumi.String("-1-eD-7-J"),
			PartitionKeyProperty: pulumi.String("c"),
			PartitionStrategy:    pulumi.String("default"),
			ResourceGroupName:    pulumi.String("rgiotoperationsmq"),
			Routes: []iotoperationsmq.KafkaRoutesArgs{
				{
					KafkaToMqtt: {
						ConsumerGroupId: pulumi.String("usork"),
						KafkaTopic:      pulumi.String("ggwhwbsr"),
						MqttTopic:       pulumi.String("jwvmmhfqqkkmqrpslbdfmpbdetfu"),
						Name:            pulumi.String("lrnvudysggscnqvmnlkrk"),
						Qos:             pulumi.Int(1),
					},
					MqttToKafka: {
						KafkaAcks:  pulumi.String("zero"),
						KafkaTopic: pulumi.String("tellycttwulueqcpqf"),
						MqttTopic:  pulumi.String("raipkrcwvdnnflywhgjwnquarf"),
						Name:       pulumi.String("qpshqcaxvxnyjzimvchngupzezdei"),
						Qos:        pulumi.Int(1),
						SharedSubscription: {
							GroupMinimumShareNumber: pulumi.Int(216),
							GroupName:               pulumi.String("nwdyccsditzhchuksmi"),
						},
					},
				},
			},
			Tags:         nil,
			TopicMapName: pulumi.String("q582ViEY-b7wF1OO2A"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
