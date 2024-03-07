package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewAssignment(ctx, "assignment", &security.AssignmentArgs{
			AdditionalData: &security.AssignmentPropertiesAdditionalDataArgs{
				ExemptionCategory: pulumi.String("waiver"),
			},
			AssignedComponent: &security.AssignedComponentItemArgs{
				Key: pulumi.String("1195afff-c881-495e-9bc5-1486211ae03f"),
			},
			AssignedStandard: &security.AssignedStandardItemArgs{
				Id: pulumi.String("/providers/Microsoft.Security/Standards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
			},
			AssignmentId: pulumi.String("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
			Description:  pulumi.String("Set of policies monitored by Azure Security Center for cross cloud"),
			DisplayName:  pulumi.String("ASC Default"),
			Effect:       pulumi.String("Exempt"),
			ExpiresOn:    pulumi.String("2022-05-01T19:50:47.083633Z"),
			Metadata: pulumi.Any{
				TicketId: 12345,
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Scope:             pulumi.String("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/ResourceGroup/rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
