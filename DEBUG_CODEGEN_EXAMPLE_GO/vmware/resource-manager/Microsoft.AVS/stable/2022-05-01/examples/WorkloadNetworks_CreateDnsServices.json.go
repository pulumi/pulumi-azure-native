package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewWorkloadNetworkDnsService(ctx, "workloadNetworkDnsService", &avs.WorkloadNetworkDnsServiceArgs{
			DefaultDnsZone: pulumi.String("defaultDnsZone1"),
			DisplayName:    pulumi.String("dnsService1"),
			DnsServiceId:   pulumi.String("dnsService1"),
			DnsServiceIp:   pulumi.String("5.5.5.5"),
			FqdnZones: pulumi.StringArray{
				pulumi.String("fqdnZone1"),
			},
			LogLevel:          pulumi.String("INFO"),
			PrivateCloudName:  pulumi.String("cloud1"),
			ResourceGroupName: pulumi.String("group1"),
			Revision:          pulumi.Float64(1),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
