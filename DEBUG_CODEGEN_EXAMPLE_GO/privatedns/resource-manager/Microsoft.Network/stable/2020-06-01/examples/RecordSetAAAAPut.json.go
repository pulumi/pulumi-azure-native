package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewPrivateRecordSet(ctx, "privateRecordSet", &network.PrivateRecordSetArgs{
			AaaaRecords: network.AaaaRecordArray{
				&network.AaaaRecordArgs{
					Ipv6Address: pulumi.String("::1"),
				},
			},
			Metadata: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			PrivateZoneName:       pulumi.String("privatezone1.com"),
			RecordType:            pulumi.String("AAAA"),
			RelativeRecordSetName: pulumi.String("recordAAAA"),
			ResourceGroupName:     pulumi.String("resourceGroup1"),
			Ttl:                   pulumi.Float64(3600),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
