package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningcompute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningcompute.NewOperationalizationCluster(ctx, "operationalizationCluster", &machinelearningcompute.OperationalizationClusterArgs{
			ClusterName: pulumi.String("myCluster"),
			ClusterType: pulumi.String("ACS"),
			ContainerService: &machinelearningcompute.AcsClusterPropertiesArgs{
				OrchestratorProperties: &machinelearningcompute.KubernetesClusterPropertiesArgs{
					ServicePrincipal: &machinelearningcompute.ServicePrincipalPropertiesArgs{
						ClientId: pulumi.String("abcdefghijklmnopqrt"),
						Secret:   pulumi.String("<secret>"),
					},
				},
				OrchestratorType: pulumi.String("Kubernetes"),
			},
			Description: pulumi.String("My Operationalization Cluster"),
			GlobalServiceConfiguration: &machinelearningcompute.GlobalServiceConfigurationArgs{
				Ssl: &machinelearningcompute.SslConfigurationArgs{
					Cert:   pulumi.String("afjdklq2131casfakld="),
					Cname:  pulumi.String("foo.bar.com"),
					Key:    pulumi.String("flksdafkldsajf="),
					Status: pulumi.String("Enabled"),
				},
			},
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("alpha"),
				"key2": pulumi.String("beta"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
