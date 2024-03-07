package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewProtectionPolicy(ctx, "protectionPolicy", &recoveryservices.ProtectionPolicyArgs{
			PolicyName: pulumi.String("testPolicy1"),
			Properties: recoveryservices.AzureIaaSVMProtectionPolicy{
				BackupManagementType: "AzureIaasVM",
				RetentionPolicy: recoveryservices.LongTermRetentionPolicy{
					DailySchedule: recoveryservices.DailyRetentionSchedule{
						RetentionDuration: recoveryservices.RetentionDuration{
							Count:        1,
							DurationType: "Days",
						},
						RetentionTimes: []string{
							"2018-01-24T02:00:00Z",
						},
					},
					RetentionPolicyType: "LongTermRetentionPolicy",
				},
				SchedulePolicy: recoveryservices.SimpleSchedulePolicy{
					SchedulePolicyType:   "SimpleSchedulePolicy",
					ScheduleRunFrequency: "Daily",
					ScheduleRunTimes: []string{
						"2018-01-24T02:00:00Z",
					},
				},
				TimeZone: "Pacific Standard Time",
			},
			ResourceGroupName: pulumi.String("SwaggerTestRg"),
			VaultName:         pulumi.String("NetSDKTestRsVault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
