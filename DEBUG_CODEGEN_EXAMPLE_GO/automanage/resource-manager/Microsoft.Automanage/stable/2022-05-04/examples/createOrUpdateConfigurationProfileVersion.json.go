package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automanage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automanage.NewConfigurationProfilesVersion(ctx, "configurationProfilesVersion", &automanage.ConfigurationProfilesVersionArgs{
			ConfigurationProfileName: pulumi.String("customConfigurationProfile"),
			Location:                 pulumi.String("East US"),
			Properties: &automanage.ConfigurationProfilePropertiesArgs{
				Configuration: pulumi.Any{
					Antimalware / Enable:                false,
					AzureSecurityCenter / Enable:        true,
					Backup / Enable:                     false,
					BootDiagnostics / Enable:            true,
					ChangeTrackingAndInventory / Enable: true,
					GuestConfiguration / Enable:         true,
					LogAnalytics / Enable:               true,
					UpdateManagement / Enable:           true,
					VMInsights / Enable:                 true,
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroupName"),
			Tags: pulumi.StringMap{
				"Organization": pulumi.String("Administration"),
			},
			VersionName: pulumi.String("version1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
