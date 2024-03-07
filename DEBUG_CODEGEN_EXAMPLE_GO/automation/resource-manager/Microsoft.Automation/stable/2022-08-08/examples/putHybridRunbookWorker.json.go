package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewHybridRunbookWorker(ctx, "hybridRunbookWorker", &automation.HybridRunbookWorkerArgs{
			AutomationAccountName:        pulumi.String("testaccount"),
			HybridRunbookWorkerGroupName: pulumi.String("TestHybridGroup"),
			HybridRunbookWorkerId:        pulumi.String("c010ad12-ef14-4a2a-aa9e-ef22c4745ddd"),
			ResourceGroupName:            pulumi.String("rg"),
			VmResourceId:                 pulumi.String("/subscriptions/vmsubid/resourceGroups/vmrg/providers/Microsoft.Compute/virtualMachines/vmname"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
