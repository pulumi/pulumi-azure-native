package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/elastic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elastic.NewMonitor(ctx, "monitor", &elastic.MonitorArgs{
			MonitorName:       pulumi.String("myMonitor"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
