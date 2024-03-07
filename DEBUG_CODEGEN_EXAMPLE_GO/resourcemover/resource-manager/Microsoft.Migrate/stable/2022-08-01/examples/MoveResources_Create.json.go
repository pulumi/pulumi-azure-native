package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewMoveResource(ctx, "moveResource", &migrate.MoveResourceArgs{
			MoveCollectionName: pulumi.String("movecollection1"),
			MoveResourceName:   pulumi.String("moveresourcename1"),
			Properties: migrate.MoveResourcePropertiesResponse{
				DependsOnOverrides: migrate.MoveResourceDependencyOverrideArray{
					&migrate.MoveResourceDependencyOverrideArgs{
						Id:       pulumi.String("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
						TargetId: pulumi.String("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/westusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
					},
				},
				ResourceSettings: migrate.VirtualMachineResourceSettings{
					ResourceType:            "Microsoft.Compute/virtualMachines",
					TargetAvailabilitySetId: "/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/availabilitySets/avset1",
					TargetAvailabilityZone:  "2",
					TargetResourceName:      "westusvm1",
					UserManagedIdentities: []string{
						"/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1",
					},
				},
				SourceId: pulumi.String("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/virtualMachines/eastusvm1"),
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
