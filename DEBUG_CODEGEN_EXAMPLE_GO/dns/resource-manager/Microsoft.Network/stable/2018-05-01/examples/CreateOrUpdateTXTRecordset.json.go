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
			RecordType:            pulumi.String("TXT"),
			RelativeRecordSetName: pulumi.String("record1"),
			ResourceGroupName:     pulumi.String("rg1"),
			Ttl:                   pulumi.Float64(3600),
			TxtRecords: network.TxtRecordArray{
				&network.TxtRecordArgs{
					Value: pulumi.StringArray{
						pulumi.String("string1"),
						pulumi.String("string2"),
					},
				},
			},
			ZoneName: pulumi.String("zone1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
