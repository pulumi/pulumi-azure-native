package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewDisk(ctx, "disk", &devtestlab.DiskArgs{
			DiskSizeGiB:       pulumi.Int(1023),
			DiskType:          pulumi.String("Standard"),
			LabName:           pulumi.String("{labName}"),
			LeasedByLabVmId:   pulumi.String("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/vmName"),
			Name:              pulumi.String("{diskName}"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			UserName:          pulumi.String("{userId}"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
