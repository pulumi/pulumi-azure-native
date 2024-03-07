package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNspAssociation(ctx, "nspAssociation", &network.NspAssociationArgs{
			AccessMode:                   pulumi.String("Enforced"),
			AssociationName:              pulumi.String("association1"),
			NetworkSecurityPerimeterName: pulumi.String("nsp1"),
			PrivateLinkResource: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/{paasSubscriptionId}/resourceGroups/{paasResourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}"),
			},
			Profile: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/{profileName}"),
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
