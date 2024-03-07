package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewReplicationRecoveryPlan(ctx, "replicationRecoveryPlan", &recoveryservices.ReplicationRecoveryPlanArgs{
			Properties: recoveryservices.RecoveryPlanPropertiesResponse{
				FailoverDeploymentModel: pulumi.String("ResourceManager"),
				Groups: []recoveryservices.RecoveryPlanGroupArgs{
					{
						EndGroupActions: recoveryservices.RecoveryPlanActionArray{},
						GroupType:       pulumi.String("Boot"),
						ReplicationProtectedItems: recoveryservices.RecoveryPlanProtectedItemArray{
							{
								Id:               pulumi.String("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
								VirtualMachineId: pulumi.String("f8491e4f-817a-40dd-a90c-af773978c75b"),
							},
						},
						StartGroupActions: recoveryservices.RecoveryPlanActionArray{},
					},
				},
				PrimaryFabricId:  pulumi.String("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
				RecoveryFabricId: pulumi.String("Microsoft Azure"),
			},
			RecoveryPlanName:  pulumi.String("RPtest1"),
			ResourceGroupName: pulumi.String("resourceGroupPS1"),
			ResourceName:      pulumi.String("vault1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
