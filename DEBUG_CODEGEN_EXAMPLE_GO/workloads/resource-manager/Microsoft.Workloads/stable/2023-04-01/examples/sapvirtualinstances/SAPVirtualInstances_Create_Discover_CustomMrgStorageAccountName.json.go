package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewSAPVirtualInstance(ctx, "sapVirtualInstance", &workloads.SAPVirtualInstanceArgs{
			Configuration: workloads.DiscoveryConfiguration{
				CentralServerVmId:           "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/sapq20scsvm0",
				ConfigurationType:           "Discovery",
				ManagedRgStorageAccountName: "q20saacssgrs",
			},
			Environment:            pulumi.String("NonProd"),
			Location:               pulumi.String("northeurope"),
			ResourceGroupName:      pulumi.String("test-rg"),
			SapProduct:             pulumi.String("S4HANA"),
			SapVirtualInstanceName: pulumi.String("X00"),
			Tags: pulumi.StringMap{
				"createdby": pulumi.String("abc@microsoft.com"),
				"test":      pulumi.String("abc"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
