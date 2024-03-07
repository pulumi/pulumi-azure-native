package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/costmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costmanagement.NewSetting(ctx, "setting", &costmanagement.SettingArgs{
			Cache: []costmanagement.SettingsPropertiesCacheArgs{
				{
					Channel:    pulumi.String("Modern"),
					Id:         pulumi.String("/providers/Microsoft.Management/managementGroups/72f988bf-86f1-41af-91ab-2d7cd011db47"),
					Name:       pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
					Parent:     pulumi.String("/providers/Microsoft.Management/managementGroups/acm"),
					Status:     pulumi.String("enabled"),
					Subchannel: pulumi.String("NotApplicable"),
				},
			},
			Scope:       pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000"),
			SettingName: pulumi.String("myscope"),
			StartOn:     pulumi.String("LastUsed"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
