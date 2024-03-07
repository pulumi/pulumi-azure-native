package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewWorkloadNetworkPublicIP(ctx, "workloadNetworkPublicIP", &avs.WorkloadNetworkPublicIPArgs{
			DisplayName:       pulumi.String("publicIP1"),
			NumberOfPublicIPs: pulumi.Float64(32),
			PrivateCloudName:  pulumi.String("cloud1"),
			PublicIPId:        pulumi.String("publicIP1"),
			ResourceGroupName: pulumi.String("group1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
