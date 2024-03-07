package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewDataCollectionEndpoint(ctx, "dataCollectionEndpoint", &insights.DataCollectionEndpointArgs{
			DataCollectionEndpointName: pulumi.String("myCollectionEndpoint"),
			Location:                   pulumi.String("eastus"),
			NetworkAcls: &insights.DataCollectionEndpointNetworkAclsArgs{
				PublicNetworkAccess: pulumi.String("Enabled"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
