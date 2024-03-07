package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/billing/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := billing.NewBillingRoleAssignmentByEnrollmentAccount(ctx, "billingRoleAssignmentByEnrollmentAccount", &billing.BillingRoleAssignmentByEnrollmentAccountArgs{
			BillingAccountName:        pulumi.String("{billingAccountName}"),
			BillingRoleAssignmentName: pulumi.String("{billingRoleAssignmentName}"),
			EnrollmentAccountName:     pulumi.String("{enrollmentAccountName}"),
			PrincipalId:               pulumi.String("99a1a759-30dd-42c2-828c-db398826bb67"),
			PrincipalTenantId:         pulumi.String("7ca289b9-c32d-4f01-8566-7ff93261d76f"),
			RoleDefinitionId:          pulumi.String("/providers/Microsoft.Billing/billingAccounts/7898901/enrollmentAccounts/225314/billingRoleDefinitions/a0bcee42-bf30-4d1b-926a-48d21664ef71"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
