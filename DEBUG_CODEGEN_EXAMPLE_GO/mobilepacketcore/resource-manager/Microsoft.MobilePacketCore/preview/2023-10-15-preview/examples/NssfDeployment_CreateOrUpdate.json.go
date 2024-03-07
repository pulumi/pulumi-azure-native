package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilepacketcore/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilepacketcore.NewNssfDeployment(ctx, "nssfDeployment", &mobilepacketcore.NssfDeploymentArgs{
			ClusterService:      pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobilePacketCore/clusterServices/byoCluster"),
			ComponentParameters: pulumi.String("{\"global\": {\"registry\": {\"docker\": []}}}"),
			Location:            pulumi.String("eastus"),
			NssfDeploymentName:  pulumi.String("nssfDeployment1"),
			ResourceGroupName:   pulumi.String("rg1"),
			SecretsParameters:   pulumi.String("{\"global\": {\"secret\": {\"secretValue\": []}}}"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
