package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewCustomImage(ctx, "customImage", &devtestlab.CustomImageArgs{
			Description:       pulumi.String("My Custom Image"),
			LabName:           pulumi.String("{labName}"),
			Name:              pulumi.String("{customImageName}"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"tagName1": pulumi.String("tagValue1"),
			},
			Vm: devtestlab.CustomImagePropertiesFromVmResponse{
				LinuxOsInfo: &devtestlab.LinuxOsInfoArgs{
					LinuxOsState: pulumi.String("NonDeprovisioned"),
				},
				SourceVmId: pulumi.String("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/{vmName}"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
