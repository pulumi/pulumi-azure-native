package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewMoveCollection(ctx, "moveCollection", &migrate.MoveCollectionArgs{
			Identity: &migrate.IdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Location:           pulumi.String("eastus2"),
			MoveCollectionName: pulumi.String("movecollection1"),
			Properties: &migrate.MoveCollectionPropertiesArgs{
				SourceRegion: pulumi.String("eastus"),
				TargetRegion: pulumi.String("westus"),
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
