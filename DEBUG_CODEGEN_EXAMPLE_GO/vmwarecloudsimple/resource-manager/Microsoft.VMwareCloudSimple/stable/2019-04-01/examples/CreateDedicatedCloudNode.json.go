package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/vmwarecloudsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vmwarecloudsimple.NewDedicatedCloudNode(ctx, "dedicatedCloudNode", &vmwarecloudsimple.DedicatedCloudNodeArgs{
			AvailabilityZoneId:     pulumi.String("az1"),
			DedicatedCloudNodeName: pulumi.String("myNode"),
			Id:                     pulumi.String("general"),
			Location:               pulumi.String("westus"),
			Name:                   pulumi.String("CS28-Node"),
			NodesCount:             pulumi.Int(1),
			PlacementGroupId:       pulumi.String("n1"),
			PurchaseId:             pulumi.String("56acbd46-3d36-4bbf-9b08-57c30fdf6932"),
			ResourceGroupName:      pulumi.String("myResourceGroup"),
			Sku: &vmwarecloudsimple.SkuArgs{
				Name: pulumi.String("VMware_CloudSimple_CS28"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
