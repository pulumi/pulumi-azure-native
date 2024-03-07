package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dashboard/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dashboard.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &dashboard.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("myConnection"),
			ResourceGroupName:             pulumi.String("myResourceGroup"),
			WorkspaceName:                 pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
