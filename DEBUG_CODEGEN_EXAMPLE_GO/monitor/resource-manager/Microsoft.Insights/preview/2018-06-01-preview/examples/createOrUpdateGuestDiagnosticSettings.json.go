package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewGuestDiagnosticsSetting(ctx, "guestDiagnosticsSetting", &insights.GuestDiagnosticsSettingArgs{
			DataSources: insights.DataSourceArray{
				&insights.DataSourceArgs{
					Configuration: &insights.DataSourceConfigurationArgs{
						PerfCounters: insights.PerformanceCounterConfigurationArray{
							&insights.PerformanceCounterConfigurationArgs{
								Name:           pulumi.String("\\Process(_Total)\\%Processor Time"),
								SamplingPeriod: pulumi.String("PT1M"),
							},
							&insights.PerformanceCounterConfigurationArgs{
								Name:           pulumi.String("\\Process(_Total)\\Working Set"),
								SamplingPeriod: pulumi.String("PT1M"),
							},
						},
					},
					Kind: pulumi.String("PerformanceCounter"),
					Sinks: insights.SinkConfigurationArray{
						&insights.SinkConfigurationArgs{
							Kind: pulumi.String("LogAnalytics"),
						},
					},
				},
				&insights.DataSourceArgs{
					Configuration: &insights.DataSourceConfigurationArgs{
						Providers: insights.EtwProviderConfigurationArray{
							&insights.EtwProviderConfigurationArgs{
								Id: pulumi.String("1"),
							},
							&insights.EtwProviderConfigurationArgs{
								Id: pulumi.String("2"),
							},
						},
					},
					Kind: pulumi.String("ETWProviders"),
					Sinks: insights.SinkConfigurationArray{
						&insights.SinkConfigurationArgs{
							Kind: pulumi.String("LogAnalytics"),
						},
					},
				},
				&insights.DataSourceArgs{
					Configuration: &insights.DataSourceConfigurationArgs{
						EventLogs: insights.EventLogConfigurationArray{
							&insights.EventLogConfigurationArgs{
								Filter:  pulumi.String("SourceName == Xyz AND EventId = \"100\" AND  $Xpath/Column=\"DCName\" = \"CatWoman\""),
								LogName: pulumi.String("Application"),
							},
							&insights.EventLogConfigurationArgs{
								Filter:  pulumi.String("SourceName == Xyz AND EventId = \"100\" AND  $Xpath/Column=\"DCName\" = \"BatMan\""),
								LogName: pulumi.String("Application"),
							},
						},
					},
					Kind: pulumi.String("WindowsEventLogs"),
					Sinks: insights.SinkConfigurationArray{
						&insights.SinkConfigurationArgs{
							Kind: pulumi.String("LogAnalytics"),
						},
					},
				},
			},
			DiagnosticSettingsName: pulumi.String("SampleDiagSetting"),
			Location:               pulumi.String("Global"),
			OsType:                 pulumi.String("Windows"),
			ResourceGroupName:      pulumi.String("Default-ResourceGroup"),
			Tags:                   nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
