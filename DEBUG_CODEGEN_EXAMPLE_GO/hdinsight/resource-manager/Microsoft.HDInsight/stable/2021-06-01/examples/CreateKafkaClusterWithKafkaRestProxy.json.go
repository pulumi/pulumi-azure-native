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
						"Kafka": pulumi.String("2.1"),
					},
					Configurations: pulumi.Any{
						Gateway: map[string]interface{}{
							"restAuthCredential.isEnabled": true,
							"restAuthCredential.password":  "**********",
							"restAuthCredential.username":  "admin",
						},
					},
					Kind: pulumi.String("kafka"),
				},
				ClusterVersion: pulumi.String("4.0"),
				ComputeProfile: &hdinsight.ComputeProfileArgs{
					Roles: hdinsight.RoleArray{
						&hdinsight.RoleArgs{
							HardwareProfile: &hdinsight.HardwareProfileArgs{
								VmSize: pulumi.String("Large"),
							},
							Name: pulumi.String("headnode"),
							OsProfile: &hdinsight.OsProfileArgs{
								LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
									Password: pulumi.String("**********"),
									Username: pulumi.String("sshuser"),
								},
							},
							TargetInstanceCount: pulumi.Int(2),
						},
						&hdinsight.RoleArgs{
							DataDisksGroups: hdinsight.DataDisksGroupsArray{
								&hdinsight.DataDisksGroupsArgs{
									DisksPerNode: pulumi.Int(8),
								},
							},
							HardwareProfile: &hdinsight.HardwareProfileArgs{
								VmSize: pulumi.String("Large"),
							},
							Name: pulumi.String("workernode"),
							OsProfile: &hdinsight.OsProfileArgs{
								LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
									Password: pulumi.String("**********"),
									Username: pulumi.String("sshuser"),
								},
							},
							TargetInstanceCount: pulumi.Int(3),
						},
						&hdinsight.RoleArgs{
							HardwareProfile: &hdinsight.HardwareProfileArgs{
								VmSize: pulumi.String("Small"),
							},
							Name: pulumi.String("zookeepernode"),
							OsProfile: &hdinsight.OsProfileArgs{
								LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
									Password: pulumi.String("**********"),
									Username: pulumi.String("sshuser"),
								},
							},
							TargetInstanceCount: pulumi.Int(3),
						},
						&hdinsight.RoleArgs{
							HardwareProfile: &hdinsight.HardwareProfileArgs{
								VmSize: pulumi.String("Standard_D4_v2"),
							},
							Name: pulumi.String("kafkamanagementnode"),
							OsProfile: &hdinsight.OsProfileArgs{
								LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
									Password: pulumi.String("**********"),
									Username: pulumi.String("kafkauser"),
								},
							},
							TargetInstanceCount: pulumi.Int(2),
						},
					},
				},
				KafkaRestProperties: &hdinsight.KafkaRestPropertiesArgs{
					ClientGroupInfo: &hdinsight.ClientGroupInfoArgs{
						GroupId:   pulumi.String("00000000-0000-0000-0000-111111111111"),
						GroupName: pulumi.String("Kafka security group name"),
					},
				},
				OsType: pulumi.String("Linux"),
				StorageProfile: &hdinsight.StorageProfileArgs{
					Storageaccounts: hdinsight.StorageAccountArray{
						&hdinsight.StorageAccountArgs{
							Container: pulumi.String("containername"),
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
