package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewGlobalReachConnection(ctx, "globalReachConnection", &avs.GlobalReachConnectionArgs{
			AuthorizationKey:          pulumi.String("01010101-0101-0101-0101-010101010101"),
			GlobalReachConnectionName: pulumi.String("connection1"),
			PeerExpressRouteCircuit:   pulumi.String("/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/mygroup/providers/Microsoft.Network/expressRouteCircuits/mypeer"),
			PrivateCloudName:          pulumi.String("cloud1"),
			ResourceGroupName:         pulumi.String("group1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
