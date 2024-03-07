package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewMonitoringSetting(ctx, "monitoringSetting", &appplatform.MonitoringSettingArgs{
			Properties: &appplatform.MonitoringSettingPropertiesArgs{
				AppInsightsInstrumentationKey: pulumi.String("00000000-0000-0000-0000-000000000000"),
				AppInsightsSamplingRate:       pulumi.Float64(10),
				TraceEnabled:                  pulumi.Bool(true),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
