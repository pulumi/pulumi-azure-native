package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewInput(ctx, "input", &streamanalytics.InputArgs{
			InputName: pulumi.String("input8899"),
			JobName:   pulumi.String("sj6695"),
			Properties: streamanalytics.StreamInputProperties{
				Datasource: streamanalytics.BlobStreamInputDataSource{
					Container:            "state",
					DateFormat:           "yyyy/MM/dd",
					PathPattern:          "{date}/{time}",
					SourcePartitionCount: 16,
					StorageAccounts: []streamanalytics.StorageAccount{
						{
							AccountKey:  "someAccountKey==",
							AccountName: "someAccountName",
						},
					},
					TimeFormat: "HH",
					Type:       "Microsoft.Storage/Blob",
				},
				Serialization: streamanalytics.CsvSerialization{
					Encoding:       "UTF8",
					FieldDelimiter: ",",
					Type:           "Csv",
				},
				Type: "Stream",
			},
			ResourceGroupName: pulumi.String("sjrg8161"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
