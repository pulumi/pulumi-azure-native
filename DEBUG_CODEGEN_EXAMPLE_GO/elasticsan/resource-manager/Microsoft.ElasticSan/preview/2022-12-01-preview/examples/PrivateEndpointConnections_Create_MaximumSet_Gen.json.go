package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/elasticsan/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticsan.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &elasticsan.PrivateEndpointConnectionArgs{
			ElasticSanName: pulumi.String("elasticsanname"),
			GroupIds: pulumi.StringArray{
				pulumi.String("sytxzqlcoapcaywthgwvwcw"),
			},
			PrivateEndpointConnectionName: pulumi.String("privateendpointconnectionname"),
			PrivateLinkServiceConnectionState: &elasticsan.PrivateLinkServiceConnectionStateArgs{
				ActionsRequired: pulumi.String("None"),
				Description:     pulumi.String("Auto-Approved"),
				Status:          pulumi.String("Pending"),
			},
			ResourceGroupName: pulumi.String("resourcegroupname"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
