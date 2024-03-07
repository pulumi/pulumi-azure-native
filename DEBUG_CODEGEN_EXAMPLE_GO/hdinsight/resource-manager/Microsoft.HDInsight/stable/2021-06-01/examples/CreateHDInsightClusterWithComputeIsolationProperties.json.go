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
							"restAuthCredential.isEnabled": true,
							"restAuthCredential.password":  "**********",
							"restAuthCredential.username":  "admin",
						},
					},
					Kind: pulumi.String("hadoop"),
				},
				ClusterVersion: pulumi.String("3.6"),
				ComputeIsolationProperties: &hdinsight.ComputeIsolationPropertiesArgs{
					EnableComputeIsolation: pulumi.Bool(true),
				},
				ComputeProfile: &hdinsight.ComputeProfileArgs{
					Roles: hdinsight.RoleArray{
						&hdinsight.RoleArgs{
							HardwareProfile: &hdinsight.HardwareProfileArgs{
								VmSize: pulumi.String("standard_d3"),
							},
							Name: pulumi.String("headnode"),
							OsProfile: &hdinsight.OsProfileArgs{
								LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
									Password: pulumi.String("**********"),
									SshProfile: &hdinsight.SshProfileArgs{
										PublicKeys: hdinsight.SshPublicKeyArray{
											&hdinsight.SshPublicKeyArgs{
												CertificateData: pulumi.String("**********"),
											},
										},
									},
									Username: pulumi.String("sshuser"),
								},
							},
							TargetInstanceCount: pulumi.Int(2),
						},
						&hdinsight.RoleArgs{
							HardwareProfile: &hdinsight.HardwareProfileArgs{
								VmSize: pulumi.String("standard_d3"),
							},
							Name: pulumi.String("workernode"),
							OsProfile: &hdinsight.OsProfileArgs{
								LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
									Password: pulumi.String("**********"),
									SshProfile: &hdinsight.SshProfileArgs{
										PublicKeys: hdinsight.SshPublicKeyArray{
											&hdinsight.SshPublicKeyArgs{
												CertificateData: pulumi.String("**********"),
											},
										},
									},
									Username: pulumi.String("sshuser"),
								},
							},
							TargetInstanceCount: pulumi.Int(2),
						},
					},
				},
				OsType: pulumi.String("Linux"),
				StorageProfile: &hdinsight.StorageProfileArgs{
					Storageaccounts: hdinsight.StorageAccountArray{
						&hdinsight.StorageAccountArgs{
							Container: pulumi.String("containername"),
							IsDefault: pulumi.Bool(true),
							Key:       pulumi.String("storage account key"),
							Name:      pulumi.String("mystorage"),
						},
					},
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
