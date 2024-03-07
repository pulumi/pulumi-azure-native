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
				InstanceCount: 1,
				PartitionDescription: servicefabric.SingletonPartitionScheme{
					PartitionScheme: "Singleton",
				},
				ServiceKind:     "Stateless",
				ServiceTypeName: "myServiceType",
			},
			ResourceGroupName: pulumi.String("resRg"),
			ServiceName:       pulumi.String("myService"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
