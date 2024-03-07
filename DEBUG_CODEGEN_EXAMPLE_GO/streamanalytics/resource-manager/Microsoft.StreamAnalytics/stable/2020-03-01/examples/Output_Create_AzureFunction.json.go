package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewOutput(ctx, "output", &streamanalytics.OutputArgs{
			Datasource: streamanalytics.AzureFunctionOutputDataSource{
				FunctionAppName: "functionappforasaautomation",
				FunctionName:    "HttpTrigger2",
				MaxBatchCount:   100,
				MaxBatchSize:    256,
				Type:            "Microsoft.AzureFunction",
			},
			JobName:           pulumi.String("sjName"),
			OutputName:        pulumi.String("azureFunction1"),
			ResourceGroupName: pulumi.String("sjrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
