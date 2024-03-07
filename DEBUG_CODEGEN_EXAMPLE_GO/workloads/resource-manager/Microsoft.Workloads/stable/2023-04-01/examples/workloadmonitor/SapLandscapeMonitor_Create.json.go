package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewSapLandscapeMonitor(ctx, "sapLandscapeMonitor", &workloads.SapLandscapeMonitorArgs{
			Grouping: &workloads.SapLandscapeMonitorPropertiesGroupingArgs{
				Landscape: workloads.SapLandscapeMonitorSidMappingArray{
					&workloads.SapLandscapeMonitorSidMappingArgs{
						Name: pulumi.String("Prod"),
						TopSid: pulumi.StringArray{
							pulumi.String("SID1"),
							pulumi.String("SID2"),
						},
					},
				},
				SapApplication: workloads.SapLandscapeMonitorSidMappingArray{
					&workloads.SapLandscapeMonitorSidMappingArgs{
						Name: pulumi.String("ERP1"),
						TopSid: pulumi.StringArray{
							pulumi.String("SID1"),
							pulumi.String("SID2"),
						},
					},
				},
			},
			MonitorName:       pulumi.String("mySapMonitor"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			TopMetricsThresholds: workloads.SapLandscapeMonitorMetricThresholdsArray{
				&workloads.SapLandscapeMonitorMetricThresholdsArgs{
					Green:  pulumi.Float64(90),
					Name:   pulumi.String("Instance Availability"),
					Red:    pulumi.Float64(50),
					Yellow: pulumi.Float64(75),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
