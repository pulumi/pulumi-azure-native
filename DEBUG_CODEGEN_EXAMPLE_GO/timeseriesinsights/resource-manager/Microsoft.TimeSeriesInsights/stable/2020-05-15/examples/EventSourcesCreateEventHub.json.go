package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/timeseriesinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := timeseriesinsights.NewIoTHubEventSource(ctx, "ioTHubEventSource", &timeseriesinsights.IoTHubEventSourceArgs{
			EnvironmentName:   pulumi.String("env1"),
			EventSourceName:   pulumi.String("es1"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
