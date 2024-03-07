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
						AmbariConf: map[string]interface{}{
							"database-name":          "{ambari database name}",
							"database-server":        "{sql server name}.database.windows.net",
							"database-user-name":     "**********",
							"database-user-password": "**********",
						},
						Gateway: map[string]interface{}{
							"restAuthCredential.isEnabled": true,
							"restAuthCredential.password":  "**********",
							"restAuthCredential.username":  "admin",
						},
						HiveEnv: map[string]interface{}{
							"hive_database":                       "Existing MSSQL Server database with SQL authentication",
							"hive_database_name":                  "{hive metastore name}",
							"hive_database_type":                  "mssql",
							"hive_existing_mssql_server_database": "{hive metastore name}",
							"hive_existing_mssql_server_host":     "{sql server name}.database.windows.net",
							"hive_hostname":                       "{sql server name}.database.windows.net",
						},
						HiveSite: map[string]interface{}{
							"javax.jdo.option.ConnectionDriverName": "com.microsoft.sqlserver.jdbc.SQLServerDriver",
							"javax.jdo.option.ConnectionPassword":   "**********!",
							"javax.jdo.option.ConnectionURL":        "jdbc:sqlserver://{sql server name}.database.windows.net;database={hive metastore name};encrypt=true;trustServerCertificate=true;create=false;loginTimeout=300;sendStringParametersAsUnicode=true;prepareSQL=0",
							"javax.jdo.option.ConnectionUserName":   "**********",
						},
						OozieEnv: map[string]interface{}{
							"oozie_database":                       "Existing MSSQL Server database with SQL authentication",
							"oozie_database_name":                  "{oozie metastore name}",
							"oozie_database_type":                  "mssql",
							"oozie_existing_mssql_server_database": "{oozie metastore name}",
							"oozie_existing_mssql_server_host":     "{sql server name}.database.windows.net",
							"oozie_hostname":                       "{sql server name}.database.windows.net",
						},
						OozieSite: map[string]interface{}{
							"oozie.db.schema.name":                   "oozie",
							"oozie.service.JPAService.jdbc.driver":   "com.microsoft.sqlserver.jdbc.SQLServerDriver",
							"oozie.service.JPAService.jdbc.password": "**********",
							"oozie.service.JPAService.jdbc.url":      "jdbc:sqlserver://{sql server name}.database.windows.net;database={oozie metastore name};encrypt=true;trustServerCertificate=true;create=false;loginTimeout=300;sendStringParametersAsUnicode=true;prepareSQL=0",
							"oozie.service.JPAService.jdbc.username": "**********",
						},
					},
					Kind: pulumi.String("hadoop"),
				},
				ClusterVersion: pulumi.String("3.6"),
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
							VirtualNetworkProfile: &hdinsight.VirtualNetworkProfileArgs{
								Id:     pulumi.String("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname"),
								Subnet: pulumi.String("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
							},
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
							VirtualNetworkProfile: &hdinsight.VirtualNetworkProfileArgs{
								Id:     pulumi.String("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname"),
								Subnet: pulumi.String("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"),
							},
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
			Zones: pulumi.StringArray{
				pulumi.String("1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
