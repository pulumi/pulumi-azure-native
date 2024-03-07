package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dataprotection/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dataprotection.NewBackupPolicy(ctx, "backupPolicy", &dataprotection.BackupPolicyArgs{
			BackupPolicyName: pulumi.String("OSSDBPolicy"),
			Properties: &dataprotection.BackupPolicyTypeArgs{
				DatasourceTypes: pulumi.StringArray{
					pulumi.String("OssDB"),
				},
				ObjectType: pulumi.String("BackupPolicy"),
				PolicyRules: pulumi.Array{
					dataprotection.AzureBackupRule{
						BackupParameters: dataprotection.AzureBackupParams{
							BackupType: "Full",
							ObjectType: "AzureBackupParams",
						},
						DataStore: dataprotection.DataStoreInfoBase{
							DataStoreType: "VaultStore",
							ObjectType:    "DataStoreInfoBase",
						},
						Name:       "BackupWeekly",
						ObjectType: "AzureBackupRule",
						Trigger: dataprotection.ScheduleBasedTriggerContext{
							ObjectType: "ScheduleBasedTriggerContext",
							Schedule: dataprotection.BackupSchedule{
								RepeatingTimeIntervals: []string{
									"R/2019-11-20T08:00:00-08:00/P1W",
								},
							},
							TaggingCriteria: []dataprotection.TaggingCriteria{
								{
									IsDefault: true,
									TagInfo: {
										TagName: "Default",
									},
									TaggingPriority: 99,
								},
								{
									Criteria: []dataprotection.ScheduleBasedBackupCriteria{
										{
											DaysOfTheWeek: []dataprotection.DayOfWeek{
												"Sunday",
											},
											ObjectType: "ScheduleBasedBackupCriteria",
											ScheduleTimes: []string{
												"2019-03-01T13:00:00Z",
											},
										},
									},
									IsDefault: false,
									TagInfo: {
										TagName: "Weekly",
									},
									TaggingPriority: 20,
								},
							},
						},
					},
					dataprotection.AzureRetentionRule{
						IsDefault: true,
						Lifecycles: []dataprotection.SourceLifeCycle{
							{
								DeleteAfter: {
									Duration:   "P1W",
									ObjectType: "AbsoluteDeleteOption",
								},
								SourceDataStore: {
									DataStoreType: "VaultStore",
									ObjectType:    "DataStoreInfoBase",
								},
							},
						},
						Name:       "Default",
						ObjectType: "AzureRetentionRule",
					},
					dataprotection.AzureRetentionRule{
						IsDefault: false,
						Lifecycles: []dataprotection.SourceLifeCycle{
							{
								DeleteAfter: {
									Duration:   "P12W",
									ObjectType: "AbsoluteDeleteOption",
								},
								SourceDataStore: {
									DataStoreType: "VaultStore",
									ObjectType:    "DataStoreInfoBase",
								},
							},
						},
						Name:       "Weekly",
						ObjectType: "AzureRetentionRule",
					},
				},
			},
			ResourceGroupName: pulumi.String("000pikumar"),
			VaultName:         pulumi.String("PrivatePreviewVault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
