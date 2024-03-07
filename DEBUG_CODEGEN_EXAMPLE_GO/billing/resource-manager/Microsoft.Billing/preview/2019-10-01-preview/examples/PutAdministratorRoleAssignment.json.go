package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/billing/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := billing.NewBillingRoleAssignmentByBillingAccount(ctx, "billingRoleAssignmentByBillingAccount", &billing.BillingRoleAssignmentByBillingAccountArgs{
			BillingAccountName:        pulumi.String("{billingAccountName}"),
			BillingRoleAssignmentName: pulumi.String("{billingRoleAssignmentName}"),
			PrincipalId:               pulumi.String("99a1a759-30dd-42c2-828c-db398826bb67"),
			PrincipalTenantId:         pulumi.String("7ca289b9-c32d-4f01-8566-7ff93261d76f"),
			RoleDefinitionId:          pulumi.String("/providers/Microsoft.Billing/billingAccounts/7898901/billingRoleDefinitions/9f1983cb-2574-400c-87e9-34cf8e2280db"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
