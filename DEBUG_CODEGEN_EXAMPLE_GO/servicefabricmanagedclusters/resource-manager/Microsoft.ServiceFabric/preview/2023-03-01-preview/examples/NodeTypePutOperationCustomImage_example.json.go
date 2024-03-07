package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabric.NewNodeType(ctx, "nodeType", &servicefabric.NodeTypeArgs{
			ClusterName:       pulumi.String("myCluster"),
			DataDiskSizeGB:    pulumi.Int(200),
			IsPrimary:         pulumi.Bool(false),
			NodeTypeName:      pulumi.String("BE"),
			ResourceGroupName: pulumi.String("resRg"),
			VmImageResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-custom-image/providers/Microsoft.Compute/galleries/myCustomImages/images/Win2019DC"),
			VmInstanceCount:   pulumi.Int(10),
			VmSize:            pulumi.String("Standard_D3"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
