package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/scom/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := scom.NewMonitoredResource(ctx, "monitoredResource", &scom.MonitoredResourceArgs{
			InstanceName:          pulumi.String("myInstance"),
			MonitoredResourceName: pulumi.String("d877b154-9a8d-4bfe-8a24-20682fcf2ed3"),
			ResourceGroupName:     pulumi.String("myResGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
