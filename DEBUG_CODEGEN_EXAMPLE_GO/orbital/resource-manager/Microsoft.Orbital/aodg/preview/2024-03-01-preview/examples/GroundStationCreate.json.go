package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/orbital/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := orbital.NewGroundStation(ctx, "groundStation", &orbital.GroundStationArgs{
			AltitudeMeters: pulumi.Float64(1500.83),
			Capabilities: pulumi.StringArray{
				pulumi.String("Communication"),
			},
			City: pulumi.String("redmond"),
			GlobalCommunicationsSite: &orbital.GroundStationsPropertiesGlobalCommunicationsSiteArgs{
				Id: pulumi.String("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/providers/Microsoft.Orbital/globalCommunicationsSites/contoso-Vernon"),
			},
			GroundStationName: pulumi.String("westus_gs1"),
			LatitudeDegrees:   -122.122,
			Location:          pulumi.String("westus"),
			LongitudeDegrees:  pulumi.Float64(47.674),
			ProviderName:      pulumi.String("Microsoft"),
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
				"key2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
