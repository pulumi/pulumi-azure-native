package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devices.NewIotHubResource(ctx, "iotHubResource", &devices.IotHubResourceArgs{
			Location: pulumi.String("centraluseuap"),
			Properties: &devices.IotHubPropertiesArgs{
				CloudToDevice: &devices.CloudToDevicePropertiesArgs{
					DefaultTtlAsIso8601: pulumi.String("PT1H"),
					Feedback: &devices.FeedbackPropertiesArgs{
						LockDurationAsIso8601: pulumi.String("PT1M"),
						MaxDeliveryCount:      pulumi.Int(10),
						TtlAsIso8601:          pulumi.String("PT1H"),
					},
					MaxDeliveryCount: pulumi.Int(10),
				},
				EnableDataResidency:           pulumi.Bool(true),
				EnableFileUploadNotifications: pulumi.Bool(false),
				EventHubEndpoints: devices.EventHubPropertiesMap{
					"events": &devices.EventHubPropertiesArgs{
						PartitionCount:      pulumi.Int(2),
						RetentionTimeInDays: pulumi.Float64(1),
					},
				},
				Features:      pulumi.String("None"),
				IpFilterRules: devices.IpFilterRuleArray{},
				MessagingEndpoints: devices.MessagingEndpointPropertiesMap{
					"fileNotifications": &devices.MessagingEndpointPropertiesArgs{
						LockDurationAsIso8601: pulumi.String("PT1M"),
						MaxDeliveryCount:      pulumi.Int(10),
						TtlAsIso8601:          pulumi.String("PT1H"),
					},
				},
				MinTlsVersion: pulumi.String("1.2"),
				NetworkRuleSets: &devices.NetworkRuleSetPropertiesArgs{
					ApplyToBuiltInEventHubEndpoint: pulumi.Bool(true),
					DefaultAction:                  pulumi.String("Deny"),
					IpRules: devices.NetworkRuleSetIpRuleArray{
						&devices.NetworkRuleSetIpRuleArgs{
							Action:     pulumi.String("Allow"),
							FilterName: pulumi.String("rule1"),
							IpMask:     pulumi.String("131.117.159.53"),
						},
						&devices.NetworkRuleSetIpRuleArgs{
							Action:     pulumi.String("Allow"),
							FilterName: pulumi.String("rule2"),
							IpMask:     pulumi.String("157.55.59.128/25"),
						},
					},
				},
				RootCertificate: &devices.RootCertificatePropertiesArgs{
					EnableRootCertificateV2: pulumi.Bool(true),
				},
				Routing: &devices.RoutingPropertiesArgs{
					Endpoints: &devices.RoutingEndpointsArgs{
						EventHubs:         devices.RoutingEventHubPropertiesArray{},
						ServiceBusQueues:  devices.RoutingServiceBusQueueEndpointPropertiesArray{},
						ServiceBusTopics:  devices.RoutingServiceBusTopicEndpointPropertiesArray{},
						StorageContainers: devices.RoutingStorageContainerPropertiesArray{},
					},
					FallbackRoute: &devices.FallbackRoutePropertiesArgs{
						Condition: pulumi.String("true"),
						EndpointNames: pulumi.StringArray{
							pulumi.String("events"),
						},
						IsEnabled: pulumi.Bool(true),
						Name:      pulumi.String("$fallback"),
						Source:    pulumi.String("DeviceMessages"),
					},
					Routes: devices.RoutePropertiesArray{},
				},
				StorageEndpoints: devices.StorageEndpointPropertiesMap{
					"$default": &devices.StorageEndpointPropertiesArgs{
						ConnectionString: pulumi.String(""),
						ContainerName:    pulumi.String(""),
						SasTtlAsIso8601:  pulumi.String("PT1H"),
					},
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ResourceName:      pulumi.String("testHub"),
			Sku: &devices.IotHubSkuInfoArgs{
				Capacity: pulumi.Float64(1),
				Name:     pulumi.String("S1"),
			},
			Tags: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
