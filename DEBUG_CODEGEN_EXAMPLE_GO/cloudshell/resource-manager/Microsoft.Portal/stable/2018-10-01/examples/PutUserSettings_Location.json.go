package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/portal/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := portal.NewUserSettingsWithLocation(ctx, "userSettingsWithLocation", &portal.UserSettingsWithLocationArgs{
			Location: pulumi.String("eastus"),
			Properties: portal.UserPropertiesResponse{
				PreferredLocation:  pulumi.String("eastus"),
				PreferredOsType:    pulumi.String("Linux"),
				PreferredShellType: pulumi.String("bash"),
				StorageProfile: &portal.StorageProfileArgs{
					DiskSizeInGB:             pulumi.Int(5),
					FileShareName:            pulumi.String("string"),
					StorageAccountResourceId: pulumi.String("string"),
				},
				TerminalSettings: &portal.TerminalSettingsArgs{
					FontSize:  pulumi.String("Medium"),
					FontStyle: pulumi.String("Monospace"),
				},
			},
			UserSettingsName: pulumi.String("cloudconsole"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
