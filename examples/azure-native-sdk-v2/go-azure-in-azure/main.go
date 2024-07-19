package main

import (
	"fmt"

	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The VM parts are mostly based on https://learn.microsoft.com/en-us/azure/virtual-machines/linux/quick-create-template
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

		vm, err := compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
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

		sshConn := remote.ConnectionArgs{
			Host:       publicIp.IpAddress.Elem(),
			User:       pulumi.String("pulumi"),
			PrivateKey: privateKey.PrivateKeyOpenssh,
		}

		// Poll the server until it responds. Because other commands depend on this command, they
		// are guaranteed to hit an already booted server.
		poll, err := remote.NewCommand(ctx, "poll", &remote.CommandArgs{
			Connection: remote.ConnectionArgs{
				Host:           publicIp.IpAddress.Elem(),
				User:           pulumi.String("pulumi"),
				PrivateKey:     privateKey.PrivateKeyOpenssh,
				DialErrorLimit: pulumi.IntPtr(-1),
			},
			Create: pulumi.String("echo 'Connection established'"),
		}, pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "10m"}), pulumi.DependsOn([]pulumi.Resource{vm}))
		if err != nil {
			return err
		}

		const innerProgram = "yaml-simple"

		copy, err := remote.NewCopyToRemote(ctx, "copy", &remote.CopyToRemoteArgs{
			Connection: sshConn,
			Source:     pulumi.NewFileArchive(innerProgram + "/"),
			RemotePath: pulumi.String(innerProgram),
			Triggers:   pulumi.ToArray([]any{publicIp.IpAddress}),
		}, pulumi.DependsOn([]pulumi.Resource{poll}))
		if err != nil {
			return err
		}

		installPulumi, err := remote.NewCommand(ctx, "installPulumi", &remote.CommandArgs{
			Connection: sshConn,
			Create:     pulumi.String("curl -fsSL https://get.pulumi.com | sh"),
			Triggers:   pulumi.ToArray([]any{publicIp.IpAddress}),
		}, pulumi.DependsOn([]pulumi.Resource{poll}))
		if err != nil {
			return err
		}

		check, err := remote.NewCommand(ctx, "check", &remote.CommandArgs{
			Connection: sshConn,
			Create:     pulumi.String(fmt.Sprintf("cat %s/Pulumi.yaml && ls .pulumi/bin/", innerProgram)),
		}, pulumi.DependsOn([]pulumi.Resource{copy, installPulumi}))
		if err != nil {
			return err
		}

		pulumiPreview, err := remote.NewCommand(ctx, "pulumiPreview", &remote.CommandArgs{
			Connection: sshConn,
			// export ARM_USE_MSI=true && \
			Create: pulumi.String(fmt.Sprintf(`cd %s && \
export PATH=~/.pulumi/bin:$PATH && \
export PULUMI_CONFIG_PASSPHRASE=pass && \
rand=$(openssl rand -hex 4) && \
stackname="%s-$rand" && \
pulumi login --local && \
pulumi stack init $stackname && \
pulumi preview && \
pulumi stack rm --yes $stackname && \
pulumi logout --local`, innerProgram, innerProgram)),
		}, pulumi.DependsOn([]pulumi.Resource{copy, installPulumi}))
		if err != nil {
			return err
		}

		ctx.Export("publicIpAddress", publicIp.IpAddress)
		ctx.Export("check", check.Stdout)
		ctx.Export("installPulumi", installPulumi.Stdout)
		ctx.Export("installPulumiStderr", installPulumi.Stderr)
		ctx.Export("pulumiPreview", pulumiPreview.Stdout)
		ctx.Export("pulumiPreviewStderr", pulumiPreview.Stderr)

		return nil
	})
}
