package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabric.NewManagedClusterApplication(ctx, "managedClusterApplication", &servicefabric.ManagedClusterApplicationArgs{
			ApplicationName:   pulumi.String("myApp"),
			ClusterName:       pulumi.String("myCluster"),
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("resRg"),
			Version:           pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType/versions/1.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
