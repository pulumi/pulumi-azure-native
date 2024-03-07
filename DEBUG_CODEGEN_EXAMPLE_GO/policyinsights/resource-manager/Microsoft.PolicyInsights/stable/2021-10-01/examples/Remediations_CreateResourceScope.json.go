package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/policyinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := policyinsights.NewRemediationAtResource(ctx, "remediationAtResource", &policyinsights.RemediationAtResourceArgs{
			PolicyAssignmentId: pulumi.String("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/myResourceGroup/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
			RemediationName:    pulumi.String("storageRemediation"),
			ResourceId:         pulumi.String("subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storAc1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
