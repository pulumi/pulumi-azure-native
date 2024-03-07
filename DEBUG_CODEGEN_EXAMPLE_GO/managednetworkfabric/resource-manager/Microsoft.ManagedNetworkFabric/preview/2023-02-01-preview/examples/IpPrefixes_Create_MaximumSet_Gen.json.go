package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewIpPrefix(ctx, "ipPrefix", &managednetworkfabric.IpPrefixArgs{
			Annotation:   pulumi.String("annotationValue"),
			IpPrefixName: pulumi.String("example-ipPrefix"),
			IpPrefixRules: managednetworkfabric.IpPrefixPropertiesIpPrefixRulesArray{
				&managednetworkfabric.IpPrefixPropertiesIpPrefixRulesArgs{
					Action:           pulumi.String("Permit"),
					Condition:        pulumi.String("EqualTo"),
					NetworkPrefix:    pulumi.String("1.1.1.0/24"),
					SequenceNumber:   pulumi.Float64(12),
					SubnetMaskLength: pulumi.Int(28),
				},
			},
			Location:          pulumi.String("EastUS"),
			ResourceGroupName: pulumi.String("resourcegroupname"),
			Tags: pulumi.StringMap{
				"key6404": pulumi.String(""),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
