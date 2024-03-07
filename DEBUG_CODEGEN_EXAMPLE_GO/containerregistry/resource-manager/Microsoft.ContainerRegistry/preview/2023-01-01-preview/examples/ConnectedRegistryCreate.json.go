package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewConnectedRegistry(ctx, "connectedRegistry", &containerregistry.ConnectedRegistryArgs{
			ClientTokenIds: pulumi.StringArray{
				pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),
			},
			ConnectedRegistryName: pulumi.String("myConnectedRegistry"),
			Mode:                  pulumi.String("ReadWrite"),
			NotificationsList: pulumi.StringArray{
				pulumi.String("hello-world:*:*"),
				pulumi.String("sample/repo/*:1.0:*"),
			},
			Parent: containerregistry.ParentPropertiesResponse{
				SyncProperties: &containerregistry.SyncPropertiesArgs{
					MessageTtl: pulumi.String("P2D"),
					Schedule:   pulumi.String("0 9 * * *"),
					SyncWindow: pulumi.String("PT3H"),
					TokenId:    pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
				},
			},
			RegistryName:      pulumi.String("myRegistry"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
