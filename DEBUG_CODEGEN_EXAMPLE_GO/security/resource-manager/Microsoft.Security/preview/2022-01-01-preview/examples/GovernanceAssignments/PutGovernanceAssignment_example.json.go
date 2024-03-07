package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewGovernanceAssignment(ctx, "governanceAssignment", &security.GovernanceAssignmentArgs{
			AdditionalData: &security.GovernanceAssignmentAdditionalDataArgs{
				TicketLink:   pulumi.String("https://snow.com"),
				TicketNumber: pulumi.Int(123123),
				TicketStatus: pulumi.String("Active"),
			},
			AssessmentName: pulumi.String("6b9421dd-5555-2251-9b3d-2be58e2f82cd"),
			AssignmentKey:  pulumi.String("6634ff9f-127b-4bf2-8e6e-b1737f5e789c"),
			GovernanceEmailNotification: &security.GovernanceEmailNotificationArgs{
				DisableManagerEmailNotification: pulumi.Bool(false),
				DisableOwnerEmailNotification:   pulumi.Bool(false),
			},
			IsGracePeriod:      pulumi.Bool(true),
			Owner:              pulumi.String("user@contoso.com"),
			RemediationDueDate: pulumi.String("2022-01-07T13:00:00.0000000Z"),
			RemediationEta: &security.RemediationEtaArgs{
				Eta:           pulumi.String("2022-01-08T13:00:00.0000000Z"),
				Justification: pulumi.String("Justification of ETA"),
			},
			Scope: pulumi.String("subscriptions/c32e05d9-7207-4e22-bdf4-4f7d9c72e5fd/resourceGroups/compute_servers/providers/Microsoft.Compute/virtualMachines/win2012"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
