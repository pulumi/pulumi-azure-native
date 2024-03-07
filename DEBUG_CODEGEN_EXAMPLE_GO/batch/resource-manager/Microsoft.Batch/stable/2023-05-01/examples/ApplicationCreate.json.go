package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/batch/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.NewApplication(ctx, "application", &batch.ApplicationArgs{
			AccountName:       pulumi.String("sampleacct"),
			AllowUpdates:      pulumi.Bool(false),
			ApplicationName:   pulumi.String("app1"),
			DisplayName:       pulumi.String("myAppName"),
			ResourceGroupName: pulumi.String("default-azurebatch-japaneast"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
