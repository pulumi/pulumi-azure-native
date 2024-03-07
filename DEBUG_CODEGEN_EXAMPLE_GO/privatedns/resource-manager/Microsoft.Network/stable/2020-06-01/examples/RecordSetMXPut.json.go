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
			MxRecords: []network.MxRecordArgs{
				{
					Exchange:   pulumi.String("mail.privatezone1.com"),
					Preference: pulumi.Int(0),
				},
			},
			PrivateZoneName:       pulumi.String("privatezone1.com"),
			RecordType:            pulumi.String("MX"),
			RelativeRecordSetName: pulumi.String("recordMX"),
			ResourceGroupName:     pulumi.String("resourceGroup1"),
			Ttl:                   pulumi.Float64(3600),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
