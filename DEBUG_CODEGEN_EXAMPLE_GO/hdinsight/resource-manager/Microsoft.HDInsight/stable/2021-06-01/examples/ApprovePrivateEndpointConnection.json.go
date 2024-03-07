package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hdinsight/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hdinsight.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &hdinsight.PrivateEndpointConnectionArgs{
			ClusterName:                   pulumi.String("cluster1"),
			PrivateEndpointConnectionName: pulumi.String("testprivateep.b3bf5fed-9b12-4560-b7d0-2abe1bba07e2"),
			PrivateLinkServiceConnectionState: &hdinsight.PrivateLinkServiceConnectionStateArgs{
				ActionsRequired: pulumi.String("None"),
				Description:     pulumi.String("update it from pending to approved."),
				Status:          pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
