package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/extendedlocation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := extendedlocation.NewCustomLocation(ctx, "customLocation", &extendedlocation.CustomLocationArgs{
			Authentication: &extendedlocation.CustomLocationPropertiesAuthenticationArgs{
				Type:  pulumi.String("KubeConfig"),
				Value: pulumi.String("<base64 KubeConfig>"),
			},
			ClusterExtensionIds: pulumi.StringArray{
				pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedCluster/someCluster/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension"),
			},
			DisplayName:    pulumi.String("customLocationLocation01"),
			HostResourceId: pulumi.String("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01"),
			Identity: &extendedlocation.IdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Location:          pulumi.String("West US"),
			Namespace:         pulumi.String("namespace01"),
			ResourceGroupName: pulumi.String("testresourcegroup"),
			ResourceName:      pulumi.String("customLocation01"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
