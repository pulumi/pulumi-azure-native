package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewWorkloadNetworkDhcp(ctx, "workloadNetworkDhcp", &avs.WorkloadNetworkDhcpArgs{
			DhcpId:           pulumi.String("dhcp1"),
			PrivateCloudName: pulumi.String("cloud1"),
			Properties: avs.WorkloadNetworkDhcpServer{
				DhcpType:      "SERVER",
				DisplayName:   "dhcpConfigurations1",
				LeaseTime:     86400,
				Revision:      1,
				ServerAddress: "40.1.5.1/24",
			},
			ResourceGroupName: pulumi.String("group1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
