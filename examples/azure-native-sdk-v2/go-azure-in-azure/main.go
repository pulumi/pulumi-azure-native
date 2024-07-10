package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		rg, err := resources.NewResourceGroup(ctx, "rg", &resources.ResourceGroupArgs{})
		if err != nil {
			return err
		}

		// Generate a new RSA key pair
		privateKey, err := tls.NewPrivateKey(ctx, "mySshKey", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("RSA"),
			RsaBits:   pulumi.Int(4096),
		})
		if err != nil {
			return err
		}

		vnet, err := network.NewVirtualNetwork(ctx, "vnet", &network.VirtualNetworkArgs{
			ResourceGroupName: rg.Name,
			AddressSpace: network.AddressSpaceArgs{
				AddressPrefixes: pulumi.StringArray{pulumi.String("10.1.0.0/16")},
			},
			Subnets: network.SubnetTypeArray{network.SubnetTypeArgs{
				Name:                              pulumi.String("subnet"),
				AddressPrefix:                     pulumi.String("10.1.0.0/24"),
				PrivateEndpointNetworkPolicies:    pulumi.String("Enabled"),
				PrivateLinkServiceNetworkPolicies: pulumi.String("Enabled"),
			}},
		})
		if err != nil {
			return err
		}

		publicIp, err := network.NewPublicIPAddress(ctx, "publicIp", &network.PublicIPAddressArgs{
			ResourceGroupName:        rg.Name,
			PublicIPAllocationMethod: pulumi.String("Dynamic"),
			PublicIPAddressVersion:   pulumi.String("IPv4"),
			IdleTimeoutInMinutes:     pulumi.IntPtr(5),
		})
		if err != nil {
			return err
		}

		nsg, err := network.NewNetworkSecurityGroup(ctx, "nsg", &network.NetworkSecurityGroupArgs{
			ResourceGroupName: rg.Name,
			SecurityRules: network.SecurityRuleTypeArray{
				network.SecurityRuleTypeArgs{
					Name:                     pulumi.String("SSH"),
					Priority:                 pulumi.IntPtr(1000),
					Protocol:                 pulumi.String("TCP"),
					Access:                   pulumi.String("Allow"),
					Direction:                pulumi.String("Inbound"),
					SourceAddressPrefix:      pulumi.String("*"),
					SourcePortRange:          pulumi.String("*"),
					DestinationAddressPrefix: pulumi.String("*"),
					DestinationPortRange:     pulumi.String("22"),
				},
			},
		})
		if err != nil {
			return err
		}

		nic, err := network.NewNetworkInterface(ctx, "networkInterface", &network.NetworkInterfaceArgs{
			ResourceGroupName: rg.Name,
			IpConfigurations: network.NetworkInterfaceIPConfigurationArray{
				&network.NetworkInterfaceIPConfigurationArgs{
					Name: pulumi.String("ipconfig1"),
					Subnet: &network.SubnetTypeArgs{
						Id: vnet.Subnets.Index(pulumi.Int(0)).Id(),
					},
					PrivateIPAllocationMethod: pulumi.String("Dynamic"),
					PrivateIPAddressVersion:   pulumi.String("IPv4"),
					PublicIPAddress: network.PublicIPAddressTypeArgs{
						Id: publicIp.ID(),
					},
				},
			},
			NetworkSecurityGroup: network.NetworkSecurityGroupTypeArgs{
				Id: nsg.ID(),
			},
		})
		if err != nil {
			return err
		}

		_, err = compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
			ResourceGroupName: rg.Name,
			HardwareProfile: &compute.HardwareProfileArgs{
				VmSize: pulumi.String(compute.VirtualMachineSizeTypes_Standard_A2_v2),
			},
			NetworkProfile: &compute.NetworkProfileArgs{
				NetworkInterfaces: compute.NetworkInterfaceReferenceArray{
					&compute.NetworkInterfaceReferenceArgs{
						Id:      nic.ID(),
						Primary: pulumi.Bool(true),
					},
				},
			},
			OsProfile: &compute.OSProfileArgs{
				AdminUsername: pulumi.String("pulumi"),
				ComputerName:  pulumi.String("myVM"),
				LinuxConfiguration: &compute.LinuxConfigurationArgs{
					DisablePasswordAuthentication: pulumi.Bool(true),
					Ssh: &compute.SshConfigurationArgs{
						PublicKeys: compute.SshPublicKeyTypeArray{
							&compute.SshPublicKeyTypeArgs{
								KeyData: privateKey.PublicKeyOpenssh,
								Path:    pulumi.String("/home/pulumi/.ssh/authorized_keys"),
							},
						},
					},
				},
			},
			StorageProfile: &compute.StorageProfileArgs{
				ImageReference: &compute.ImageReferenceArgs{
					Offer:     pulumi.String("ubuntu-24_04-lts"),
					Publisher: pulumi.String("canonical"),
					Sku:       pulumi.String("server-gen1"),
					Version:   pulumi.String("latest"),
				},
				OsDisk: &compute.OSDiskArgs{
					Caching:      compute.CachingTypesReadWrite,
					CreateOption: pulumi.String(compute.DiskCreateOptionTypesFromImage),
					ManagedDisk: &compute.ManagedDiskParametersArgs{
						StorageAccountType: pulumi.String(compute.StorageAccountTypes_Standard_LRS),
					},
					Name: pulumi.String("myVMosdisk"),
				},
			},
			VmName: pulumi.String("myVM"),
		})
		if err != nil {
			return err
		}

		ctx.Export("publicIpAddress", publicIp.IpAddress)

		return nil
	})
}
