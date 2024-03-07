package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databoxedge/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databoxedge.NewMonitoringConfig(ctx, "monitoringConfig", &databoxedge.MonitoringConfigArgs{
			DeviceName: pulumi.String("testedgedevice"),
			MetricConfigurations: databoxedge.MetricConfigurationArray{
				&databoxedge.MetricConfigurationArgs{
					CounterSets: databoxedge.MetricCounterSetArray{
						&databoxedge.MetricCounterSetArgs{
							Counters: databoxedge.MetricCounterArray{
								&databoxedge.MetricCounterArgs{
									Name: pulumi.String("test"),
								},
							},
						},
					},
					MdmAccount:      pulumi.String("test"),
					MetricNameSpace: pulumi.String("test"),
					ResourceId:      pulumi.String("test"),
				},
			},
			ResourceGroupName: pulumi.String("GroupForEdgeAutomation"),
			RoleName:          pulumi.String("testrole"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
