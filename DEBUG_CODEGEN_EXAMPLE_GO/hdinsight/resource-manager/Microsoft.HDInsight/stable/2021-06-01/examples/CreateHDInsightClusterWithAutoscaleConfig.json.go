package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hdinsight/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hdinsight.NewCluster(ctx, "cluster", &hdinsight.ClusterArgs{
			ClusterName: pulumi.String("cluster1"),
			Properties: &hdinsight.ClusterCreatePropertiesArgs{
				ClusterDefinition: &hdinsight.ClusterDefinitionArgs{
					ComponentVersion: pulumi.StringMap{
						"Hadoop": pulumi.String("2.7"),
					},
					Configurations: pulumi.Any{
						Gateway: map[string]interface{}{
							"restAuthCredential.isEnabled": true,
							"restAuthCredential.password":  "**********",
							"restAuthCredential.username":  "admin",
						},
					},
					Kind: pulumi.String("hadoop"),
				},
				ClusterVersion: pulumi.String("3.6"),
				ComputeProfile: &hdinsight.ComputeProfileArgs{
					Roles: hdinsight.RoleArray{
						&hdinsight.RoleArgs{
							AutoscaleConfiguration: &hdinsight.AutoscaleArgs{
								Recurrence: &hdinsight.AutoscaleRecurrenceArgs{
									Schedule: hdinsight.AutoscaleScheduleArray{
										&hdinsight.AutoscaleScheduleArgs{
											Days: pulumi.StringArray{
												pulumi.String("Monday"),
												pulumi.String("Tuesday"),
												pulumi.String("Wednesday"),
												pulumi.String("Thursday"),
												pulumi.String("Friday"),
											},
											TimeAndCapacity: &hdinsight.AutoscaleTimeAndCapacityArgs{
												MaxInstanceCount: pulumi.Int(3),
												MinInstanceCount: pulumi.Int(3),
												Time:             pulumi.String("09:00"),
											},
										},
										&hdinsight.AutoscaleScheduleArgs{
											Days: pulumi.StringArray{
												pulumi.String("Monday"),
												pulumi.String("Tuesday"),
												pulumi.String("Wednesday"),
												pulumi.String("Thursday"),
												pulumi.String("Friday"),
											},
											TimeAndCapacity: &hdinsight.AutoscaleTimeAndCapacityArgs{
												MaxInstanceCount: pulumi.Int(6),
												MinInstanceCount: pulumi.Int(6),
												Time:             pulumi.String("18:00"),
											},
										},
										&hdinsight.AutoscaleScheduleArgs{
											Days: pulumi.StringArray{
												pulumi.String("Saturday"),
												pulumi.String("Sunday"),
											},
											TimeAndCapacity: &hdinsight.AutoscaleTimeAndCapacityArgs{
												MaxInstanceCount: pulumi.Int(2),
												MinInstanceCount: pulumi.Int(2),
												Time:             pulumi.String("09:00"),
											},
										},
										&hdinsight.AutoscaleScheduleArgs{
											Days: pulumi.StringArray{
												pulumi.String("Saturday"),
												pulumi.String("Sunday"),
											},
											TimeAndCapacity: &hdinsight.AutoscaleTimeAndCapacityArgs{
												MaxInstanceCount: pulumi.Int(4),
												MinInstanceCount: pulumi.Int(4),
												Time:             pulumi.String("18:00"),
											},
										},
									},
									TimeZone: pulumi.String("China Standard Time"),
								},
							},
							HardwareProfile: &hdinsight.HardwareProfileArgs{
								VmSize: pulumi.String("Standard_D4_V2"),
							},
							Name: pulumi.String("workernode"),
							OsProfile: &hdinsight.OsProfileArgs{
								LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
									Password: pulumi.String("**********"),
									Username: pulumi.String("sshuser"),
								},
							},
							ScriptActions:       hdinsight.ScriptActionArray{},
							TargetInstanceCount: pulumi.Int(4),
						},
					},
				},
				OsType: pulumi.String("Linux"),
				StorageProfile: &hdinsight.StorageProfileArgs{
					Storageaccounts: hdinsight.StorageAccountArray{
						&hdinsight.StorageAccountArgs{
							Container: pulumi.String("hdinsight-autoscale-tes-2019-06-18t05-49-16-591z"),
							IsDefault: pulumi.Bool(true),
							Key:       pulumi.String("storagekey"),
							Name:      pulumi.String("mystorage.blob.core.windows.net"),
						},
					},
				},
				Tier: pulumi.String("Standard"),
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
