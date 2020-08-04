package main

import (
	"encoding/base64"

	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
	insights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/insights/v20150401"
	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20200501"
	resources "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/resources/v20200601"
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		c := config.New(ctx, "")

		var adminUser, adminPassword, domain pulumi.StringOutput

		user := c.Get("adminUser")
		if user == "" {
			adminUser = pulumi.String("azureuser").ToStringOutput()
		} else {
			adminUser = pulumi.String(user).ToStringOutput()
		}

		password := c.Get("adminPassword")
		if password == "" {
			pwd, err := random.NewRandomPassword(ctx, "pwd", &random.RandomPasswordArgs{
				Length:  pulumi.Int(20),
				Special: pulumi.Bool(true),
			})
			if err != nil {
				return err
			}

			adminPassword = pwd.Result.ToStringOutput()
		} else {
			adminPassword = pulumi.String(password).ToStringOutput()
		}

		d := c.Get("domain")
		if d == "" {
			randomDomain, err := random.NewRandomString(ctx, "domain", &random.RandomStringArgs{
				Length:  pulumi.Int(10),
				Number:  pulumi.Bool(false),
				Special: pulumi.Bool(false),
				Upper:   pulumi.Bool(false),
			})
			if err != nil {
				return err
			}

			domain = randomDomain.Result.ToStringOutput()
		} else {
			domain = pulumi.String(d).ToStringOutput()
		}

		applicationPort := pulumi.Int(80).ToIntOutput()
		if port := c.GetInt("applicationPort"); port != 0 {
			applicationPort = pulumi.Int(port).ToIntOutput()
		}

		location := pulumi.String("westus2").ToStringOutput()
		if loc := c.Get("location"); loc != "" {
			location = pulumi.String(loc).ToStringOutput()
		}

		// Create an Azure Resource Group
		resourceGroup, err := resources.NewResourceGroup(ctx, "azurerm-govmset", &resources.ResourceGroupArgs{
			Name:     pulumi.String("azurerm-govmset"),
			Location: location,
			Tags: pulumi.StringMap{
				"Owner": pulumi.String("azurerm-test"),
			},
		})
		if err != nil {
			return err
		}

		publicIP, err := network.NewPublicIPAddress(ctx, "public-ip", &network.PublicIPAddressArgs{
			Name:              pulumi.String("publicip"),
			Location:          location, // Required by service
			ResourceGroupName: resourceGroup.Name,
			Sku: network.PublicIPAddressSkuPtr(&network.PublicIPAddressSkuArgs{
				Name: pulumi.String("Standard"),
			}).ToPublicIPAddressSkuPtrOutput(),
			PublicIPAllocationMethod: pulumi.String("Static"),
			DnsSettings: network.PublicIPAddressDnsSettingsPtrInput(&network.PublicIPAddressDnsSettingsArgs{
				DomainNameLabel: domain,
			}),
		})
		if err != nil {
			return err
		}

		vnet, err := network.NewVirtualNetwork(ctx, "vnet", &network.VirtualNetworkArgs{
			ResourceGroupName: resourceGroup.Name,
			Name:              pulumi.String("vnet"),
			Location:          location, // Required by service
			AddressSpace: network.AddressSpacePtr(&network.AddressSpaceArgs{
				AddressPrefixes: pulumi.StringArray{pulumi.String("10.0.0.0/16")},
			}).ToAddressSpacePtrOutput(),
		})
		if err != nil {
			return err
		}

		subnet, err := network.NewSubnet(ctx, "subnet", &network.SubnetArgs{
			ResourceGroupName:  resourceGroup.Name,
			Name:               pulumi.String("subnet"),
			AddressPrefix:      pulumi.String("10.0.2.0/24"),
			VirtualNetworkName: vnet.Name,
		})
		if err != nil {
			return err
		}

		loadBalancer, err := network.NewLoadBalancer(ctx, "lb", &network.LoadBalancerArgs{
			ResourceGroupName: resourceGroup.Name,
			Name:              pulumi.String("lb"),
			Location:          location, // Required by service
			Sku: network.LoadBalancerSkuPtr(&network.LoadBalancerSkuArgs{
				Name: pulumi.String("Standard"),
			}).ToLoadBalancerSkuPtrOutput(),
			FrontendIPConfigurations: network.FrontendIPConfigurationArray{
				network.FrontendIPConfigurationArgs{
					Name: pulumi.String("PublicIPAddress"),
					PublicIPAddress: network.PublicIPAddressTypePtr(&network.PublicIPAddressTypeArgs{
						Id: publicIP.ID(),
					}).ToPublicIPAddressTypePtrOutput(),
				},
			},

			// Error: azurerm:network:LoadBalancer (lb):
			//      error: Code="InvalidRequestFormat" Message="Cannot parse the request." Details=[]
			//Probes: network.ProbeArray{
			//	network.ProbeArgs{
			//		Protocol: pulumi.String("Tcp"),
			//		Port:     applicationPort,
			//	},
			//},
		}, pulumi.DependsOn([]pulumi.Resource{publicIP}))
		if err != nil {
			return err
		}

		bpepool, err := network.NewLoadBalancerBackendAddressPool(ctx, "bpepool", &network.LoadBalancerBackendAddressPoolArgs{
			ResourceGroupName: resourceGroup.Name,
			Name:              pulumi.String("bpepool"),
			LoadBalancerName:  loadBalancer.Name,
		}, pulumi.DependsOn([]pulumi.Resource{loadBalancer}))
		if err != nil {
			return err
		}

		if _, err = network.NewInboundNatRule(ctx, "natrule", &network.InboundNatRuleArgs{
			BackendPort:       applicationPort,
			FrontendPort:      applicationPort,
			Protocol:          pulumi.StringPtr("Tcp"),
			ResourceGroupName: resourceGroup.Name,
			LoadBalancerName:  loadBalancer.Name,
			Name:              pulumi.String("allow-80"),
			FrontendIPConfiguration: network.SubResourcePtr(&network.SubResourceArgs{
				Id: loadBalancer.Properties.FrontendIPConfigurations().ToFrontendIPConfigurationResponseArrayOutput().Index(pulumi.Int(0)).Id(),
			}).ToSubResourcePtrOutput(),
		}); err != nil {
			return err
		}

		customData := base64.StdEncoding.EncodeToString([]byte(`#cloud-config
		packages:
			- nginx`))

		scaleSet, err := compute.NewVirtualMachineScaleSet(ctx, "vmscaleset", &compute.VirtualMachineScaleSetArgs{
			ResourceGroupName: resourceGroup.Name,
			Name:              pulumi.String("vmscaleset"),
			Location:          location,
			VirtualMachineProfile: compute.VirtualMachineScaleSetVMProfilePtr(&compute.VirtualMachineScaleSetVMProfileArgs{
				NetworkProfile: compute.VirtualMachineScaleSetNetworkProfilePtr(&compute.VirtualMachineScaleSetNetworkProfileArgs{
					NetworkInterfaceConfigurations: compute.VirtualMachineScaleSetNetworkConfigurationArray{
						compute.VirtualMachineScaleSetNetworkConfigurationArgs{
							Name: pulumi.String("NICConfiguration"),
							IpConfigurations: compute.VirtualMachineScaleSetIPConfigurationArray{
								compute.VirtualMachineScaleSetIPConfigurationArgs{
									Name:                            pulumi.String("VMSetIPConfiguration"),
									LoadBalancerBackendAddressPools: compute.SubResourceArray{compute.SubResourceArgs{Id: bpepool.ID().ToStringOutput()}},
									Primary:                         pulumi.Bool(true),
									Subnet: compute.ApiEntityReferencePtr(&compute.ApiEntityReferenceArgs{
										Id: subnet.ID(),
									}),
								},
							},
							Primary: pulumi.Bool(true),
						},
					},
				}),
				OsProfile: compute.VirtualMachineScaleSetOSProfilePtr(
					&compute.VirtualMachineScaleSetOSProfileArgs{
						AdminUsername:      adminUser,
						AdminPassword:      adminPassword,
						ComputerNamePrefix: pulumi.String("vmlab"),
						CustomData:         pulumi.String(customData),
						LinuxConfiguration: compute.LinuxConfigurationPtr(&compute.LinuxConfigurationArgs{
							DisablePasswordAuthentication: pulumi.Bool(false),
						}),
					}),
				DiagnosticsProfile: compute.DiagnosticsProfilePtr(&compute.DiagnosticsProfileArgs{
					BootDiagnostics: compute.BootDiagnosticsPtr(&compute.BootDiagnosticsArgs{
						Enabled: pulumi.Bool(true),
					}),
				}),
				StorageProfile: compute.VirtualMachineScaleSetStorageProfilePtr(&compute.VirtualMachineScaleSetStorageProfileArgs{
					DataDisks: compute.VirtualMachineScaleSetDataDiskArray{
						compute.VirtualMachineScaleSetDataDiskArgs{
							Caching:      pulumi.StringPtr("ReadWrite"),
							CreateOption: pulumi.String("Empty"),
							DiskSizeGB:   pulumi.IntPtr(10),
							Lun:          pulumi.Int(0),
						},
					},
					ImageReference: compute.ImageReferencePtr(&compute.ImageReferenceArgs{
						Offer:     pulumi.String("UbuntuServer"),
						Publisher: pulumi.String("Canonical"),
						Sku:       pulumi.StringPtr("16.04-LTS"),
						Version:   pulumi.StringPtr("latest"),
					}),
					OsDisk: compute.VirtualMachineScaleSetOSDiskPtr(&compute.VirtualMachineScaleSetOSDiskArgs{
						Caching:      pulumi.StringPtr("ReadWrite"),
						CreateOption: pulumi.String("FromImage"),
						ManagedDisk: compute.VirtualMachineScaleSetManagedDiskParametersPtr(
							&compute.VirtualMachineScaleSetManagedDiskParametersArgs{
								StorageAccountType: pulumi.StringPtr("Standard_LRS"),
							}),
					}),
				}),
			}).ToVirtualMachineScaleSetVMProfilePtrOutput(),
			UpgradePolicy: compute.UpgradePolicyPtr(&compute.UpgradePolicyArgs{
				Mode: pulumi.StringPtr("Manual"),
			}).ToUpgradePolicyPtrOutput(),
			Sku: compute.SkuPtr(&compute.SkuArgs{
				Capacity: pulumi.Int(1),
				Name:     pulumi.StringPtr("Standard_DS1_v2"),
			}).ToSkuPtrOutput(),
		}, pulumi.DependsOn([]pulumi.Resource{subnet}))
		if err != nil {
			return err
		}

		_, err = insights.NewAutoscaleSetting(ctx, "vmss-autoscale", &insights.AutoscaleSettingArgs{
			ResourceGroupName: resourceGroup.Name,
			Location:          location,
			Name:              pulumi.String("vmss-autoscale"),
			Notifications: insights.AutoscaleNotificationArray{
				insights.AutoscaleNotificationArgs{
					Operation: pulumi.String("Scale"),
					Email: insights.EmailNotificationPtr(&insights.EmailNotificationArgs{
						CustomEmails:                       pulumi.StringArray{pulumi.String("admin@contoso.com")},
						SendToSubscriptionAdministrator:    pulumi.BoolPtr(true),
						SendToSubscriptionCoAdministrators: pulumi.BoolPtr(true),
					}).ToEmailNotificationPtrOutput(),
				},
			},
			Profiles: insights.AutoscaleProfileArray{
				insights.AutoscaleProfileArgs{
					Capacity: insights.ScaleCapacityArgs{
						Default: pulumi.String("1"), // strings?
						Maximum: pulumi.String("10"),
						Minimum: pulumi.String("1"),
					},
					Name: pulumi.String("defaultProfile"),
					Rules: insights.ScaleRuleArray{
						insights.ScaleRuleArgs{
							MetricTrigger: insights.MetricTriggerArgs{
								MetricName:        pulumi.String("Percentage CPU"),
								MetricResourceUri: scaleSet.ID().ToStringOutput(),
								Operator:          pulumi.String("GreaterThan"),
								Statistic:         pulumi.String("Average"),
								Threshold:         pulumi.Float64(75),
								TimeAggregation:   pulumi.String("Average"),
								TimeGrain:         pulumi.String("PT1M"),
								TimeWindow:        pulumi.String("PT5M"),
							},
							ScaleAction: insights.ScaleActionArgs{
								Cooldown:  pulumi.String("PT1M"),
								Direction: pulumi.String("Increase"),
								Type:      pulumi.String("ChangeCount"),
								Value:     pulumi.StringPtr("1"),
							},
						},
						insights.ScaleRuleArgs{
							MetricTrigger: insights.MetricTriggerArgs{
								MetricName:        pulumi.String("Percentage CPU"),
								MetricResourceUri: scaleSet.ID(),
								Operator:          pulumi.String("LessThan"),
								Statistic:         pulumi.String("Average"),
								Threshold:         pulumi.Float64(25),
								TimeAggregation:   pulumi.String("Average"),
								TimeGrain:         pulumi.String("PT1M"),
								TimeWindow:        pulumi.String("PT5M"),
							},
							ScaleAction: insights.ScaleActionArgs{
								Cooldown:  pulumi.String("PT1M"),
								Direction: pulumi.String("Decrease"),
								Type:      pulumi.String("ChangeCount"),
								Value:     pulumi.StringPtr("1"),
							},
						},
					},
				},
			},
			TargetResourceUri: scaleSet.ID(),
		})
		if err != nil {
			return err
		}

		ctx.Export("publicAddress", publicIP.Properties.DnsSettings().Fqdn())
		return nil
	})
}
