package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewWebAppSlot(ctx, "webAppSlot", &web.WebAppSlotArgs{
			CloningInfo: &web.CloningInfoArgs{
				AppSettingsOverrides: pulumi.StringMap{
					"Setting1": pulumi.String("NewValue1"),
					"Setting3": pulumi.String("NewValue5"),
				},
				CloneCustomHostNames:   pulumi.Bool(true),
				CloneSourceControl:     pulumi.Bool(true),
				ConfigureLoadBalancing: pulumi.Bool(false),
				HostingEnvironment:     pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg456/providers/Microsoft.Web/hostingenvironments/aseforsites"),
				Overwrite:              pulumi.Bool(false),
				SourceWebAppId:         pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg456/providers/Microsoft.Web/sites/srcsiteg478/slot/qa"),
				SourceWebAppLocation:   pulumi.String("West Europe"),
			},
			Kind:              pulumi.String("app"),
			Location:          pulumi.String("East US"),
			Name:              pulumi.String("sitef6141"),
			ResourceGroupName: pulumi.String("testrg123"),
			Slot:              pulumi.String("staging"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
