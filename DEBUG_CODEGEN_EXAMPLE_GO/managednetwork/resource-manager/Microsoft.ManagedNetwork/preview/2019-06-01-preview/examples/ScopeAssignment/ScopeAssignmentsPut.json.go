package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetwork.NewScopeAssignment(ctx, "scopeAssignment", &managednetwork.ScopeAssignmentArgs{
			AssignedManagedNetwork: pulumi.String("/subscriptions/subscriptionA/resourceGroups/myResourceGroup/providers/Microsoft.ManagedNetwork/managedNetworks/myManagedNetwork"),
			Scope:                  pulumi.String("subscriptions/subscriptionC"),
			ScopeAssignmentName:    pulumi.String("subscriptionCAssignment"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
