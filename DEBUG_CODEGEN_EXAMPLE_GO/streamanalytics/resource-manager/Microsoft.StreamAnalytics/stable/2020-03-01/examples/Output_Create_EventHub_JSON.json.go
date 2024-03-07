package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewOutput(ctx, "output", &streamanalytics.OutputArgs{
			Datasource: streamanalytics.EventHubOutputDataSource{
				EventHubName:           "sdkeventhub",
				PartitionKey:           "partitionKey",
				ServiceBusNamespace:    "sdktest",
				SharedAccessPolicyKey:  "sharedAccessPolicyKey=",
				SharedAccessPolicyName: "RootManageSharedAccessKey",
				Type:                   "Microsoft.ServiceBus/EventHub",
			},
			JobName:           pulumi.String("sj3310"),
			OutputName:        pulumi.String("output5195"),
			ResourceGroupName: pulumi.String("sjrg6912"),
			Serialization: streamanalytics.JsonSerialization{
				Encoding: "UTF8",
				Format:   "Array",
				Type:     "Json",
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
