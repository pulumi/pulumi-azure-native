package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachineScaleSetVMRunCommand(ctx, "virtualMachineScaleSetVMRunCommand", &compute.VirtualMachineScaleSetVMRunCommandArgs{
			AsyncExecution:           pulumi.Bool(false),
			ErrorBlobManagedIdentity: nil,
			ErrorBlobUri:             pulumi.String("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
			InstanceId:               pulumi.String("0"),
			Location:                 pulumi.String("West US"),
			OutputBlobManagedIdentity: &compute.RunCommandManagedIdentityArgs{
				ClientId: pulumi.String("22d35efb-0c99-4041-8c5b-6d24db33a69a"),
			},
			OutputBlobUri: pulumi.String("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
			Parameters: compute.RunCommandInputParameterArray{
				&compute.RunCommandInputParameterArgs{
					Name:  pulumi.String("param1"),
					Value: pulumi.String("value1"),
				},
				&compute.RunCommandInputParameterArgs{
					Name:  pulumi.String("param2"),
					Value: pulumi.String("value2"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			RunAsPassword:     pulumi.String("<runAsPassword>"),
			RunAsUser:         pulumi.String("user1"),
			RunCommandName:    pulumi.String("myRunCommand"),
			Source: &compute.VirtualMachineRunCommandScriptSourceArgs{
				ScriptUri: pulumi.String("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1"),
				ScriptUriManagedIdentity: &compute.RunCommandManagedIdentityArgs{
					ObjectId: pulumi.String("4231e4d2-33e4-4e23-96b2-17888afa6072"),
				},
			},
			TimeoutInSeconds:                pulumi.Int(3600),
			TreatFailureAsDeploymentFailure: pulumi.Bool(true),
			VmScaleSetName:                  pulumi.String("myvmScaleSet"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
