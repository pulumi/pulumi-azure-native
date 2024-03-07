package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/netapp/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netapp.NewVolumeGroup(ctx, "volumeGroup", &netapp.VolumeGroupArgs{
			AccountName: pulumi.String("account1"),
			GroupMetaData: &netapp.VolumeGroupMetaDataArgs{
				ApplicationIdentifier: pulumi.String("DEV"),
				ApplicationType:       pulumi.String("SAP-HANA"),
				DeploymentSpecId:      pulumi.String("20542149-bfca-5618-1879-9863dc6767f1"),
				GroupDescription:      pulumi.String("Volume group"),
			},
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("myRG"),
			VolumeGroupName:   pulumi.String("group1"),
			Volumes: netapp.VolumeGroupVolumePropertiesArray{
				&netapp.VolumeGroupVolumePropertiesArgs{
					CapacityPoolResourceId:  pulumi.String("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
					CreationToken:           pulumi.String("test-data-mnt00001"),
					Name:                    pulumi.String("test-data-mnt00001"),
					ProximityPlacementGroup: pulumi.String("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
					ServiceLevel:            pulumi.String("Premium"),
					SubnetId:                pulumi.String("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
					ThroughputMibps:         pulumi.Float64(10),
					UsageThreshold:          pulumi.Float64(107374182400),
					VolumeSpecName:          pulumi.String("data"),
				},
				&netapp.VolumeGroupVolumePropertiesArgs{
					CapacityPoolResourceId:  pulumi.String("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
					CreationToken:           pulumi.String("test-log-mnt00001"),
					Name:                    pulumi.String("test-log-mnt00001"),
					ProximityPlacementGroup: pulumi.String("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
					ServiceLevel:            pulumi.String("Premium"),
					SubnetId:                pulumi.String("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
					ThroughputMibps:         pulumi.Float64(10),
					UsageThreshold:          pulumi.Float64(107374182400),
					VolumeSpecName:          pulumi.String("log"),
				},
				&netapp.VolumeGroupVolumePropertiesArgs{
					CapacityPoolResourceId:  pulumi.String("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
					CreationToken:           pulumi.String("test-shared"),
					Name:                    pulumi.String("test-shared"),
					ProximityPlacementGroup: pulumi.String("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
					ServiceLevel:            pulumi.String("Premium"),
					SubnetId:                pulumi.String("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
					ThroughputMibps:         pulumi.Float64(10),
					UsageThreshold:          pulumi.Float64(107374182400),
					VolumeSpecName:          pulumi.String("shared"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
