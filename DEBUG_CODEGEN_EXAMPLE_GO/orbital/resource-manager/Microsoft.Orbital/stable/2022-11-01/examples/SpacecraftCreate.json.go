package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/orbital/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := orbital.NewSpacecraft(ctx, "spacecraft", &orbital.SpacecraftArgs{
			Links: orbital.SpacecraftLinkArray{
				&orbital.SpacecraftLinkArgs{
					BandwidthMHz:       pulumi.Float64(2),
					CenterFrequencyMHz: pulumi.Float64(2250),
					Direction:          pulumi.String("Uplink"),
					Name:               pulumi.String("uplink_lhcp1"),
					Polarization:       pulumi.String("LHCP"),
				},
				&orbital.SpacecraftLinkArgs{
					BandwidthMHz:       pulumi.Float64(15),
					CenterFrequencyMHz: pulumi.Float64(8160),
					Direction:          pulumi.String("Downlink"),
					Name:               pulumi.String("downlink_rhcp1"),
					Polarization:       pulumi.String("RHCP"),
				},
			},
			Location:          pulumi.String("eastus2"),
			NoradId:           pulumi.String("36411"),
			ResourceGroupName: pulumi.String("contoso-Rgp"),
			SpacecraftName:    pulumi.String("CONTOSO_SAT"),
			TitleLine:         pulumi.String("CONTOSO_SAT"),
			TleLine1:          pulumi.String("1 27424U 02022A   22167.05119303  .00000638  00000+0  15103-3 0  9994"),
			TleLine2:          pulumi.String("2 27424  98.2477 108.9546 0000928  92.9194 327.0802 14.57300770 69982"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
