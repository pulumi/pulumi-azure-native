package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewDscpConfiguration(ctx, "dscpConfiguration", &network.DscpConfigurationArgs{
			DscpConfigurationName: pulumi.String("mydscpconfig"),
			Location:              pulumi.String("eastus"),
			QosDefinitionCollection: []network.QosDefinitionArgs{
				{
					DestinationIpRanges: network.QosIpRangeArray{
						{
							EndIP:   pulumi.String("127.0.10.2"),
							StartIP: pulumi.String("127.0.10.1"),
						},
					},
					DestinationPortRanges: network.QosPortRangeArray{
						{
							End:   pulumi.Int(15),
							Start: pulumi.Int(15),
						},
					},
					Markings: pulumi.IntArray{
						pulumi.Int(1),
					},
					Protocol: pulumi.String("Tcp"),
					SourceIpRanges: network.QosIpRangeArray{
						{
							EndIP:   pulumi.String("127.0.0.2"),
							StartIP: pulumi.String("127.0.0.1"),
						},
					},
					SourcePortRanges: network.QosPortRangeArray{
						{
							End:   pulumi.Int(11),
							Start: pulumi.Int(10),
						},
						{
							End:   pulumi.Int(21),
							Start: pulumi.Int(20),
						},
					},
				},
				{
					DestinationIpRanges: network.QosIpRangeArray{
						{
							EndIP:   pulumi.String("12.0.10.2"),
							StartIP: pulumi.String("12.0.10.1"),
						},
					},
					DestinationPortRanges: network.QosPortRangeArray{
						{
							End:   pulumi.Int(52),
							Start: pulumi.Int(51),
						},
					},
					Markings: pulumi.IntArray{
						pulumi.Int(2),
					},
					Protocol: pulumi.String("Udp"),
					SourceIpRanges: network.QosIpRangeArray{
						{
							EndIP:   pulumi.String("12.0.0.2"),
							StartIP: pulumi.String("12.0.0.1"),
						},
					},
					SourcePortRanges: network.QosPortRangeArray{
						{
							End:   pulumi.Int(12),
							Start: pulumi.Int(11),
						},
					},
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
