package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabric.NewManagedClusterService(ctx, "managedClusterService", &servicefabric.ManagedClusterServiceArgs{
			ApplicationName: pulumi.String("myApp"),
			ClusterName:     pulumi.String("myCluster"),
			Location:        pulumi.String("eastus"),
			Properties: servicefabric.StatelessServiceProperties{
				CorrelationScheme: []servicefabric.ServiceCorrelation{
					{
						Scheme:      "AlignedAffinity",
						ServiceName: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applications/myApp/services/myService1",
					},
				},
				DefaultMoveCost:       "Medium",
				InstanceCount:         5,
				MinInstanceCount:      3,
				MinInstancePercentage: 30,
				PartitionDescription: servicefabric.SingletonPartitionScheme{
					PartitionScheme: "Singleton",
				},
				PlacementConstraints: "NodeType==frontend",
				ScalingPolicies: []servicefabric.ScalingPolicy{
					{
						ScalingMechanism: {
							Kind:             "ScalePartitionInstanceCount",
							MaxInstanceCount: 9,
							MinInstanceCount: 3,
							ScaleIncrement:   2,
						},
						ScalingTrigger: {
							Kind:               "AveragePartitionLoadTrigger",
							LowerLoadThreshold: 2,
							MetricName:         "metricName",
							ScaleInterval:      "00:01:00",
							UpperLoadThreshold: 8,
						},
					},
				},
				ServiceDnsName: "myservicednsname.myApp",
				ServiceKind:    "Stateless",
				ServiceLoadMetrics: []servicefabric.ServiceLoadMetric{
					{
						DefaultLoad: 3,
						Name:        "metric1",
						Weight:      "Low",
					},
				},
				ServicePackageActivationMode: "SharedProcess",
				ServicePlacementPolicies: []interface{}{
					servicefabric.ServicePlacementNonPartiallyPlaceServicePolicy{
						Type: "NonPartiallyPlaceService",
					},
				},
				ServiceTypeName: "myServiceType",
			},
			ResourceGroupName: pulumi.String("resRg"),
			ServiceName:       pulumi.String("myService"),
			Tags: pulumi.StringMap{
				"a": pulumi.String("b"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
