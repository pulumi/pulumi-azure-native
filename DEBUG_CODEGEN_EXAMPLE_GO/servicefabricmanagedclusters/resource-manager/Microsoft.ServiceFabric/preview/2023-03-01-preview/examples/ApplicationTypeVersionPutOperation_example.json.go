package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabric.NewManagedClusterApplicationTypeVersion(ctx, "managedClusterApplicationTypeVersion", &servicefabric.ManagedClusterApplicationTypeVersionArgs{
			AppPackageUrl:       pulumi.String("http://fakelink.test.com/MyAppType"),
			ApplicationTypeName: pulumi.String("myAppType"),
			ClusterName:         pulumi.String("myCluster"),
			Location:            pulumi.String("eastus"),
			ResourceGroupName:   pulumi.String("resRg"),
			Version:             pulumi.String("1.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
