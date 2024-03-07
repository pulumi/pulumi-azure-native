package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewStreamingJob(ctx, "streamingJob", &streamanalytics.StreamingJobArgs{
			CompatibilityLevel:                 pulumi.String("1.0"),
			DataLocale:                         pulumi.String("en-US"),
			EventsLateArrivalMaxDelayInSeconds: pulumi.Int(5),
			EventsOutOfOrderMaxDelayInSeconds:  pulumi.Int(0),
			EventsOutOfOrderPolicy:             pulumi.String("Drop"),
			Functions:                          streamanalytics.FunctionTypeArray{},
			Inputs: streamanalytics.InputTypeArray{
				&streamanalytics.InputTypeArgs{
					Name: pulumi.String("inputtest"),
					Properties: streamanalytics.StreamInputProperties{
						Datasource: streamanalytics.BlobStreamInputDataSource{
							Container:   "containerName",
							PathPattern: "",
							StorageAccounts: []streamanalytics.StorageAccount{
								{
									AccountKey:  "yourAccountKey==",
									AccountName: "yourAccountName",
								},
							},
							Type: "Microsoft.Storage/Blob",
						},
						Serialization: streamanalytics.JsonSerialization{
							Encoding: "UTF8",
							Type:     "Json",
						},
						Type: "Stream",
					},
				},
			},
			JobName:           pulumi.String("sj7804"),
			Location:          pulumi.String("West US"),
			OutputErrorPolicy: pulumi.String("Drop"),
			Outputs: streamanalytics.OutputTypeArray{
				&streamanalytics.OutputTypeArgs{
					Datasource: streamanalytics.AzureSqlDatabaseOutputDataSource{
						Database: "databaseName",
						Password: "userPassword",
						Server:   "serverName",
						Table:    "tableName",
						Type:     "Microsoft.Sql/Server/Database",
						User:     "<user>",
					},
					Name: pulumi.String("outputtest"),
				},
			},
			ResourceGroupName: pulumi.String("sjrg3276"),
			Sku: &streamanalytics.SkuArgs{
				Name: pulumi.String("Standard"),
			},
			Tags: pulumi.StringMap{
				"key1":      pulumi.String("value1"),
				"key3":      pulumi.String("value3"),
				"randomKey": pulumi.String("randomValue"),
			},
			Transformation: &streamanalytics.TransformationArgs{
				Name:           pulumi.String("transformationtest"),
				Query:          pulumi.String("Select Id, Name from inputtest"),
				StreamingUnits: pulumi.Int(1),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
