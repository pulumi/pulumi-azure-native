package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewOutput(ctx, "output", &streamanalytics.OutputArgs{
			Datasource: streamanalytics.AzureTableOutputDataSource{
				AccountKey:  "accountKey==",
				AccountName: "someAccountName",
				BatchSize:   25,
				ColumnsToRemove: []string{
					"column1",
					"column2",
				},
				PartitionKey: "partitionKey",
				RowKey:       "rowKey",
				Table:        "samples",
				Type:         "Microsoft.Storage/Table",
			},
			JobName:           pulumi.String("sj2790"),
			OutputName:        pulumi.String("output958"),
			ResourceGroupName: pulumi.String("sjrg5176"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
