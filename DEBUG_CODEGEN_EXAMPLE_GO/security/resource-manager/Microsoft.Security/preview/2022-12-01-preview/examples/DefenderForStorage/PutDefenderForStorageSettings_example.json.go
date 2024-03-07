package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewDefenderForStorage(ctx, "defenderForStorage", &security.DefenderForStorageArgs{
			IsEnabled: pulumi.Bool(true),
			MalwareScanning: &security.MalwareScanningPropertiesArgs{
				CapGBPerMonth:                       -1,
				IsEnabled:                           pulumi.Bool(true),
				ScanResultsEventGridTopicResourceId: pulumi.String("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.EventGrid/topics/sampletopic"),
			},
			OverrideSubscriptionLevelSettings: pulumi.Bool(true),
			ResourceId:                        pulumi.String("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount"),
			SensitiveDataDiscovery: &security.SensitiveDataDiscoveryPropertiesArgs{
				IsEnabled: pulumi.Bool(true),
			},
			SettingName: pulumi.String("current"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
