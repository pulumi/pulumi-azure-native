package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kubernetesconfiguration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kubernetesconfiguration.NewExtension(ctx, "extension", &kubernetesconfiguration.ExtensionArgs{
			AutoUpgradeMinorVersion: pulumi.Bool(true),
			ClusterName:             pulumi.String("clusterName1"),
			ClusterResourceName:     pulumi.String("connectedClusters"),
			ClusterRp:               pulumi.String("Microsoft.Kubernetes"),
			ConfigurationProtectedSettings: pulumi.StringMap{
				"omsagent.secret.key": pulumi.String("secretKeyValue01"),
			},
			ConfigurationSettings: pulumi.StringMap{
				"omsagent.env.clusterName": pulumi.String("clusterName1"),
				"omsagent.secret.wsid":     pulumi.String("a38cef99-5a89-52ed-b6db-22095c23664b"),
			},
			ExtensionName:     pulumi.String("ClusterMonitor"),
			ExtensionType:     pulumi.String("azuremonitor-containers"),
			ReleaseTrain:      pulumi.String("Preview"),
			ResourceGroupName: pulumi.String("rg1"),
			Scope: kubernetesconfiguration.ScopeResponse{
				Cluster: &kubernetesconfiguration.ScopeClusterArgs{
					ReleaseNamespace: pulumi.String("kube-system"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
