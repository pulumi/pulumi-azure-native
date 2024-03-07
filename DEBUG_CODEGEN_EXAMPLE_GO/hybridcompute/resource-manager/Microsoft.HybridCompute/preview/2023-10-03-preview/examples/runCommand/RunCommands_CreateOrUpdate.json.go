package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcompute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcompute.NewMachineRunCommand(ctx, "machineRunCommand", &hybridcompute.MachineRunCommandArgs{
			AsyncExecution: pulumi.Bool(false),
			ErrorBlobUri:   pulumi.String("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
			Location:       pulumi.String("eastus2"),
			MachineName:    pulumi.String("myMachine"),
			OutputBlobUri:  pulumi.String("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
			Parameters: hybridcompute.RunCommandInputParameterArray{
				&hybridcompute.RunCommandInputParameterArgs{
					Name:  pulumi.String("param1"),
					Value: pulumi.String("value1"),
				},
				&hybridcompute.RunCommandInputParameterArgs{
					Name:  pulumi.String("param2"),
					Value: pulumi.String("value2"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			RunAsPassword:     pulumi.String("<runAsPassword>"),
			RunAsUser:         pulumi.String("user1"),
			RunCommandName:    pulumi.String("myRunCommand"),
			Source: &hybridcompute.MachineRunCommandScriptSourceArgs{
				Script: pulumi.String("Write-Host Hello World!"),
			},
			TimeoutInSeconds: pulumi.Int(3600),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
