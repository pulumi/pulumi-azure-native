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
			PrivateZoneName:       pulumi.String("privatezone1.com"),
			RecordType:            pulumi.String("TXT"),
			RelativeRecordSetName: pulumi.String("recordTXT"),
			ResourceGroupName:     pulumi.String("resourceGroup1"),
			Ttl:                   pulumi.Float64(3600),
			TxtRecords: network.TxtRecordArray{
				&network.TxtRecordArgs{
					Value: pulumi.StringArray{
						pulumi.String("string1"),
						pulumi.String("string2"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
