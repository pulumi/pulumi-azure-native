package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkcloud.NewCloudServicesNetwork(ctx, "cloudServicesNetwork", &networkcloud.CloudServicesNetworkArgs{
			AdditionalEgressEndpoints: []networkcloud.EgressEndpointArgs{
				{
					Category: pulumi.String("azure-resource-management"),
					Endpoints: networkcloud.EndpointDependencyArray{
						{
							DomainName: pulumi.String("https://storageaccountex.blob.core.windows.net"),
							Port:       pulumi.Float64(443),
						},
					},
				},
			},
			CloudServicesNetworkName:     pulumi.String("cloudServicesNetworkName"),
			EnableDefaultEgressEndpoints: pulumi.String("False"),
			ExtendedLocation: &networkcloud.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
				Type: pulumi.String("CustomLocation"),
			},
			Location:          pulumi.String("location"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("myvalue1"),
				"key2": pulumi.String("myvalue2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
