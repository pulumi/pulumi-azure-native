package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewDataCollectionRule(ctx, "dataCollectionRule", &insights.DataCollectionRuleArgs{
			DataCollectionRuleName: pulumi.String("myCollectionRule"),
			DataFlows: insights.DataFlowArray{
				&insights.DataFlowArgs{
					Destinations: pulumi.StringArray{
						pulumi.String("centralWorkspace"),
					},
					Streams: pulumi.StringArray{
						pulumi.String("Microsoft-Perf"),
						pulumi.String("Microsoft-Syslog"),
						pulumi.String("Microsoft-WindowsEvent"),
					},
				},
			},
			DataSources: &insights.DataCollectionRuleDataSourcesArgs{
				PerformanceCounters: insights.PerfCounterDataSourceArray{
					&insights.PerfCounterDataSourceArgs{
						CounterSpecifiers: pulumi.StringArray{
							pulumi.String("\\Processor(_Total)\\% Processor Time"),
							pulumi.String("\\Memory\\Committed Bytes"),
							pulumi.String("\\LogicalDisk(_Total)\\Free Megabytes"),
							pulumi.String("\\PhysicalDisk(_Total)\\Avg. Disk Queue Length"),
						},
						Name:                       pulumi.String("cloudTeamCoreCounters"),
						SamplingFrequencyInSeconds: pulumi.Int(15),
						Streams: pulumi.StringArray{
							pulumi.String("Microsoft-Perf"),
						},
					},
					&insights.PerfCounterDataSourceArgs{
						CounterSpecifiers: pulumi.StringArray{
							pulumi.String("\\Process(_Total)\\Thread Count"),
						},
						Name:                       pulumi.String("appTeamExtraCounters"),
						SamplingFrequencyInSeconds: pulumi.Int(30),
						Streams: pulumi.StringArray{
							pulumi.String("Microsoft-Perf"),
						},
					},
				},
				Syslog: insights.SyslogDataSourceArray{
					&insights.SyslogDataSourceArgs{
						FacilityNames: pulumi.StringArray{
							pulumi.String("cron"),
						},
						LogLevels: pulumi.StringArray{
							pulumi.String("Debug"),
							pulumi.String("Critical"),
							pulumi.String("Emergency"),
						},
						Name: pulumi.String("cronSyslog"),
						Streams: pulumi.StringArray{
							pulumi.String("Microsoft-Syslog"),
						},
					},
					&insights.SyslogDataSourceArgs{
						FacilityNames: pulumi.StringArray{
							pulumi.String("syslog"),
						},
						LogLevels: pulumi.StringArray{
							pulumi.String("Alert"),
							pulumi.String("Critical"),
							pulumi.String("Emergency"),
						},
						Name: pulumi.String("syslogBase"),
						Streams: pulumi.StringArray{
							pulumi.String("Microsoft-Syslog"),
						},
					},
				},
				WindowsEventLogs: insights.WindowsEventLogDataSourceArray{
					&insights.WindowsEventLogDataSourceArgs{
						Name: pulumi.String("cloudSecurityTeamEvents"),
						Streams: pulumi.StringArray{
							pulumi.String("Microsoft-WindowsEvent"),
						},
						XPathQueries: pulumi.StringArray{
							pulumi.String("Security!"),
						},
					},
					&insights.WindowsEventLogDataSourceArgs{
						Name: pulumi.String("appTeam1AppEvents"),
						Streams: pulumi.StringArray{
							pulumi.String("Microsoft-WindowsEvent"),
						},
						XPathQueries: pulumi.StringArray{
							pulumi.String("System![System[(Level = 1 or Level = 2 or Level = 3)]]"),
							pulumi.String("Application!*[System[(Level = 1 or Level = 2 or Level = 3)]]"),
						},
					},
				},
			},
			Destinations: &insights.DataCollectionRuleDestinationsArgs{
				LogAnalytics: insights.LogAnalyticsDestinationArray{
					&insights.LogAnalyticsDestinationArgs{
						Name:                pulumi.String("centralWorkspace"),
						WorkspaceResourceId: pulumi.String("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/centralTeamWorkspace"),
					},
				},
			},
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
