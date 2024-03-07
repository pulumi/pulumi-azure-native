package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewInput(ctx, "input", &streamanalytics.InputArgs{
			InputName: pulumi.String("input7225"),
			JobName:   pulumi.String("sj9597"),
			Properties: streamanalytics.ReferenceInputProperties{
				Datasource: streamanalytics.BlobReferenceInputDataSource{
					Container:   "state",
					DateFormat:  "yyyy/MM/dd",
					PathPattern: "{date}/{time}",
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
				Type: "Reference",
			},
			ResourceGroupName: pulumi.String("sjrg8440"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
