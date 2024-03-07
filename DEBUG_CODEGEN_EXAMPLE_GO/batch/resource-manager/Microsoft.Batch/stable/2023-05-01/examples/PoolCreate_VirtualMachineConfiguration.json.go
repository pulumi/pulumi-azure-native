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
					DataDisks: batch.DataDiskArray{
						&batch.DataDiskArgs{
							Caching:            batch.CachingTypeReadWrite,
							DiskSizeGB:         pulumi.Int(30),
							Lun:                pulumi.Int(0),
							StorageAccountType: batch.StorageAccountType_Premium_LRS,
						},
						&batch.DataDiskArgs{
							Caching:            batch.CachingTypeNone,
							DiskSizeGB:         pulumi.Int(200),
							Lun:                pulumi.Int(1),
							StorageAccountType: batch.StorageAccountType_Standard_LRS,
						},
					},
					DiskEncryptionConfiguration: &batch.DiskEncryptionConfigurationArgs{
						Targets: batch.DiskEncryptionTargetArray{
							batch.DiskEncryptionTargetOsDisk,
							batch.DiskEncryptionTargetTemporaryDisk,
						},
					},
					ImageReference: &batch.ImageReferenceArgs{
						Offer:     pulumi.String("WindowsServer"),
						Publisher: pulumi.String("MicrosoftWindowsServer"),
						Sku:       pulumi.String("2016-Datacenter-SmallDisk"),
						Version:   pulumi.String("latest"),
					},
					LicenseType:    pulumi.String("Windows_Server"),
					NodeAgentSkuId: pulumi.String("batch.node.windows amd64"),
					NodePlacementConfiguration: &batch.NodePlacementConfigurationArgs{
						Policy: batch.NodePlacementPolicyTypeZonal,
					},
					OsDisk: &batch.OSDiskArgs{
						EphemeralOSDiskSettings: &batch.DiffDiskSettingsArgs{
							Placement: batch.DiffDiskPlacementCacheDisk,
						},
					},
					WindowsConfiguration: &batch.WindowsConfigurationArgs{
						EnableAutomaticUpdates: pulumi.Bool(false),
					},
				},
			},
			NetworkConfiguration: &batch.NetworkConfigurationArgs{
				EndpointConfiguration: &batch.PoolEndpointConfigurationArgs{
					InboundNatPools: batch.InboundNatPoolArray{
						&batch.InboundNatPoolArgs{
							BackendPort:            pulumi.Int(12001),
							FrontendPortRangeEnd:   pulumi.Int(15100),
							FrontendPortRangeStart: pulumi.Int(15000),
							Name:                   pulumi.String("testnat"),
							NetworkSecurityGroupRules: batch.NetworkSecurityGroupRuleArray{
								&batch.NetworkSecurityGroupRuleArgs{
									Access:              batch.NetworkSecurityGroupRuleAccessAllow,
									Priority:            pulumi.Int(150),
									SourceAddressPrefix: pulumi.String("192.100.12.45"),
									SourcePortRanges: pulumi.StringArray{
										pulumi.String("1"),
										pulumi.String("2"),
									},
								},
								&batch.NetworkSecurityGroupRuleArgs{
									Access:              batch.NetworkSecurityGroupRuleAccessDeny,
									Priority:            pulumi.Int(3500),
									SourceAddressPrefix: pulumi.String("*"),
									SourcePortRanges: pulumi.StringArray{
										pulumi.String("*"),
									},
								},
							},
							Protocol: batch.InboundEndpointProtocolTCP,
						},
					},
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
