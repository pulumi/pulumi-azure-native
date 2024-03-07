package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewOutput(ctx, "output", &streamanalytics.OutputArgs{
			Datasource: streamanalytics.BlobOutputDataSource{
				Container:   "state",
				DateFormat:  "yyyy/MM/dd",
				PathPattern: "{date}/{time}",
				StorageAccounts: []streamanalytics.StorageAccount{
					{
						AccountKey:  "accountKey==",
						AccountName: "someAccountName",
					},
				},
				TimeFormat: "HH",
				Type:       "Microsoft.Storage/Blob",
			},
			JobName:           pulumi.String("sj900"),
			OutputName:        pulumi.String("output1623"),
			ResourceGroupName: pulumi.String("sjrg5023"),
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
