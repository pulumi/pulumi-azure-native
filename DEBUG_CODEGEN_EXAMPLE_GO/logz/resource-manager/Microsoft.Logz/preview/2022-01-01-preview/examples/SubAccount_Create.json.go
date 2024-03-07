package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logz/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logz.NewSubAccount(ctx, "subAccount", &logz.SubAccountArgs{
			MonitorName:       pulumi.String("myMonitor"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SubAccountName:    pulumi.String("SubAccount1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
