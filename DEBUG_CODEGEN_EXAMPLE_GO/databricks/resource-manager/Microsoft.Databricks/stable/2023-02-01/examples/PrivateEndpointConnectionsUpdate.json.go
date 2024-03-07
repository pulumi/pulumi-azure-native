package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databricks/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databricks.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &databricks.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("myWorkspace.23456789-1111-1111-1111-111111111111"),
			Properties: &databricks.PrivateEndpointConnectionPropertiesArgs{
				PrivateLinkServiceConnectionState: &databricks.PrivateLinkServiceConnectionStateArgs{
					Description: pulumi.String("Approved by databricksadmin@contoso.com"),
					Status:      pulumi.String("Approved"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
