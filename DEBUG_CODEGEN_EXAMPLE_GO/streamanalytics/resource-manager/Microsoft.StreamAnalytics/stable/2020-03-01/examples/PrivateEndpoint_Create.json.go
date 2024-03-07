package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewPrivateEndpoint(ctx, "privateEndpoint", &streamanalytics.PrivateEndpointArgs{
			ClusterName: pulumi.String("testcluster"),
			ManualPrivateLinkServiceConnections: streamanalytics.PrivateLinkServiceConnectionArray{
				&streamanalytics.PrivateLinkServiceConnectionArgs{
					GroupIds: pulumi.StringArray{
						pulumi.String("groupIdFromResource"),
					},
					PrivateLinkServiceId: pulumi.String("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
				},
			},
			PrivateEndpointName: pulumi.String("testpe"),
			ResourceGroupName:   pulumi.String("sjrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
