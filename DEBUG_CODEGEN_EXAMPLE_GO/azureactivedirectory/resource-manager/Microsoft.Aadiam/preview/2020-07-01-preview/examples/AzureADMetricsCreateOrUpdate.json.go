package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/aadiam/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aadiam.NewAzureADMetric(ctx, "azureADMetric", &aadiam.AzureADMetricArgs{
			AzureADMetricsName: pulumi.String("ddb1"),
			Location:           pulumi.String("West US"),
			ResourceGroupName:  pulumi.String("rg1"),
			Tags:               nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
