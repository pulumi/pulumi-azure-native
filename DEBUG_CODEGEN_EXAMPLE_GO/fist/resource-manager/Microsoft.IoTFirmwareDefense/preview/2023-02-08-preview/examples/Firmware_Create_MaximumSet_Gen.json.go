package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotfirmwaredefense/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotfirmwaredefense.NewFirmware(ctx, "firmware", &iotfirmwaredefense.FirmwareArgs{
			Description:       pulumi.String("uz"),
			FileName:          pulumi.String("wresexxulcdsdd"),
			FileSize:          pulumi.Float64(17),
			FirmwareId:        pulumi.String("umrkdttp"),
			Model:             pulumi.String("f"),
			ResourceGroupName: pulumi.String("rgworkspaces-firmwares"),
			Status:            pulumi.String("Pending"),
			StatusMessages: pulumi.Array{
				pulumi.Any{
					Message: "ulvhmhokezathzzauiitu",
				},
			},
			Vendor:        pulumi.String("vycmdhgtmepcptyoubztiuudpkcpd"),
			Version:       pulumi.String("s"),
			WorkspaceName: pulumi.String("A7"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
