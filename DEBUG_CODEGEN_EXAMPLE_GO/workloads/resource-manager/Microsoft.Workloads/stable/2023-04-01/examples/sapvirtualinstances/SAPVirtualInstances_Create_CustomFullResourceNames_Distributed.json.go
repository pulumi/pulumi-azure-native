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
					AppResourceGroup: "X00-RG",
					ApplicationServer: workloads.ApplicationServerConfiguration{
						InstanceCount: 6,
						SubnetId:      "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet",
						VirtualMachineConfiguration: workloads.VirtualMachineConfiguration{
							ImageReference: workloads.ImageReference{
								Offer:     "RHEL-SAP",
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
					CentralServer: workloads.CentralServerConfiguration{
						InstanceCount: 1,
						SubnetId:      "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet",
						VirtualMachineConfiguration: workloads.VirtualMachineConfiguration{
							ImageReference: workloads.ImageReference{
								Offer:     "RHEL-SAP",
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
							VmSize: "Standard_E16ds_v4",
						},
					},
					CustomResourceNames: workloads.ThreeTierFullResourceNames{
						ApplicationServer: workloads.ApplicationServerFullResourceNames{
							AvailabilitySetName: "appAvSet",
							VirtualMachines: []workloads.VirtualMachineResourceNames{
								{
									DataDiskNames: {
										"default": []string{
											"app0disk0",
										},
									},
									HostName: "apphostName0",
									NetworkInterfaces: []workloads.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: "appnic0",
										},
									},
									OsDiskName: "app0osdisk",
									VmName:     "appvm0",
								},
								{
									DataDiskNames: {
										"default": []string{
											"app1disk0",
										},
									},
									HostName: "apphostName1",
									NetworkInterfaces: []workloads.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: "appnic1",
										},
									},
									OsDiskName: "app1osdisk",
									VmName:     "appvm1",
								},
							},
						},
						CentralServer: workloads.CentralServerFullResourceNames{
							VirtualMachines: []workloads.VirtualMachineResourceNames{
								{
									DataDiskNames: {
										"default": []string{
											"ascsdisk0",
										},
									},
									HostName: "ascshostName",
									NetworkInterfaces: []workloads.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: "ascsnic",
										},
									},
									OsDiskName: "ascsosdisk",
									VmName:     "ascsvm",
								},
							},
						},
						DatabaseServer: workloads.DatabaseServerFullResourceNames{
							VirtualMachines: []workloads.VirtualMachineResourceNames{
								{
									DataDiskNames: {
										"hanaData": []string{
											"hanadata0",
											"hanadata1",
										},
										"hanaLog": []string{
											"hanalog0",
											"hanalog1",
											"hanalog2",
										},
										"hanaShared": []string{
											"hanashared0",
											"hanashared1",
										},
										"usrSap": []string{
											"usrsap0",
										},
									},
									HostName: "dbhostName",
									NetworkInterfaces: []workloads.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: "dbnic",
										},
									},
									OsDiskName: "dbosdisk",
									VmName:     "dbvm",
								},
							},
						},
						NamingPatternType: "FullResourceName",
						SharedStorage: workloads.SharedStorageResourceNames{
							SharedStorageAccountName:                "storageacc",
							SharedStorageAccountPrivateEndPointName: "peForxNFS",
						},
					},
					DatabaseServer: workloads.DatabaseConfiguration{
						DatabaseType:  "HANA",
						InstanceCount: 1,
						SubnetId:      "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet",
						VirtualMachineConfiguration: workloads.VirtualMachineConfiguration{
							ImageReference: workloads.ImageReference{
								Offer:     "RHEL-SAP",
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
							VmSize: "Standard_M32ts",
						},
					},
					DeploymentType: "ThreeTier",
				},
				OsSapConfiguration: workloads.OsSapConfiguration{
					SapFqdn: "xyz.test.com",
				},
			},
			Environment:            pulumi.String("Prod"),
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
