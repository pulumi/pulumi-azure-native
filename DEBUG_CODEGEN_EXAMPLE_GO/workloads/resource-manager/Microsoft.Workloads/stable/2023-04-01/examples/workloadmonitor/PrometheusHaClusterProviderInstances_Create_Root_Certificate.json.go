package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewProviderInstance(ctx, "providerInstance", &workloads.ProviderInstanceArgs{
			MonitorName:          pulumi.String("mySapMonitor"),
			ProviderInstanceName: pulumi.String("myProviderInstance"),
			ProviderSettings: workloads.PrometheusHaClusterProviderInstanceProperties{
				ClusterName:   "clusterName",
				Hostname:      "hostname",
				PrometheusUrl: "http://192.168.0.0:9090/metrics",
				ProviderType:  "PrometheusHaCluster",
				Sid:           "sid",
				SslPreference: "RootCertificate",
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
