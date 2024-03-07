package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewWebAppSlot(ctx, "webAppSlot", &web.WebAppSlotArgs{
			Kind:              pulumi.String("app"),
			Location:          pulumi.String("East US"),
			Name:              pulumi.String("sitef6141"),
			ResourceGroupName: pulumi.String("testrg123"),
			ServerFarmId:      pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/DefaultAsp"),
			Slot:              pulumi.String("staging"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
