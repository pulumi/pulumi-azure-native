package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logz/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logz.NewMetricsSource(ctx, "metricsSource", &logz.MetricsSourceArgs{
			MetricsSourceName: pulumi.String("MetricsSource1"),
			MonitorName:       pulumi.String("myMonitor"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
