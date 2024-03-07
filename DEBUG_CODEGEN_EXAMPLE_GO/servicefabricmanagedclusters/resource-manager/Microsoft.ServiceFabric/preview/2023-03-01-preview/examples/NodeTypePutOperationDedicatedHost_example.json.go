package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabric.NewNodeType(ctx, "nodeType", &servicefabric.NodeTypeArgs{
			Capacities:          nil,
			ClusterName:         pulumi.String("myCluster"),
			DataDiskSizeGB:      pulumi.Int(200),
			DataDiskType:        pulumi.String("StandardSSD_LRS"),
			HostGroupId:         pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testhostgroupRG/providers/Microsoft.Compute/hostGroups/testHostGroup"),
			IsPrimary:           pulumi.Bool(false),
			NodeTypeName:        pulumi.String("BE"),
			PlacementProperties: nil,
			ResourceGroupName:   pulumi.String("resRg"),
			VmImageOffer:        pulumi.String("WindowsServer"),
			VmImagePublisher:    pulumi.String("MicrosoftWindowsServer"),
			VmImageSku:          pulumi.String("2019-Datacenter"),
			VmImageVersion:      pulumi.String("latest"),
			VmInstanceCount:     pulumi.Int(10),
			VmSize:              pulumi.String("Standard_D8s_v3"),
			Zones: pulumi.StringArray{
				pulumi.String("1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
