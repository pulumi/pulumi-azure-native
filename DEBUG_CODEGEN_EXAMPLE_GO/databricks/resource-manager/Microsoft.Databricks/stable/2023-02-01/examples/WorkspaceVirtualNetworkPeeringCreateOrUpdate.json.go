package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databricks/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databricks.NewVNetPeering(ctx, "vNetPeering", &databricks.VNetPeeringArgs{
			AllowForwardedTraffic:     pulumi.Bool(false),
			AllowGatewayTransit:       pulumi.Bool(false),
			AllowVirtualNetworkAccess: pulumi.Bool(true),
			PeeringName:               pulumi.String("vNetPeeringTest"),
			RemoteVirtualNetwork: &databricks.VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkArgs{
				Id: pulumi.String("/subscriptions/0140911e-1040-48da-8bc9-b99fb3dd88a6/resourceGroups/subramantest/providers/Microsoft.Network/virtualNetworks/subramanvnet"),
			},
			ResourceGroupName: pulumi.String("rg"),
			UseRemoteGateways: pulumi.Bool(false),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
