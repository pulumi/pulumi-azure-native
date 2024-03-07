package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/standbypool/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := standbypool.NewStandbyVirtualMachinePool(ctx, "standbyVirtualMachinePool", &standbypool.StandbyVirtualMachinePoolArgs{
			AttachedVirtualMachineScaleSetId: pulumi.String("/subscriptions/8CC31D61-82D7-4B2B-B9DC-6B924DE7D229/resourceGroups/vmssRg/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
			ElasticityProfile: &standbypool.StandbyVirtualMachinePoolElasticityProfileArgs{
				MaxReadyCapacity: pulumi.Float64(304),
			},
			Location:                      pulumi.String("West US"),
			ResourceGroupName:             pulumi.String("rgstandbypool"),
			StandbyVirtualMachinePoolName: pulumi.String("pool"),
			Tags:                          nil,
			VirtualMachineState:           pulumi.String("Running"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
