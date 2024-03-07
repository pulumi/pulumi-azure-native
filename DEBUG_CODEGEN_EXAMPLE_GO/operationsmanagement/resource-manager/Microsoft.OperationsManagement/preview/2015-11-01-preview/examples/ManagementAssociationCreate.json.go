package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationsmanagement.NewManagementAssociation(ctx, "managementAssociation", &operationsmanagement.ManagementAssociationArgs{
			Location:                  pulumi.String("East US"),
			ManagementAssociationName: pulumi.String("managementAssociation1"),
			Properties: &operationsmanagement.ManagementAssociationPropertiesArgs{
				ApplicationId: pulumi.String("/subscriptions/sub1/resourcegroups/rg1/providers/Microsoft.Appliance/Appliances/appliance1"),
			},
			ProviderName:      pulumi.String("providerName"),
			ResourceGroupName: pulumi.String("rg1"),
			ResourceName:      pulumi.String("resourceName"),
			ResourceType:      pulumi.String("resourceType"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
