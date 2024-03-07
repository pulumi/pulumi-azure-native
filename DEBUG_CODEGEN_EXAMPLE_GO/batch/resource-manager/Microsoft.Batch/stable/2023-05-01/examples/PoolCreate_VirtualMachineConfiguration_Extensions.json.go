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
					Extensions: batch.VMExtensionArray{
						&batch.VMExtensionArgs{
							AutoUpgradeMinorVersion: pulumi.Bool(true),
							EnableAutomaticUpgrade:  pulumi.Bool(true),
							Name:                    pulumi.String("batchextension1"),
							Publisher:               pulumi.String("Microsoft.Azure.KeyVault"),
							Settings: pulumi.Any{
								AuthenticationSettingsKey:    "authenticationSettingsValue",
								SecretsManagementSettingsKey: "secretsManagementSettingsValue",
							},
							Type:               pulumi.String("KeyVaultForLinux"),
							TypeHandlerVersion: pulumi.String("2.0"),
						},
					},
					ImageReference: &batch.ImageReferenceArgs{
						Offer:     pulumi.String("0001-com-ubuntu-server-focal"),
						Publisher: pulumi.String("Canonical"),
						Sku:       pulumi.String("20_04-lts"),
					},
					NodeAgentSkuId: pulumi.String("batch.node.ubuntu 20.04"),
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
			TargetNodeCommunicationMode: batch.NodeCommunicationModeDefault,
			VmSize:                      pulumi.String("STANDARD_D4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
