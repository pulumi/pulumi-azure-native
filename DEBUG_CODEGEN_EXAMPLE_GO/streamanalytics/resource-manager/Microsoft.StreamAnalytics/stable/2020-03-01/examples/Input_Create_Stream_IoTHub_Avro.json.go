package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewInput(ctx, "input", &streamanalytics.InputArgs{
			InputName: pulumi.String("input7970"),
			JobName:   pulumi.String("sj9742"),
			Properties: streamanalytics.StreamInputProperties{
				Datasource: streamanalytics.IoTHubStreamInputDataSource{
					ConsumerGroupName:      "sdkconsumergroup",
					Endpoint:               "messages/events",
					IotHubNamespace:        "iothub",
					SharedAccessPolicyKey:  "sharedAccessPolicyKey=",
					SharedAccessPolicyName: "owner",
					Type:                   "Microsoft.Devices/IotHubs",
				},
				Serialization: streamanalytics.AvroSerialization{
					Type: "Avro",
				},
				Type: "Stream",
			},
			ResourceGroupName: pulumi.String("sjrg3467"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
