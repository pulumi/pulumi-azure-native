package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewDaprComponent(ctx, "daprComponent", &app.DaprComponentArgs{
			ComponentName:   pulumi.String("reddog"),
			ComponentType:   pulumi.String("state.azure.cosmosdb"),
			EnvironmentName: pulumi.String("myenvironment"),
			IgnoreErrors:    pulumi.Bool(false),
			InitTimeout:     pulumi.String("50s"),
			Metadata: app.DaprMetadataArray{
				&app.DaprMetadataArgs{
					Name:  pulumi.String("url"),
					Value: pulumi.String("<COSMOS-URL>"),
				},
				&app.DaprMetadataArgs{
					Name:  pulumi.String("database"),
					Value: pulumi.String("itemsDB"),
				},
				&app.DaprMetadataArgs{
					Name:  pulumi.String("collection"),
					Value: pulumi.String("items"),
				},
				&app.DaprMetadataArgs{
					Name:      pulumi.String("masterkey"),
					SecretRef: pulumi.String("masterkey"),
				},
			},
			ResourceGroupName: pulumi.String("examplerg"),
			Scopes: pulumi.StringArray{
				pulumi.String("container-app-1"),
				pulumi.String("container-app-2"),
			},
			SecretStoreComponent: pulumi.String("my-secret-store"),
			Version:              pulumi.String("v1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
