package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewDeviceSecurityGroup(ctx, "deviceSecurityGroup", &security.DeviceSecurityGroupArgs{
			DeviceSecurityGroupName: pulumi.String("samplesecuritygroup"),
			ResourceId:              pulumi.String("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Devices/iotHubs/sampleiothub"),
			TimeWindowRules: []security.TimeWindowCustomAlertRuleArgs{
				nil,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
