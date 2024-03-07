package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &automation.PrivateEndpointConnectionArgs{
			AutomationAccountName:         pulumi.String("ddb1"),
			PrivateEndpointConnectionName: pulumi.String("privateEndpointConnectionName"),
			PrivateLinkServiceConnectionState: &automation.PrivateLinkServiceConnectionStatePropertyArgs{
				Description: pulumi.String("Approved by johndoe@contoso.com"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
