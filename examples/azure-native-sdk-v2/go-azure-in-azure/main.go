package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/managedidentity/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const innerProgram = "yaml-simple"

// This program creates a VM that's open to SSH connections. It then copies a separate, "inner"
// Pulumi program to the VM and runs `pulumi up` and `down` on it. The intent is to prove that
// managed identity authentication, which needs to happen on an Azure resource, works as expected.
// The VM parts are mostly based on
// https://learn.microsoft.com/en-us/azure/virtual-machines/linux/quick-create-template.
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		clientConf, err := authorization.GetClientConfig(ctx)
		if err != nil {
			return err
		}

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
			// The actual IP address is not known until after the resource is created, so we ignore
			// changes to the property and look the IP up later via LookupPublicIPAddress.
		}, pulumi.IgnoreChanges([]string{"ipAddress"}))
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

		// clientId is required for user-assigned identity to disambiguate between several identities.
		var clientId pulumi.StringOutput = pulumi.String("").ToStringOutput()
		vmIdentity := &compute.VirtualMachineIdentityArgs{Type: compute.ResourceIdentityTypeSystemAssigned}
		var umi *managedidentity.UserAssignedIdentity
		if os.Getenv("PULUMI_TEST_USER_IDENTITY") == "true" {
			fmt.Printf("go-azure-in-azure: using user-assigned identity\n")

			umi, err = managedidentity.NewUserAssignedIdentity(ctx, "umi", &managedidentity.UserAssignedIdentityArgs{
				ResourceGroupName: rg.Name,
			})
			if err != nil {
				return err
			}
			clientId = umi.ClientId

			// Create a second user-assigned identity to test multiple identities. With multiple identities, the one to
			// use needs to be specified via clientId.
			umi2, err := managedidentity.NewUserAssignedIdentity(ctx, "umi2", &managedidentity.UserAssignedIdentityArgs{
				ResourceGroupName: rg.Name,
			})
			if err != nil {
				return err
			}

			vmIdentity = &compute.VirtualMachineIdentityArgs{
				Type:                   compute.ResourceIdentityTypeUserAssigned,
				UserAssignedIdentities: pulumi.StringArray{umi.ID(), umi2.ID()},
			}
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
			VmName:   pulumi.String("myVM"),
			Identity: vmIdentity,
		})
		if err != nil {
			return err
		}

		var principalId pulumi.StringOutput
		if umi != nil {
			principalId = umi.PrincipalId
		} else {
			principalId = vm.Identity.Elem().PrincipalId()
		}

		// Grant the new VM identity access to the resource group
		roleAssignment, err := authorization.NewRoleAssignment(ctx, "rgAccess",
			&authorization.RoleAssignmentArgs{
				PrincipalId:      principalId,
				PrincipalType:    authorization.PrincipalTypeServicePrincipal,
				RoleDefinitionId: pulumi.String("/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c"), // Contributor
				Scope:            rg.ID(),
			}, pulumi.DeleteBeforeReplace(true))
		if err != nil {
			return err
		}

		ipLookup := vm.ID().ApplyT(func(_ pulumi.ID) network.LookupPublicIPAddressResultOutput {
			return network.LookupPublicIPAddressOutput(ctx, network.LookupPublicIPAddressOutputArgs{
				ResourceGroupName:   rg.Name,
				PublicIpAddressName: publicIp.Name,
			})
		}).(network.LookupPublicIPAddressResultOutput)

		// Poll the server until it responds. Because other commands depend on this command, they
		// are guaranteed to hit an already booted server.
		poll, err := remote.NewCommand(ctx, "poll", &remote.CommandArgs{
			Connection: remote.ConnectionArgs{
				Host:           ipLookup.IpAddress().Elem(),
				User:           pulumi.String("pulumi"),
				PrivateKey:     privateKey.PrivateKeyOpenssh,
				DialErrorLimit: pulumi.IntPtr(-1),
			},
			Create: pulumi.String("echo 'Connection established'"),
		}, pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "10m"}), pulumi.DependsOn([]pulumi.Resource{publicIp, vm}))
		if err != nil {
			return err
		}

		sshConn := remote.ConnectionArgs{
			Host:       ipLookup.IpAddress().Elem(),
			User:       pulumi.String("pulumi"),
			PrivateKey: privateKey.PrivateKeyOpenssh,
		}

		copy, err := remote.NewCopyToRemote(ctx, "copy", &remote.CopyToRemoteArgs{
			Connection: sshConn,
			Source:     pulumi.NewFileArchive(innerProgram + "/"),
			RemotePath: pulumi.String(innerProgram),
			Triggers:   pulumi.ToArray([]any{vm.ID()}),
		}, pulumi.DependsOn([]pulumi.Resource{poll}))
		if err != nil {
			return err
		}

		installPulumi, err := remote.NewCommand(ctx, "installPulumi", &remote.CommandArgs{
			Connection: sshConn,
			Create:     pulumi.String("curl -fsSL https://get.pulumi.com | sh"),
			Triggers:   pulumi.ToArray([]any{vm.ID()}),
		}, pulumi.DependsOn([]pulumi.Resource{poll}))
		if err != nil {
			return err
		}

		// Copy the provider binary under test (the one on PATH) to the VM.
		// We put it in the same directory as the pulumi binary which is on the PATH.
		// Note that this can only work if the VM has the same architecture as the local machine.
		providerBinaryPath, err := exec.LookPath("pulumi-resource-azure-native")
		if err != nil {
			return err
		}
		copyProvider, err := remote.NewCopyToRemote(ctx, "copyProvider", &remote.CopyToRemoteArgs{
			Connection: sshConn,
			Source:     pulumi.NewFileAsset(providerBinaryPath),
			RemotePath: pulumi.String("/home/pulumi/.pulumi/bin/"),
			Triggers:   pulumi.ToArray([]any{vm.ID()}),
		}, pulumi.DependsOn([]pulumi.Resource{poll, installPulumi}))
		if err != nil {
			return err
		}
		chmodProvider, err := remote.NewCommand(ctx, "chmodProvider", &remote.CommandArgs{
			Connection: sshConn,
			Create:     pulumi.String("chmod +x /home/pulumi/.pulumi/bin/pulumi-resource-azure-native"),
			Triggers:   pulumi.ToArray([]any{vm.ID()}),
		}, pulumi.DependsOn([]pulumi.Resource{copyProvider}))
		if err != nil {
			return err
		}

		// Pass feature flags into the VM.
		useAutorest := os.Getenv("PULUMI_USE_AUTOREST")
		useLegacyAuth := os.Getenv("PULUMI_USE_LEGACY_AUTH")

		var tenantId pulumi.StringOutput = pulumi.String(os.Getenv("ARM_TENANT_ID")).ToStringOutput()

		// We pass the resource group's ID into the inner program via config so the program can
		// create a resource in the resource group.
		create := pulumi.Sprintf(`cd %s && \
set -euxo pipefail && \
export ARM_USE_MSI=true && \
export ARM_SUBSCRIPTION_ID=%s && \
export PATH="$HOME/.pulumi/bin:$PATH" && \
export PULUMI_CONFIG_PASSPHRASE=pass && \
export PULUMI_USE_AUTOREST=%s && \
export PULUMI_USE_LEGACY_AUTH=%s && \
rand=$(openssl rand -hex 4) && \
stackname="%s-$rand" && \
pulumi login --local && \
pulumi stack init $stackname && \
pulumi config set azure-native:clientId "%s" -s $stackname && \
pulumi config set azure-native:tenantId "%s" -s $stackname && \
pulumi config set objectId "%s" -s $stackname && \
pulumi config set rgId "%s" -s $stackname && \
pulumi config -s $stackname && \
pulumi up -s $stackname --skip-preview --logtostderr --logflow -v=9 && \
pulumi down -s $stackname --skip-preview --logtostderr --logflow -v=9 && \
pulumi stack rm --yes $stackname && \
pulumi logout --local`, innerProgram, clientConf.SubscriptionId, useAutorest, useLegacyAuth, innerProgram, clientId, tenantId, principalId, rg.ID())

		pulumiPreview, err := remote.NewCommand(ctx, "pulumiUpDown", &remote.CommandArgs{
			Connection: sshConn,
			Triggers:   pulumi.ToArray([]any{vm.ID(), principalId, roleAssignment.ID()}),
			Create:     create,
		}, pulumi.DependsOn([]pulumi.Resource{roleAssignment, copy, copyProvider, chmodProvider, installPulumi}))
		if err != nil {
			return err
		}

		ctx.Export("rg", rg.ID())
		ctx.Export("vm", vm.Name)
		ctx.Export("principal", principalId)
		ctx.Export("publicIpAddress", ipLookup.IpAddress().Elem())
		ctx.Export("installPulumi", installPulumi.Stdout)
		ctx.Export("installPulumiStderr", installPulumi.Stderr)
		ctx.Export("providerBinary", copyProvider.Source)
		ctx.Export("pulumiStdout", pulumiPreview.Stdout)
		ctx.Export("pulumiStderr", pulumiPreview.Stderr)

		return nil
	})
}
