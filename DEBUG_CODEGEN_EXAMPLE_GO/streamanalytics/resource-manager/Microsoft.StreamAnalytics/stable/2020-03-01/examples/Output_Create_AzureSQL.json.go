package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewOutput(ctx, "output", &streamanalytics.OutputArgs{
			Datasource: streamanalytics.AzureSqlDatabaseOutputDataSource{
				Database: "someDatabase",
				Password: "somePassword",
				Server:   "someServer",
				Table:    "someTable",
				Type:     "Microsoft.Sql/Server/Database",
				User:     "<user>",
			},
			JobName:           pulumi.String("sj6458"),
			OutputName:        pulumi.String("output1755"),
			ResourceGroupName: pulumi.String("sjrg2157"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
