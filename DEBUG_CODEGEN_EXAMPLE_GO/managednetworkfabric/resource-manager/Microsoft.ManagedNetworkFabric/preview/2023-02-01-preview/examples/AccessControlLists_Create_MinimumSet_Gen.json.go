package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewAccessControlList(ctx, "accessControlList", &managednetworkfabric.AccessControlListArgs{
			AccessControlListName: pulumi.String("aclOne"),
			AddressFamily:         pulumi.String("ipv4"),
			Conditions: []managednetworkfabric.AccessControlListConditionPropertiesArgs{
				{
					Action:             pulumi.String("allow"),
					DestinationAddress: pulumi.String("1.1.1.1"),
					DestinationPort:    pulumi.String("21"),
					Protocol:           pulumi.Int(6),
					SequenceNumber:     pulumi.Int(3),
					SourceAddress:      pulumi.String("2.2.2.2"),
					SourcePort:         pulumi.String("65000"),
				},
			},
			Location:          pulumi.String("EastUs"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
