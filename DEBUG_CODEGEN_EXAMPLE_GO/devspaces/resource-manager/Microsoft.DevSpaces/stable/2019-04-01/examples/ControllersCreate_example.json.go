package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devspaces/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devspaces.NewController(ctx, "controller", &devspaces.ControllerArgs{
			Location:          pulumi.String("eastus"),
			Name:              pulumi.String("myControllerResource"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &devspaces.SkuArgs{
				Name: pulumi.String("S1"),
				Tier: pulumi.String("Standard"),
			},
			Tags:                                 nil,
			TargetContainerHostCredentialsBase64: pulumi.String("QmFzZTY0IEVuY29kZWQgVmFsdWUK"),
			TargetContainerHostResourceId:        pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
