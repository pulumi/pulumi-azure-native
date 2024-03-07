package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewWorkloadNetworkDnsZone(ctx, "workloadNetworkDnsZone", &avs.WorkloadNetworkDnsZoneArgs{
			DisplayName: pulumi.String("dnsZone1"),
			DnsServerIps: pulumi.StringArray{
				pulumi.String("1.1.1.1"),
			},
			DnsZoneId:         pulumi.String("dnsZone1"),
			Domain:            pulumi.StringArray{},
			PrivateCloudName:  pulumi.String("cloud1"),
			ResourceGroupName: pulumi.String("group1"),
			Revision:          pulumi.Float64(1),
			SourceIp:          pulumi.String("8.8.8.8"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
