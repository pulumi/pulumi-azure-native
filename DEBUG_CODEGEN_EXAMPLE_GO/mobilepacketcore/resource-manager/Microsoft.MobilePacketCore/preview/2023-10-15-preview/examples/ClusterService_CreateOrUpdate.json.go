package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilepacketcore/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilepacketcore.NewClusterService(ctx, "clusterService", &mobilepacketcore.ClusterServiceArgs{
			ClusterServiceName: pulumi.String("clusterService1"),
			ClusterTypeSpecificData: mobilepacketcore.ClusterServiceNexusAksClusterData{
				CustomLocationId: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/cluster124Location",
				Type:             "NexusAks",
			},
			ComponentParameters: []mobilepacketcore.QualifiedComponentDeploymentParametersArgs{
				{
					Parameters: pulumi.String("{\"global\": {\"registry\": {\"docker\": []}}}"),
					Secrets:    pulumi.String("{\"global\": {\"secret\": {\"secretValue\": []}}}"),
					Type:       pulumi.String("fed-crds"),
				},
			},
			DeploymentType:    pulumi.String("Production"),
			Location:          pulumi.String("eastus"),
			ReleaseVersion:    pulumi.String("4.3.0-alpha"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
