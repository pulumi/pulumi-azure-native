package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewPlacementPolicy(ctx, "placementPolicy", &avs.PlacementPolicyArgs{
			ClusterName:         pulumi.String("cluster1"),
			PlacementPolicyName: pulumi.String("policy1"),
			PrivateCloudName:    pulumi.String("cloud1"),
			Properties: avs.VmHostPlacementPolicyProperties{
				AffinityStrength:       "Must",
				AffinityType:           "AntiAffinity",
				AzureHybridBenefitType: "SqlHost",
				HostMembers: []string{
					"fakehost22.nyc1.kubernetes.center",
					"fakehost23.nyc1.kubernetes.center",
					"fakehost24.nyc1.kubernetes.center",
				},
				Type: "VmHost",
				VmMembers: []string{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128",
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-256",
				},
			},
			ResourceGroupName: pulumi.String("group1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
