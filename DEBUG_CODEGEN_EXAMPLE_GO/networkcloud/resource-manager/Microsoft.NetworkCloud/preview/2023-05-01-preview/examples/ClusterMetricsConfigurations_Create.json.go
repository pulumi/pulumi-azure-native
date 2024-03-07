package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkcloud.NewMetricsConfiguration(ctx, "metricsConfiguration", &networkcloud.MetricsConfigurationArgs{
			ClusterName:        pulumi.String("clusterName"),
			CollectionInterval: pulumi.Float64(15),
			EnabledMetrics: pulumi.StringArray{
				pulumi.String("metric1"),
				pulumi.String("metric2"),
			},
			ExtendedLocation: &networkcloud.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
				Type: pulumi.String("CustomLocation"),
			},
			Location:                 pulumi.String("location"),
			MetricsConfigurationName: pulumi.String("default"),
			ResourceGroupName:        pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("myvalue1"),
				"key2": pulumi.String("myvalue2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
