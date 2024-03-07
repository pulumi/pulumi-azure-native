package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewSchedule(ctx, "schedule", &machinelearningservices.ScheduleArgs{
			Name:              pulumi.String("string"),
			ResourceGroupName: pulumi.String("test-rg"),
			ScheduleProperties: &machinelearningservices.ScheduleTypeArgs{
				Action: machinelearningservices.EndpointScheduleAction{
					ActionType: "InvokeBatchEndpoint",
					EndpointInvocationDefinition: map[string]interface{}{
						"9965593e-526f-4b89-bb36-761138cf2794": nil,
					},
				},
				Description: pulumi.String("string"),
				DisplayName: pulumi.String("string"),
				IsEnabled:   pulumi.Bool(false),
				Properties: pulumi.StringMap{
					"string": pulumi.String("string"),
				},
				Tags: pulumi.StringMap{
					"string": pulumi.String("string"),
				},
				Trigger: machinelearningservices.CronTrigger{
					EndTime:     "string",
					Expression:  "string",
					StartTime:   "string",
					TimeZone:    "string",
					TriggerType: "Cron",
				},
			},
			WorkspaceName: pulumi.String("my-aml-workspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
