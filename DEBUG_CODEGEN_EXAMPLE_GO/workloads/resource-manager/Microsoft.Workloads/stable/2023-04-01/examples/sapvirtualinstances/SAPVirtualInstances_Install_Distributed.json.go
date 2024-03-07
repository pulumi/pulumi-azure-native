package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewSAPVirtualInstance(ctx, "sapVirtualInstance", &workloads.SAPVirtualInstanceArgs{
			Configuration: workloads.DeploymentWithOSConfiguration{
				AppLocation:       "eastus",
				ConfigurationType: "DeploymentWithOSConfig",
				InfrastructureConfiguration: workloads.ThreeTierConfiguration{
					AppResourceGroup: "{{resourcegrp}}",
					ApplicationServer: workloads.ApplicationServerConfiguration{
						InstanceCount: 2,
						SubnetId:      "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app",
						VirtualMachineConfiguration: workloads.VirtualMachineConfiguration{
							ImageReference: workloads.ImageReference{
								Offer:     "RHEL-SAP-HA",
								Publisher: "RedHat",
								Sku:       "8.2",
								Version:   "8.2.2021091201",
							},
							OsProfile: workloads.OSProfile{
								AdminUsername: "azureuser",
								OsConfiguration: workloads.LinuxConfiguration{
									DisablePasswordAuthentication: true,
									OsType:                        "Linux",
									SshKeyPair: workloads.SshKeyPair{
										PrivateKey: "{{privateKey}}",
										PublicKey:  "{{sshkey}}",
									},
								},
							},
							VmSize: "Standard_E4ds_v4",
						},
					},
					CentralServer: workloads.CentralServerConfiguration{
						InstanceCount: 1,
						SubnetId:      "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app",
						VirtualMachineConfiguration: workloads.VirtualMachineConfiguration{
							ImageReference: workloads.ImageReference{
								Offer:     "RHEL-SAP-HA",
								Publisher: "RedHat",
								Sku:       "8.2",
								Version:   "8.2.2021091201",
							},
							OsProfile: workloads.OSProfile{
								AdminUsername: "azureuser",
								OsConfiguration: workloads.LinuxConfiguration{
									DisablePasswordAuthentication: true,
									OsType:                        "Linux",
									SshKeyPair: workloads.SshKeyPair{
										PrivateKey: "{{privateKey}}",
										PublicKey:  "{{sshkey}}",
									},
								},
							},
							VmSize: "Standard_E4ds_v4",
						},
					},
					DatabaseServer: workloads.DatabaseConfiguration{
						InstanceCount: 1,
						SubnetId:      "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app",
						VirtualMachineConfiguration: workloads.VirtualMachineConfiguration{
							ImageReference: workloads.ImageReference{
								Offer:     "RHEL-SAP-HA",
								Publisher: "RedHat",
								Sku:       "8.2",
								Version:   "8.2.2021091201",
							},
							OsProfile: workloads.OSProfile{
								AdminUsername: "azureuser",
								OsConfiguration: workloads.LinuxConfiguration{
									DisablePasswordAuthentication: true,
									OsType:                        "Linux",
									SshKeyPair: workloads.SshKeyPair{
										PrivateKey: "{{privateKey}}",
										PublicKey:  "{{sshkey}}",
									},
								},
							},
							VmSize: "Standard_M32ts",
						},
					},
					DeploymentType: "ThreeTier",
					NetworkConfiguration: workloads.NetworkConfiguration{
						IsSecondaryIpEnabled: true,
					},
				},
				OsSapConfiguration: workloads.OsSapConfiguration{
					SapFqdn: "sap.bpaas.com",
				},
				SoftwareConfiguration: workloads.SAPInstallWithoutOSConfigSoftwareConfiguration{
					BomUrl:                   "https://teststorageaccount.blob.core.windows.net/sapbits/sapfiles/boms/S41909SPS03_v0011ms/S41909SPS03_v0011ms.yaml",
					SapBitsStorageAccountId:  "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/teststorageaccount",
					SoftwareInstallationType: "SAPInstallWithoutOSConfig",
					SoftwareVersion:          "SAP S/4HANA 1909 SPS 03",
				},
			},
			Environment:            pulumi.String("Prod"),
			Location:               pulumi.String("eastus2"),
			ResourceGroupName:      pulumi.String("test-rg"),
			SapProduct:             pulumi.String("S4HANA"),
			SapVirtualInstanceName: pulumi.String("X00"),
			Tags: pulumi.StringMap{
				"created by": pulumi.String("azureuser"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
