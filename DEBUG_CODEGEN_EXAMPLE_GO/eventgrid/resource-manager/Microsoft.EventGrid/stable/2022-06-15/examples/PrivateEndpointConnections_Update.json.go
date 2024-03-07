package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &eventgrid.PrivateEndpointConnectionArgs{
			ParentName:                    pulumi.String("exampletopic1"),
			ParentType:                    pulumi.String("topics"),
			PrivateEndpointConnectionName: pulumi.String("BMTPE5.8A30D251-4C61-489D-A1AA-B37C4A329B8B"),
			PrivateLinkServiceConnectionState: &eventgrid.ConnectionStateArgs{
				ActionsRequired: pulumi.String("None"),
				Description:     pulumi.String("approving connection"),
				Status:          pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
