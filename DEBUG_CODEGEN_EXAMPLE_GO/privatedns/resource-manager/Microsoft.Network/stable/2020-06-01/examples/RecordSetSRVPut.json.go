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
			RecordType:            pulumi.String("SRV"),
			RelativeRecordSetName: pulumi.String("recordSRV"),
			ResourceGroupName:     pulumi.String("resourceGroup1"),
			SrvRecords: []network.SrvRecordArgs{
				{
					Port:     pulumi.Int(80),
					Priority: pulumi.Int(0),
					Target:   pulumi.String("contoso.com"),
					Weight:   pulumi.Int(10),
				},
			},
			Ttl: pulumi.Float64(3600),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
