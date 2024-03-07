package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewComponentCurrentBillingFeature(ctx, "componentCurrentBillingFeature", &insights.ComponentCurrentBillingFeatureArgs{
			CurrentBillingFeatures: pulumi.StringArray{
				pulumi.String("Basic"),
				pulumi.String("Application Insights Enterprise"),
			},
			DataVolumeCap: &insights.ApplicationInsightsComponentDataVolumeCapArgs{
				Cap:                            pulumi.Float64(100),
				StopSendNotificationWhenHitCap: pulumi.Bool(true),
			},
			ResourceGroupName: pulumi.String("my-resource-group"),
			ResourceName:      pulumi.String("my-component"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
