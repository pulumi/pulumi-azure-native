package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewPrivateLinkAssociation(ctx, "privateLinkAssociation", &authorization.PrivateLinkAssociationArgs{
			GroupId: pulumi.String("my-management-group"),
			PlaId:   pulumi.String("00000000-0000-0000-0000-000000000000"),
			Properties: &authorization.PrivateLinkAssociationPropertiesArgs{
				PrivateLink:         pulumi.String("00000000-0000-0000-0000-000000000000"),
				PublicNetworkAccess: pulumi.String("Enabled"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
