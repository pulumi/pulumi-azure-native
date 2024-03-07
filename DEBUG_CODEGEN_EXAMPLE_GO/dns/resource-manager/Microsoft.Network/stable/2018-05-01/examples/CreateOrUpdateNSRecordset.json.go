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
			NsRecords: []network.NsRecordArgs{
				{
					Nsdname: pulumi.String("ns1.contoso.com"),
				},
			},
			RecordType:            pulumi.String("NS"),
			RelativeRecordSetName: pulumi.String("record1"),
			ResourceGroupName:     pulumi.String("rg1"),
			Ttl:                   pulumi.Float64(3600),
			ZoneName:              pulumi.String("zone1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
