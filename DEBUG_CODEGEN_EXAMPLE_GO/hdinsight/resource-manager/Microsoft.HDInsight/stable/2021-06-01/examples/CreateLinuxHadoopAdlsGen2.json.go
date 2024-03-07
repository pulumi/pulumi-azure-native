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
					Configurations: pulumi.Any{
						Gateway: map[string]interface{}{
							"restAuthCredential.isEnabled": "true",
							"restAuthCredential.password":  "**********",
							"restAuthCredential.username":  "admin",
						},
					},
					Kind: pulumi.String("Hadoop"),
				},
				ClusterVersion: pulumi.String("3.6"),
				ComputeProfile: &hdinsight.ComputeProfileArgs{
					Roles: hdinsight.RoleArray{
						&hdinsight.RoleArgs{
							HardwareProfile: &hdinsight.HardwareProfileArgs{
								VmSize: pulumi.String("Standard_D3_V2"),
							},
							MinInstanceCount: pulumi.Int(1),
							Name:             pulumi.String("headnode"),
							OsProfile: &hdinsight.OsProfileArgs{
								LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
									Password: pulumi.String("**********"),
									Username: pulumi.String("sshuser"),
								},
							},
							TargetInstanceCount: pulumi.Int(2),
						},
						&hdinsight.RoleArgs{
							HardwareProfile: &hdinsight.HardwareProfileArgs{
								VmSize: pulumi.String("Standard_D3_V2"),
							},
							MinInstanceCount: pulumi.Int(1),
							Name:             pulumi.String("workernode"),
							OsProfile: &hdinsight.OsProfileArgs{
								LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
									Password: pulumi.String("**********"),
									Username: pulumi.String("sshuser"),
								},
							},
							TargetInstanceCount: pulumi.Int(4),
						},
						&hdinsight.RoleArgs{
							HardwareProfile: &hdinsight.HardwareProfileArgs{
								VmSize: pulumi.String("Small"),
							},
							MinInstanceCount: pulumi.Int(1),
							Name:             pulumi.String("zookeepernode"),
							OsProfile: &hdinsight.OsProfileArgs{
								LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
									Password: pulumi.String("**********"),
									Username: pulumi.String("sshuser"),
								},
							},
							TargetInstanceCount: pulumi.Int(3),
						},
					},
				},
				OsType: pulumi.String("Linux"),
				StorageProfile: &hdinsight.StorageProfileArgs{
					Storageaccounts: hdinsight.StorageAccountArray{
						&hdinsight.StorageAccountArgs{
							FileSystem: pulumi.String("default"),
							IsDefault:  pulumi.Bool(true),
							Key:        pulumi.String("storagekey"),
							Name:       pulumi.String("mystorage.dfs.core.windows.net"),
						},
					},
				},
				Tier: pulumi.String("Standard"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("val1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
