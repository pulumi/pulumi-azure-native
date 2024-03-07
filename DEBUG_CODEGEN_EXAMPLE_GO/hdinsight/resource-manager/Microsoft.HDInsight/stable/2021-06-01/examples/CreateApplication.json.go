package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hdinsight/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hdinsight.NewApplication(ctx, "application", &hdinsight.ApplicationArgs{
			ApplicationName: pulumi.String("hue"),
			ClusterName:     pulumi.String("cluster1"),
			Properties: &hdinsight.ApplicationPropertiesArgs{
				ApplicationType: pulumi.String("CustomApplication"),
				ComputeProfile: &hdinsight.ComputeProfileArgs{
					Roles: hdinsight.RoleArray{
						&hdinsight.RoleArgs{
							HardwareProfile: &hdinsight.HardwareProfileArgs{
								VmSize: pulumi.String("Standard_D12_v2"),
							},
							Name:                pulumi.String("edgenode"),
							TargetInstanceCount: pulumi.Int(1),
						},
					},
				},
				Errors: hdinsight.ErrorsArray{},
				HttpsEndpoints: hdinsight.ApplicationGetHttpsEndpointArray{
					&hdinsight.ApplicationGetHttpsEndpointArgs{
						AccessModes: pulumi.StringArray{
							pulumi.String("WebPage"),
						},
						DestinationPort: pulumi.Int(20000),
						SubDomainSuffix: pulumi.String("dss"),
					},
				},
				InstallScriptActions: hdinsight.RuntimeScriptActionArray{
					&hdinsight.RuntimeScriptActionArgs{
						Name:       pulumi.String("app-install-app1"),
						Parameters: pulumi.String("-version latest -port 20000"),
						Roles: pulumi.StringArray{
							pulumi.String("edgenode"),
						},
						Uri: pulumi.String("https://.../install.sh"),
					},
				},
				UninstallScriptActions: hdinsight.RuntimeScriptActionArray{},
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
