package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/serialconsole/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := serialconsole.NewSerialPort(ctx, "serialPort", &serialconsole.SerialPortArgs{
			ParentResource:            pulumi.String("myVM"),
			ParentResourceType:        pulumi.String("virtualMachines"),
			ResourceGroupName:         pulumi.String("myResourceGroup"),
			ResourceProviderNamespace: pulumi.String("Microsoft.Compute"),
			SerialPort:                pulumi.String("0"),
			State:                     serialconsole.SerialPortStateEnabled,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
