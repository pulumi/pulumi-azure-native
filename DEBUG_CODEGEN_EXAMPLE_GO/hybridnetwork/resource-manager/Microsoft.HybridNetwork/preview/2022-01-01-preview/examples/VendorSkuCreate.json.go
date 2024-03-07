package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewVendorSkus(ctx, "vendorSkus", &hybridnetwork.VendorSkusArgs{
			DeploymentMode:             pulumi.String("PrivateEdgeZone"),
			ManagedApplicationTemplate: nil,
			NetworkFunctionTemplate: &hybridnetwork.NetworkFunctionTemplateArgs{
				NetworkFunctionRoleConfigurations: hybridnetwork.NetworkFunctionRoleConfigurationArray{
					&hybridnetwork.NetworkFunctionRoleConfigurationArgs{
						CustomProfile: &hybridnetwork.CustomProfileArgs{
							MetadataConfigurationPath: pulumi.String("/var/logs/network.cfg"),
						},
						NetworkInterfaces: hybridnetwork.NetworkInterfaceArray{
							&hybridnetwork.NetworkInterfaceArgs{
								IpConfigurations: hybridnetwork.NetworkInterfaceIPConfigurationArray{
									&hybridnetwork.NetworkInterfaceIPConfigurationArgs{
										Gateway:            pulumi.String(""),
										IpAddress:          pulumi.String(""),
										IpAllocationMethod: pulumi.String("Dynamic"),
										IpVersion:          pulumi.String("IPv4"),
										Subnet:             pulumi.String(""),
									},
								},
								MacAddress:           pulumi.String(""),
								NetworkInterfaceName: pulumi.String("nic1"),
								VmSwitchType:         pulumi.String("Wan"),
							},
							&hybridnetwork.NetworkInterfaceArgs{
								IpConfigurations: hybridnetwork.NetworkInterfaceIPConfigurationArray{
									&hybridnetwork.NetworkInterfaceIPConfigurationArgs{
										Gateway:            pulumi.String(""),
										IpAddress:          pulumi.String(""),
										IpAllocationMethod: pulumi.String("Dynamic"),
										IpVersion:          pulumi.String("IPv4"),
										Subnet:             pulumi.String(""),
									},
								},
								MacAddress:           pulumi.String(""),
								NetworkInterfaceName: pulumi.String("nic2"),
								VmSwitchType:         pulumi.String("Management"),
							},
						},
						OsProfile: &hybridnetwork.OsProfileArgs{
							AdminUsername: pulumi.String("dummyuser"),
							CustomData:    pulumi.String("base-64 encoded string of custom data"),
							LinuxConfiguration: &hybridnetwork.LinuxConfigurationArgs{
								Ssh: &hybridnetwork.SshConfigurationArgs{
									PublicKeys: hybridnetwork.SshPublicKeyArray{
										&hybridnetwork.SshPublicKeyArgs{
											KeyData: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAgEAwrr66r8n6B8Y0zMF3dOpXEapIQD9DiYQ6D6/zwor9o39jSkHNiMMER/GETBbzP83LOcekm02aRjo55ArO7gPPVvCXbrirJu9pkm4AC4BBre5xSLS= user@constoso-DSH"),
											Path:    pulumi.String("home/user/.ssh/authorized_keys"),
										},
									},
								},
							},
						},
						RoleName: pulumi.String("test"),
						RoleType: pulumi.String("VirtualMachine"),
						StorageProfile: &hybridnetwork.StorageProfileArgs{
							DataDisks: hybridnetwork.DataDiskArray{
								&hybridnetwork.DataDiskArgs{
									CreateOption: pulumi.String("Empty"),
									DiskSizeGB:   pulumi.Int(10),
									Name:         pulumi.String("DataDisk1"),
								},
							},
							ImageReference: &hybridnetwork.ImageReferenceArgs{
								Offer:     pulumi.String("UbuntuServer"),
								Publisher: pulumi.String("Canonical"),
								Sku:       pulumi.String("18.04-LTS"),
								Version:   pulumi.String("18.04.201804262"),
							},
							OsDisk: &hybridnetwork.OsDiskArgs{
								DiskSizeGB: pulumi.Int(30),
								Name:       pulumi.String("vhdName"),
								OsType:     pulumi.String("Linux"),
								Vhd: &hybridnetwork.VirtualHardDiskArgs{
									Uri: pulumi.String("https://contoso.net/link/vnd.vhd?sp=rl&st=2020-10-08T20:38:19Z&se=2020-12-09T19:38:00Z&sv=2019-12-12&sr=b&sig=7BM2f4yOw%3D"),
								},
							},
						},
						VirtualMachineSize: pulumi.String("Standard_D3_v2"),
					},
				},
			},
			NetworkFunctionType: pulumi.String("VirtualNetworkFunction"),
			Preview:             pulumi.Bool(true),
			SkuName:             pulumi.String("TestSku"),
			VendorName:          pulumi.String("TestVendor"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
