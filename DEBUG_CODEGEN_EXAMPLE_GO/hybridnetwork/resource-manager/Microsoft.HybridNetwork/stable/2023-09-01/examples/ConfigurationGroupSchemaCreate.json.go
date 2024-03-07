package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewConfigurationGroupSchema(ctx, "configurationGroupSchema", &hybridnetwork.ConfigurationGroupSchemaArgs{
			ConfigurationGroupSchemaName: pulumi.String("testConfigurationGroupSchema"),
			Location:                     pulumi.String("westUs2"),
			Properties: &hybridnetwork.ConfigurationGroupSchemaPropertiesFormatArgs{
				Description:      pulumi.String("Schema with no secrets"),
				SchemaDefinition: pulumi.String("{\"type\":\"object\",\"properties\":{\"interconnect-groups\":{\"type\":\"object\",\"properties\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"international-interconnects\":{\"type\":\"array\",\"item\":{\"type\":\"string\"}},\"domestic-interconnects\":{\"type\":\"array\",\"item\":{\"type\":\"string\"}}}}},\"interconnect-group-assignments\":{\"type\":\"object\",\"properties\":{\"type\":\"object\",\"properties\":{\"ssc\":{\"type\":\"string\"},\"interconnects-interconnects\":{\"type\":\"string\"}}}}},\"required\":[\"interconnect-groups\",\"interconnect-group-assignments\"]}"),
			},
			PublisherName:     pulumi.String("testPublisher"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
