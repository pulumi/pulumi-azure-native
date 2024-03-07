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
					AppResourceGroup: "X00-RG",
					DatabaseType:     "HANA",
					DeploymentType:   "SingleServer",
					NetworkConfiguration: workloads.NetworkConfiguration{
						IsSecondaryIpEnabled: true,
					},
					SubnetId: "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet",
					VirtualMachineConfiguration: workloads.VirtualMachineConfiguration{
						ImageReference: workloads.ImageReference{
							Offer:     "RHEL-SAP-HA",
							Publisher: "RedHat",
							Sku:       "84sapha-gen2",
							Version:   "latest",
						},
						OsProfile: workloads.OSProfile{
							AdminUsername: "{your-username}",
							OsConfiguration: workloads.LinuxConfiguration{
								DisablePasswordAuthentication: true,
								OsType:                        "Linux",
								SshKeyPair: workloads.SshKeyPair{
									PrivateKey: "xyz",
									PublicKey:  "abc",
								},
							},
						},
						VmSize: "Standard_E32ds_v4",
					},
				},
				OsSapConfiguration: workloads.OsSapConfiguration{
					SapFqdn: "xyz.test.com",
				},
				SoftwareConfiguration: workloads.ExternalInstallationSoftwareConfiguration{
					CentralServerVmId:        "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/sapq20scsvm0",
					SoftwareInstallationType: "External",
				},
			},
			Environment:            pulumi.String("NonProd"),
			Location:               pulumi.String("westcentralus"),
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
