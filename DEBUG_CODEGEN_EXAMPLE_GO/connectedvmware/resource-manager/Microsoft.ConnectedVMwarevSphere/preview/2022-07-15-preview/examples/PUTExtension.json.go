package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/connectedvmwarevsphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connectedvmwarevsphere.NewMachineExtension(ctx, "machineExtension", &connectedvmwarevsphere.MachineExtensionArgs{
			ExtensionName:     pulumi.String("CustomScriptExtension"),
			Location:          pulumi.String("eastus2euap"),
			Publisher:         pulumi.String("Microsoft.Compute"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Settings: pulumi.Any{
				CommandToExecute: "powershell.exe -c \"Get-Process | Where-Object { $_.CPU -gt 10000 }\"",
			},
			Type:               pulumi.String("CustomScriptExtension"),
			TypeHandlerVersion: pulumi.String("1.10"),
			VirtualMachineName: pulumi.String("myMachine"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
