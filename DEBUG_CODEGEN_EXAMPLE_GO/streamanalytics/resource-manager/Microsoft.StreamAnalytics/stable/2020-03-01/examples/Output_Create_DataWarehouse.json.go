package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewOutput(ctx, "output", &streamanalytics.OutputArgs{
			Datasource: streamanalytics.AzureSynapseOutputDataSource{
				Database: "zhayaSQLpool",
				Password: "password123",
				Server:   "asatestserver",
				Table:    "test2",
				Type:     "Microsoft.Sql/Server/DataWarehouse",
				User:     "tolladmin",
			},
			JobName:           pulumi.String("sjName"),
			OutputName:        pulumi.String("dwOutput"),
			ResourceGroupName: pulumi.String("sjrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
