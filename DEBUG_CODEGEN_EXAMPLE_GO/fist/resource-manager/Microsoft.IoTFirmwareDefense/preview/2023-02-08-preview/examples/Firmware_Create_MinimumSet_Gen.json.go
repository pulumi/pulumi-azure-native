package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotfirmwaredefense/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotfirmwaredefense.NewFirmware(ctx, "firmware", &iotfirmwaredefense.FirmwareArgs{
			FirmwareId:        pulumi.String("umrkdttp"),
			ResourceGroupName: pulumi.String("rgworkspaces-firmwares"),
			WorkspaceName:     pulumi.String("A7"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
