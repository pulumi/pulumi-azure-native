package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcompute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcompute.NewMachine(ctx, "machine", &hybridcompute.MachineArgs{
			ClientPublicKey: pulumi.String("string"),
			Identity: &hybridcompute.IdentityArgs{
				Type: hybridcompute.ResourceIdentityTypeSystemAssigned,
			},
			Location: pulumi.String("eastus2euap"),
			LocationData: &hybridcompute.LocationDataArgs{
				Name: pulumi.String("Redmond"),
			},
			MachineName:                pulumi.String("myMachine"),
			ParentClusterResourceId:    pulumi.String("{AzureStackHCIResourceId}"),
			PrivateLinkScopeResourceId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
			ResourceGroupName:          pulumi.String("myResourceGroup"),
			VmId:                       pulumi.String("b7a098cc-b0b8-46e8-a205-62f301a62a8f"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
