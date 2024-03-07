package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewConnectedEnvironmentsStorage(ctx, "connectedEnvironmentsStorage", &app.ConnectedEnvironmentsStorageArgs{
			ConnectedEnvironmentName: pulumi.String("env"),
			Properties: app.ConnectedEnvironmentStorageResponseProperties{
				AzureFile: &app.AzureFilePropertiesArgs{
					AccessMode:  pulumi.String("ReadOnly"),
					AccountKey:  pulumi.String("key"),
					AccountName: pulumi.String("account1"),
					ShareName:   pulumi.String("share1"),
				},
			},
			ResourceGroupName: pulumi.String("examplerg"),
			StorageName:       pulumi.String("jlaw-demo1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
