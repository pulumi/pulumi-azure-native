package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewSoftwareUpdateConfigurationByName(ctx, "softwareUpdateConfigurationByName", &automation.SoftwareUpdateConfigurationByNameArgs{
			AutomationAccountName: pulumi.String("myaccount"),
			ResourceGroupName:     pulumi.String("mygroup"),
			ScheduleInfo: &automation.SUCSchedulePropertiesArgs{
				AdvancedSchedule: &automation.AdvancedScheduleArgs{
					WeekDays: pulumi.StringArray{
						pulumi.String("Monday"),
						pulumi.String("Thursday"),
					},
				},
				ExpiryTime: pulumi.String("2018-11-09T11:22:57+00:00"),
				Frequency:  pulumi.String("Hour"),
				Interval:   pulumi.Float64(1),
				StartTime:  pulumi.String("2017-10-19T12:22:57+00:00"),
				TimeZone:   pulumi.String("America/Los_Angeles"),
			},
			SoftwareUpdateConfigurationName: pulumi.String("testpatch"),
			Tasks: &automation.SoftwareUpdateConfigurationTasksArgs{
				PostTask: &automation.TaskPropertiesArgs{
					Source: pulumi.String("GetCache"),
				},
				PreTask: &automation.TaskPropertiesArgs{
					Parameters: pulumi.StringMap{
						"COMPUTERNAME": pulumi.String("Computer1"),
					},
					Source: pulumi.String("HelloWorld"),
				},
			},
			UpdateConfiguration: &automation.UpdateConfigurationArgs{
				AzureVirtualMachines: pulumi.StringArray{
					pulumi.String("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01"),
					pulumi.String("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-02"),
					pulumi.String("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-03"),
				},
				Duration: pulumi.String("PT2H0M"),
				NonAzureComputerNames: pulumi.StringArray{
					pulumi.String("box1.contoso.com"),
					pulumi.String("box2.contoso.com"),
				},
				OperatingSystem: automation.OperatingSystemTypeWindows,
				Targets: &automation.TargetPropertiesArgs{
					AzureQueries: automation.AzureQueryPropertiesArray{
						&automation.AzureQueryPropertiesArgs{
							Locations: pulumi.StringArray{
								pulumi.String("Japan East"),
								pulumi.String("UK South"),
							},
							Scope: pulumi.StringArray{
								pulumi.String("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources"),
								pulumi.String("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067"),
							},
							TagSettings: &automation.TagSettingsPropertiesArgs{
								FilterOperator: automation.TagOperatorsAll,
								Tags: pulumi.StringArrayMap{
									"tag1": pulumi.StringArray{
										pulumi.String("tag1Value1"),
										pulumi.String("tag1Value2"),
										pulumi.String("tag1Value3"),
									},
									"tag2": pulumi.StringArray{
										pulumi.String("tag2Value1"),
										pulumi.String("tag2Value2"),
										pulumi.String("tag2Value3"),
									},
								},
							},
						},
					},
					NonAzureQueries: automation.NonAzureQueryPropertiesArray{
						&automation.NonAzureQueryPropertiesArgs{
							FunctionAlias: pulumi.String("SavedSearch1"),
							WorkspaceId:   pulumi.String("WorkspaceId1"),
						},
						&automation.NonAzureQueryPropertiesArgs{
							FunctionAlias: pulumi.String("SavedSearch2"),
							WorkspaceId:   pulumi.String("WorkspaceId2"),
						},
					},
				},
				Windows: &automation.WindowsPropertiesArgs{
					ExcludedKbNumbers: pulumi.StringArray{
						pulumi.String("168934"),
						pulumi.String("168973"),
					},
					IncludedUpdateClassifications: pulumi.String("Critical"),
					RebootSetting:                 pulumi.String("IfRequired"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
