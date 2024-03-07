package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/deviceregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := deviceregistry.NewAsset(ctx, "asset", &deviceregistry.AssetArgs{
			AssetEndpointProfileUri: pulumi.String("https://www.example.com/myAssetEndpointProfile"),
			AssetName:               pulumi.String("my-asset"),
			AssetType:               pulumi.String("MyAssetType"),
			DataPoints: []deviceregistry.DataPointArgs{
				{
					CapabilityId:           pulumi.String("dtmi:com:example:Thermostat:__temperature;1"),
					DataPointConfiguration: pulumi.String("{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}"),
					DataSource:             pulumi.String("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt1"),
					ObservabilityMode:      pulumi.String("counter"),
				},
				{
					CapabilityId:           pulumi.String("dtmi:com:example:Thermostat:__pressure;1"),
					DataPointConfiguration: pulumi.String("{\"publishingInterval\":4,\"samplingInterval\":4,\"queueSize\":7}"),
					DataSource:             pulumi.String("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt2"),
					ObservabilityMode:      pulumi.String("none"),
				},
			},
			DefaultDataPointsConfiguration: pulumi.String("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
			DefaultEventsConfiguration:     pulumi.String("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}"),
			Description:                    pulumi.String("This is a sample Asset"),
			DisplayName:                    pulumi.String("AssetDisplayName"),
			DocumentationUri:               pulumi.String("https://www.example.com/manual"),
			Enabled:                        pulumi.Bool(true),
			Events: []deviceregistry.EventArgs{
				{
					CapabilityId:       pulumi.String("dtmi:com:example:Thermostat:__temperature;1"),
					EventConfiguration: pulumi.String("{\"publishingInterval\":7,\"samplingInterval\":1,\"queueSize\":8}"),
					EventNotifier:      pulumi.String("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt3"),
					ObservabilityMode:  pulumi.String("none"),
				},
				{
					CapabilityId:       pulumi.String("dtmi:com:example:Thermostat:__pressure;1"),
					EventConfiguration: pulumi.String("{\"publishingInterval\":7,\"samplingInterval\":8,\"queueSize\":4}"),
					EventNotifier:      pulumi.String("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt4"),
					ObservabilityMode:  pulumi.String("log"),
				},
			},
			ExtendedLocation: &deviceregistry.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
				Type: pulumi.String("CustomLocation"),
			},
			HardwareRevision:  pulumi.String("1.0"),
			Location:          pulumi.String("West Europe"),
			Manufacturer:      pulumi.String("Contoso"),
			ManufacturerUri:   pulumi.String("https://www.contoso.com/manufacturerUri"),
			Model:             pulumi.String("ContosoModel"),
			ProductCode:       pulumi.String("SA34VDG"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SerialNumber:      pulumi.String("64-103816-519918-8"),
			SoftwareRevision:  pulumi.String("2.0"),
			Tags: pulumi.StringMap{
				"site": pulumi.String("building-1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
