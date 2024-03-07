package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewSchemaRegistry(ctx, "schemaRegistry", &eventhub.SchemaRegistryArgs{
			GroupProperties:     nil,
			NamespaceName:       pulumi.String("ali-ua-test-eh-system-1"),
			ResourceGroupName:   pulumi.String("alitest"),
			SchemaCompatibility: pulumi.String("Forward"),
			SchemaGroupName:     pulumi.String("testSchemaGroup1"),
			SchemaType:          pulumi.String("Avro"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
