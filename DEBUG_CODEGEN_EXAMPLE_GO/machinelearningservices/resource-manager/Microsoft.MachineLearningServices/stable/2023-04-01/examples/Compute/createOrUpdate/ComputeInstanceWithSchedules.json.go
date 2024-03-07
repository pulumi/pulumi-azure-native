package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewCompute(ctx, "compute", &machinelearningservices.ComputeArgs{
			ComputeName: pulumi.String("compute123"),
			Location:    pulumi.String("eastus"),
			Properties: machinelearningservices.ComputeInstance{
				ComputeType: "ComputeInstance",
				Properties: machinelearningservices.ComputeInstanceProperties{
					ApplicationSharingPolicy:         "Personal",
					ComputeInstanceAuthorizationType: "personal",
					PersonalComputeInstanceSettings: machinelearningservices.PersonalComputeInstanceSettings{
						AssignedUser: machinelearningservices.AssignedUser{
							ObjectId: "00000000-0000-0000-0000-000000000000",
							TenantId: "00000000-0000-0000-0000-000000000000",
						},
					},
					Schedules: machinelearningservices.ComputeSchedules{
						ComputeStartStop: []machinelearningservices.ComputeStartStopSchedule{
							{
								Action: "Stop",
								Cron: {
									Expression: "0 18 * * *",
									StartTime:  "2021-04-23T01:30:00",
									TimeZone:   "Pacific Standard Time",
								},
								Status:      "Enabled",
								TriggerType: "Cron",
							},
						},
					},
					SshSettings: machinelearningservices.ComputeInstanceSshSettings{
						SshPublicAccess: "Disabled",
					},
					VmSize: "STANDARD_NC6",
				},
			},
			ResourceGroupName: pulumi.String("testrg123"),
			WorkspaceName:     pulumi.String("workspaces123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
