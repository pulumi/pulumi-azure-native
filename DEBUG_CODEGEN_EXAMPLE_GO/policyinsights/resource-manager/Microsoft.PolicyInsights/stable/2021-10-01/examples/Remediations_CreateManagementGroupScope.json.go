package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/policyinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := policyinsights.NewRemediationAtManagementGroup(ctx, "remediationAtManagementGroup", &policyinsights.RemediationAtManagementGroupArgs{
			ManagementGroupId:         pulumi.String("financeMg"),
			ManagementGroupsNamespace: pulumi.String("Microsoft.Management"),
			PolicyAssignmentId:        pulumi.String("/providers/microsoft.management/managementGroups/financeMg/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
			RemediationName:           pulumi.String("storageRemediation"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
