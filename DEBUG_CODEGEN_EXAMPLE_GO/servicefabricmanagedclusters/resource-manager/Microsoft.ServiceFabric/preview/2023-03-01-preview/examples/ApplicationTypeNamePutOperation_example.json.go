package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabric.NewManagedClusterApplicationType(ctx, "managedClusterApplicationType", &servicefabric.ManagedClusterApplicationTypeArgs{
			ApplicationTypeName: pulumi.String("myAppType"),
			ClusterName:         pulumi.String("myCluster"),
			Location:            pulumi.String("eastus"),
			ResourceGroupName:   pulumi.String("resRg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
