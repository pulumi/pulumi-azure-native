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
			ExtensionName:           pulumi.String("azureVote"),
			ExtensionType:           pulumi.String("azure-vote"),
			Plan: &kubernetesconfiguration.PlanArgs{
				Name:      pulumi.String("azure-vote-standard"),
				Product:   pulumi.String("azure-vote-standard-offer-id"),
				Publisher: pulumi.String("Microsoft"),
			},
			ReleaseTrain:      pulumi.String("Preview"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
