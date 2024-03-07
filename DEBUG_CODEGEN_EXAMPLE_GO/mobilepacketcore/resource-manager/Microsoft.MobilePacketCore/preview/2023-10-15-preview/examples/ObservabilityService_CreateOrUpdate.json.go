package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilepacketcore/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilepacketcore.NewObservabilityService(ctx, "observabilityService", &mobilepacketcore.ObservabilityServiceArgs{
			ClusterService: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobilePacketCore/clusterServices/byoCluster"),
			ComponentParameters: mobilepacketcore.QualifiedComponentDeploymentParametersArray{
				&mobilepacketcore.QualifiedComponentDeploymentParametersArgs{
					Parameters: pulumi.String("{\"global\": {\"registry\": {\"docker\": []}}}"),
					Secrets:    pulumi.String("{\"global\": {\"secret\": {\"secretValue\": []}}}"),
					Type:       pulumi.String("fed-crds"),
				},
			},
			Location:                 pulumi.String("eastus"),
			ObservabilityServiceName: pulumi.String("observabilityService1"),
			ResourceGroupName:        pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
