package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewVirtualMachineRunCommandByVirtualMachine(ctx, "virtualMachineRunCommandByVirtualMachine", &compute.VirtualMachineRunCommandByVirtualMachineArgs{
			AsyncExecution: pulumi.Bool(false),
			ErrorBlobUri:   pulumi.String("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt?sp=racw&st=2022-10-07T19:40:21Z&se=2022-10-08T03:40:21Z&spr=https&sv=2021-06-08&sr=b&sig=Yh7B%2Fy83olbYBdfsfbUREvd7ol8Dq5EVP3lAO4Kj4xDcN8%3D"),
			Location:       pulumi.String("West US"),
			OutputBlobManagedIdentity: &compute.RunCommandManagedIdentityArgs{
				ClientId: pulumi.String("22d35efb-0c99-4041-8c5b-6d24db33a69a"),
			},
			OutputBlobUri: pulumi.String("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
			Parameters: []compute.RunCommandInputParameterArgs{
				{
					Name:  pulumi.String("param1"),
					Value: pulumi.String("value1"),
				},
				{
					Name:  pulumi.String("param2"),
					Value: pulumi.String("value2"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			RunAsPassword:     pulumi.String("<runAsPassword>"),
			RunAsUser:         pulumi.String("user1"),
			RunCommandName:    pulumi.String("myRunCommand"),
			Source: &compute.VirtualMachineRunCommandScriptSourceArgs{
				ScriptUri: pulumi.String("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1?sp=r&st=2022-10-07T19:52:54Z&se=2022-10-08T03:52:54Z&spr=https&sv=2021-06-08&sr=b&sig=zfYFYCgea1PqVERZuwJiewrte5gjTnKGtVJngcw5oc828%3D"),
			},
			TimeoutInSeconds:                pulumi.Int(3600),
			TreatFailureAsDeploymentFailure: pulumi.Bool(false),
			VmName:                          pulumi.String("myVM"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
