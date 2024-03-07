package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewInput(ctx, "input", &streamanalytics.InputArgs{
			InputName: pulumi.String("input7425"),
			JobName:   pulumi.String("sj197"),
			Properties: streamanalytics.StreamInputProperties{
				Datasource: streamanalytics.EventHubStreamInputDataSource{
					ConsumerGroupName:      "sdkconsumergroup",
					EventHubName:           "sdkeventhub",
					ServiceBusNamespace:    "sdktest",
					SharedAccessPolicyKey:  "someSharedAccessPolicyKey==",
					SharedAccessPolicyName: "RootManageSharedAccessKey",
					Type:                   "Microsoft.ServiceBus/EventHub",
				},
				Serialization: streamanalytics.JsonSerialization{
					Encoding: "UTF8",
					Type:     "Json",
				},
				Type: "Stream",
			},
			ResourceGroupName: pulumi.String("sjrg3139"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
