package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewRecordSet(ctx, "recordSet", &network.RecordSetArgs{
			Metadata: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			RecordType:            pulumi.String("A"),
			RelativeRecordSetName: pulumi.String("record1"),
			ResourceGroupName:     pulumi.String("rg1"),
			TargetResource: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/726f8cd6-6459-4db4-8e6d-2cd2716904e2/resourceGroups/test/providers/Microsoft.Network/trafficManagerProfiles/testpp2"),
			},
			Ttl:      pulumi.Float64(3600),
			ZoneName: pulumi.String("zone1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
