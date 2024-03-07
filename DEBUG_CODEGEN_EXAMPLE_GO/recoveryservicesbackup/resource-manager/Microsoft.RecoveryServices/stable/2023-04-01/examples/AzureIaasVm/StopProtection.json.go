package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewProtectedItem(ctx, "protectedItem", &recoveryservices.ProtectedItemArgs{
			ContainerName: pulumi.String("IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
			FabricName:    pulumi.String("Azure"),
			Properties: recoveryservices.AzureIaaSComputeVMProtectedItem{
				ProtectedItemType: "Microsoft.Compute/virtualMachines",
				ProtectionState:   "ProtectionStopped",
				SourceResourceId:  "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1",
			},
			ProtectedItemName: pulumi.String("VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
			ResourceGroupName: pulumi.String("SwaggerTestRg"),
			VaultName:         pulumi.String("NetSDKTestRsVault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
