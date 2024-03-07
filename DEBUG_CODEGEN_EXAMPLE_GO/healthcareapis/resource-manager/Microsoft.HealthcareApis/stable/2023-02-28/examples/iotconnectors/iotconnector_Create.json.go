package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/healthcareapis/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := healthcareapis.NewIotConnector(ctx, "iotConnector", &healthcareapis.IotConnectorArgs{
			DeviceMapping: &healthcareapis.IotMappingPropertiesArgs{
				Content: pulumi.Any{
					Template: []map[string]interface{}{
						map[string]interface{}{
							"template": map[string]interface{}{
								"deviceIdExpression":  "$.deviceid",
								"timestampExpression": "$.measurementdatetime",
								"typeMatchExpression": "$..[?(@heartrate)]",
								"typeName":            "heartrate",
								"values": []map[string]interface{}{
									map[string]interface{}{
										"required":        "true",
										"valueExpression": "$.heartrate",
										"valueName":       "hr",
									},
								},
							},
							"templateType": "JsonPathContent",
						},
					},
					TemplateType: "CollectionContent",
				},
			},
			Identity: &healthcareapis.ServiceManagedIdentityIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			IngestionEndpointConfiguration: &healthcareapis.IotEventHubIngestionEndpointConfigurationArgs{
				ConsumerGroup:                   pulumi.String("ConsumerGroupA"),
				EventHubName:                    pulumi.String("MyEventHubName"),
				FullyQualifiedEventHubNamespace: pulumi.String("myeventhub.servicesbus.windows.net"),
			},
			IotConnectorName:  pulumi.String("blue"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("testRG"),
			Tags: pulumi.StringMap{
				"additionalProp1": pulumi.String("string"),
				"additionalProp2": pulumi.String("string"),
				"additionalProp3": pulumi.String("string"),
			},
			WorkspaceName: pulumi.String("workspace1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
