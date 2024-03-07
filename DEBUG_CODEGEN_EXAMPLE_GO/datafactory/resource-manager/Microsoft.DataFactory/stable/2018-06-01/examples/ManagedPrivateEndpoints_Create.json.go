package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datafactory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datafactory.NewManagedPrivateEndpoint(ctx, "managedPrivateEndpoint", &datafactory.ManagedPrivateEndpointArgs{
			FactoryName:                pulumi.String("exampleFactoryName"),
			ManagedPrivateEndpointName: pulumi.String("exampleManagedPrivateEndpointName"),
			ManagedVirtualNetworkName:  pulumi.String("exampleManagedVirtualNetworkName"),
			Properties: &datafactory.ManagedPrivateEndpointTypeArgs{
				Fqdns:                 pulumi.StringArray{},
				GroupId:               pulumi.String("blob"),
				PrivateLinkResourceId: pulumi.String("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleBlobStorage"),
			},
			ResourceGroupName: pulumi.String("exampleResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
