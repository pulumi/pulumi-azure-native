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
					Kind: pulumi.String("Hadoop"),
				},
				ClusterVersion: pulumi.String("3.5"),
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
							ScriptActions:       hdinsight.ScriptActionArray{},
							TargetInstanceCount: pulumi.Int(2),
							VirtualNetworkProfile: &hdinsight.VirtualNetworkProfileArgs{
								Id:     pulumi.String("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname"),
								Subnet: pulumi.String("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
							},
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
							ScriptActions:       hdinsight.ScriptActionArray{},
							TargetInstanceCount: pulumi.Int(4),
							VirtualNetworkProfile: &hdinsight.VirtualNetworkProfileArgs{
								Id:     pulumi.String("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname"),
								Subnet: pulumi.String("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
							},
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
							ScriptActions:       hdinsight.ScriptActionArray{},
							TargetInstanceCount: pulumi.Int(3),
							VirtualNetworkProfile: &hdinsight.VirtualNetworkProfileArgs{
								Id:     pulumi.String("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname"),
								Subnet: pulumi.String("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
							},
						},
					},
				},
				OsType: pulumi.String("Linux"),
				SecurityProfile: &hdinsight.SecurityProfileArgs{
					ClusterUsersGroupDNs: pulumi.StringArray{
						pulumi.String("hdiusers"),
					},
					DirectoryType:      pulumi.String("ActiveDirectory"),
					Domain:             pulumi.String("DomainName"),
					DomainUserPassword: pulumi.String("**********"),
					DomainUsername:     pulumi.String("DomainUsername"),
					LdapsUrls: pulumi.StringArray{
						pulumi.String("ldaps://10.10.0.4:636"),
					},
					OrganizationalUnitDN: pulumi.String("OU=Hadoop,DC=hdinsight,DC=test"),
				},
				StorageProfile: &hdinsight.StorageProfileArgs{
					Storageaccounts: hdinsight.StorageAccountArray{
						&hdinsight.StorageAccountArgs{
							Container: pulumi.String("containername"),
							IsDefault: pulumi.Bool(true),
							Key:       pulumi.String("storage account key"),
							Name:      pulumi.String("mystorage.blob.core.windows.net"),
						},
					},
				},
				Tier: pulumi.String("Premium"),
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
