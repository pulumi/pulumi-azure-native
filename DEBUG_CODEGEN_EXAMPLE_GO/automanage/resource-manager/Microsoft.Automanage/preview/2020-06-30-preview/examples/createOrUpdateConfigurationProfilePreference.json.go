package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automanage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automanage.NewConfigurationProfilePreference(ctx, "configurationProfilePreference", &automanage.ConfigurationProfilePreferenceArgs{
			ConfigurationProfilePreferenceName: pulumi.String("defaultProfilePreference"),
			Location:                           pulumi.String("East US"),
			Properties: automanage.ConfigurationProfilePreferencePropertiesResponse{
				AntiMalware: &automanage.ConfigurationProfilePreferenceAntiMalwareArgs{
					EnableRealTimeProtection: pulumi.String("True"),
				},
				VmBackup: &automanage.ConfigurationProfilePreferenceVmBackupArgs{
					TimeZone: pulumi.String("Pacific Standard Time"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroupName"),
			Tags: pulumi.StringMap{
				"Organization": pulumi.String("Administration"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
