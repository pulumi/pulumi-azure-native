package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcompute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcompute.NewMachineExtension(ctx, "machineExtension", &hybridcompute.MachineExtensionArgs{
			ExtensionName: pulumi.String("CustomScriptExtension"),
			Location:      pulumi.String("eastus2euap"),
			MachineName:   pulumi.String("myMachine"),
			Properties: &hybridcompute.MachineExtensionPropertiesArgs{
				Publisher: pulumi.String("Microsoft.Compute"),
				Settings: pulumi.Any{
					CommandToExecute: "powershell.exe -c \"Get-Process | Where-Object { $_.CPU -gt 10000 }\"",
				},
				Type:               pulumi.String("CustomScriptExtension"),
				TypeHandlerVersion: pulumi.String("1.10"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
