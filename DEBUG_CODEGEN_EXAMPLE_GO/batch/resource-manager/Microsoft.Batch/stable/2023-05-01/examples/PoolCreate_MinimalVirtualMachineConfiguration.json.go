package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/batch/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.NewPool(ctx, "pool", &batch.PoolArgs{
			AccountName: pulumi.String("sampleacct"),
			DeploymentConfiguration: &batch.DeploymentConfigurationArgs{
				VirtualMachineConfiguration: &batch.VirtualMachineConfigurationArgs{
					ImageReference: &batch.ImageReferenceArgs{
						Offer:     pulumi.String("UbuntuServer"),
						Publisher: pulumi.String("Canonical"),
						Sku:       pulumi.String("18.04-LTS"),
						Version:   pulumi.String("latest"),
					},
					NodeAgentSkuId: pulumi.String("batch.node.ubuntu 18.04"),
				},
			},
			PoolName:          pulumi.String("testpool"),
			ResourceGroupName: pulumi.String("default-azurebatch-japaneast"),
			ScaleSettings: &batch.ScaleSettingsArgs{
				AutoScale: &batch.AutoScaleSettingsArgs{
					EvaluationInterval: pulumi.String("PT5M"),
					Formula:            pulumi.String("$TargetDedicatedNodes=1"),
				},
			},
			VmSize: pulumi.String("STANDARD_D4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
