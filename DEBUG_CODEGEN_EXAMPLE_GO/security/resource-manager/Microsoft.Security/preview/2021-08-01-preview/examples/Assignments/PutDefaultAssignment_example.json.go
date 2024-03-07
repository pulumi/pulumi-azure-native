package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewAssignment(ctx, "assignment", &security.AssignmentArgs{
			AssignedStandard: &security.AssignedStandardItemArgs{
				Id: pulumi.String("/providers/Microsoft.Security/Standards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
			},
			AssignmentId:      pulumi.String("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
			Description:       pulumi.String("Set of policies monitored by Azure Security Center for cross cloud"),
			DisplayName:       pulumi.String("ASC Default"),
			Effect:            pulumi.String("audit"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Scope:             pulumi.String("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/ResourceGroup/rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
