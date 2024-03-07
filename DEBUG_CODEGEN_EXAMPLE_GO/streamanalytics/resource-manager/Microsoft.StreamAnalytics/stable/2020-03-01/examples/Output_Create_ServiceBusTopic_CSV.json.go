package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewOutput(ctx, "output", &streamanalytics.OutputArgs{
			Datasource: streamanalytics.ServiceBusTopicOutputDataSource{
				PropertyColumns: []string{
					"column1",
					"column2",
				},
				ServiceBusNamespace:    "sdktest",
				SharedAccessPolicyKey:  "sharedAccessPolicyKey=",
				SharedAccessPolicyName: "RootManageSharedAccessKey",
				TopicName:              "sdktopic",
				Type:                   "Microsoft.ServiceBus/Topic",
			},
			JobName:           pulumi.String("sj7094"),
			OutputName:        pulumi.String("output7886"),
			ResourceGroupName: pulumi.String("sjrg6450"),
			Serialization: streamanalytics.CsvSerialization{
				Encoding:       "UTF8",
				FieldDelimiter: ",",
				Type:           "Csv",
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
