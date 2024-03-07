package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewPrivateRecordSet(ctx, "privateRecordSet", &network.PrivateRecordSetArgs{
			Metadata: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			PrivateZoneName: pulumi.String("0.0.127.in-addr.arpa"),
			PtrRecords: network.PtrRecordArray{
				&network.PtrRecordArgs{
					Ptrdname: pulumi.String("localhost"),
				},
			},
			RecordType:            pulumi.String("PTR"),
			RelativeRecordSetName: pulumi.String("1"),
			ResourceGroupName:     pulumi.String("resourceGroup1"),
			Ttl:                   pulumi.Float64(3600),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
