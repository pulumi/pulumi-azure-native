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
				InfrastructureConfiguration: workloads.SingleServerConfiguration{
					AppResourceGroup: "test-rg",
					DeploymentType:   "SingleServer",
					SubnetId:         "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/testsubnet",
					VirtualMachineConfiguration: workloads.VirtualMachineConfiguration{
						ImageReference: workloads.ImageReference{
							Offer:     "SLES-SAP",
							Publisher: "SUSE",
							Sku:       "12-sp4-gen2",
							Version:   "2022.02.01",
						},
						OsProfile: workloads.OSProfile{
							AdminUsername: "azureappadmin",
							OsConfiguration: workloads.LinuxConfiguration{
								DisablePasswordAuthentication: true,
								OsType:                        "Linux",
								SshKeyPair: workloads.SshKeyPair{
									PrivateKey: "{{privateKey}}",
									PublicKey:  "{{sshkey}}",
								},
							},
						},
						VmSize: "Standard_E32ds_v4",
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
			Environment:            pulumi.String("NonProd"),
			Location:               pulumi.String("eastus2"),
			ResourceGroupName:      pulumi.String("test-rg"),
			SapProduct:             pulumi.String("S4HANA"),
			SapVirtualInstanceName: pulumi.String("X00"),
			Tags:                   nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
