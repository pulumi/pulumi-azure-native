package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewJobSchedule(ctx, "jobSchedule", &automation.JobScheduleArgs{
			AutomationAccountName: pulumi.String("ContoseAutomationAccount"),
			JobScheduleId:         pulumi.String("0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc"),
			Parameters: pulumi.StringMap{
				"jobscheduletag01": pulumi.String("jobschedulevalue01"),
				"jobscheduletag02": pulumi.String("jobschedulevalue02"),
			},
			ResourceGroupName: pulumi.String("rg"),
			Runbook: &automation.RunbookAssociationPropertyArgs{
				Name: pulumi.String("TestRunbook"),
			},
			Schedule: &automation.ScheduleAssociationPropertyArgs{
				Name: pulumi.String("ScheduleNameGoesHere332204b5-debe-4348-a5c7-6357457189f2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
