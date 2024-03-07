package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewOutput(ctx, "output", &streamanalytics.OutputArgs{
			Datasource: streamanalytics.ServiceBusQueueOutputDataSource{
				PropertyColumns: []string{
					"column1",
					"column2",
				},
				QueueName:              "sdkqueue",
				ServiceBusNamespace:    "sdktest",
				SharedAccessPolicyKey:  "sharedAccessPolicyKey=",
				SharedAccessPolicyName: "RootManageSharedAccessKey",
				SystemPropertyColumns: map[string]interface{}{
					"MessageId":    "col3",
					"PartitionKey": "col4",
				},
				Type: "Microsoft.ServiceBus/Queue",
			},
			JobName:           pulumi.String("sj5095"),
			OutputName:        pulumi.String("output3456"),
			ResourceGroupName: pulumi.String("sjrg3410"),
			Serialization: streamanalytics.AvroSerialization{
				Type: "Avro",
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
