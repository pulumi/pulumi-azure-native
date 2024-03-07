package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewLogProfile(ctx, "logProfile", &insights.LogProfileArgs{
			Categories: pulumi.StringArray{
				pulumi.String("Write"),
				pulumi.String("Delete"),
				pulumi.String("Action"),
			},
			Location: pulumi.String(""),
			Locations: pulumi.StringArray{
				pulumi.String("global"),
			},
			LogProfileName: pulumi.String("Rac46PostSwapRG"),
			RetentionPolicy: &insights.RetentionPolicyArgs{
				Days:    pulumi.Int(3),
				Enabled: pulumi.Bool(true),
			},
			ServiceBusRuleId: pulumi.String(""),
			StorageAccountId: pulumi.String("/subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/JohnKemTest/providers/Microsoft.Storage/storageAccounts/johnkemtest8162"),
			Tags:             nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
