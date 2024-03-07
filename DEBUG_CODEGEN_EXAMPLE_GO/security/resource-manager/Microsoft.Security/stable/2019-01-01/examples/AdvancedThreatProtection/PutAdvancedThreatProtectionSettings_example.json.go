package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewAdvancedThreatProtection(ctx, "advancedThreatProtection", &security.AdvancedThreatProtectionArgs{
			IsEnabled:   pulumi.Bool(true),
			ResourceId:  pulumi.String("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount"),
			SettingName: pulumi.String("current"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
