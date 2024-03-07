package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewConfigurationGroupValue(ctx, "configurationGroupValue", &hybridnetwork.ConfigurationGroupValueArgs{
			ConfigurationGroupValueName: pulumi.String("testConfigurationGroupValue"),
			Location:                    pulumi.String("eastus"),
			Properties: hybridnetwork.ConfigurationValueWithSecrets{
				ConfigurationGroupSchemaResourceReference: hybridnetwork.OpenDeploymentResourceReference{
					Id:     "/subscriptions/subid/resourcegroups/testRG/providers/microsoft.hybridnetwork/publishers/testPublisher/configurationGroupSchemas/testConfigurationGroupSchemaName",
					IdType: "Open",
				},
				ConfigurationType:        "Secret",
				SecretConfigurationValue: "{\"interconnect-groups\":{\"stripe-one\":{\"name\":\"Stripe one\",\"international-interconnects\":[\"france\",\"germany\"],\"domestic-interconnects\":[\"birmingham\",\"edinburgh\"]},\"stripe-two\":{\"name\":\"Stripe two\",\"international-interconnects\":[\"germany\",\"italy\"],\"domestic-interconnects\":[\"edinburgh\",\"london\"]}},\"interconnect-group-assignments\":{\"ssc-one\":{\"ssc\":\"SSC 1\",\"interconnects\":\"stripe-one\"},\"ssc-two\":{\"ssc\":\"SSC 2\",\"interconnects\":\"stripe-two\"}}}",
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
