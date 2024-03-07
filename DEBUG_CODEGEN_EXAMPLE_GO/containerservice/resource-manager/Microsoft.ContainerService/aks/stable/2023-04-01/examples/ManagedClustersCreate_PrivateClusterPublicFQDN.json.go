package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewManagedCluster(ctx, "managedCluster", &containerservice.ManagedClusterArgs{
			AddonProfiles: nil,
			AgentPoolProfiles: containerservice.ManagedClusterAgentPoolProfileArray{
				&containerservice.ManagedClusterAgentPoolProfileArgs{
					Count:                  pulumi.Int(3),
					EnableEncryptionAtHost: pulumi.Bool(true),
					EnableNodePublicIP:     pulumi.Bool(true),
					Mode:                   pulumi.String("System"),
					Name:                   pulumi.String("nodepool1"),
					OsType:                 pulumi.String("Linux"),
					Type:                   pulumi.String("VirtualMachineScaleSets"),
					VmSize:                 pulumi.String("Standard_DS2_v2"),
				},
			},
			ApiServerAccessProfile: &containerservice.ManagedClusterAPIServerAccessProfileArgs{
				EnablePrivateCluster:           pulumi.Bool(true),
				EnablePrivateClusterPublicFQDN: pulumi.Bool(true),
			},
			AutoScalerProfile: &containerservice.ManagedClusterPropertiesAutoScalerProfileArgs{
				ScaleDownDelayAfterAdd: pulumi.String("15m"),
				ScanInterval:           pulumi.String("20s"),
			},
			DnsPrefix:               pulumi.String("dnsprefix1"),
			EnablePodSecurityPolicy: pulumi.Bool(true),
			EnableRBAC:              pulumi.Bool(true),
			KubernetesVersion:       pulumi.String(""),
			LinuxProfile: &containerservice.ContainerServiceLinuxProfileArgs{
				AdminUsername: pulumi.String("azureuser"),
				Ssh: &containerservice.ContainerServiceSshConfigurationArgs{
					PublicKeys: containerservice.ContainerServiceSshPublicKeyArray{
						&containerservice.ContainerServiceSshPublicKeyArgs{
							KeyData: pulumi.String("keydata"),
						},
					},
				},
			},
			Location: pulumi.String("location1"),
			NetworkProfile: &containerservice.ContainerServiceNetworkProfileArgs{
				LoadBalancerProfile: &containerservice.ManagedClusterLoadBalancerProfileArgs{
					ManagedOutboundIPs: &containerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs{
						Count: pulumi.Int(2),
					},
				},
				LoadBalancerSku: pulumi.String("standard"),
				OutboundType:    pulumi.String("loadBalancer"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ResourceName:      pulumi.String("clustername1"),
			ServicePrincipalProfile: &containerservice.ManagedClusterServicePrincipalProfileArgs{
				ClientId: pulumi.String("clientid"),
				Secret:   pulumi.String("secret"),
			},
			Sku: &containerservice.ManagedClusterSKUArgs{
				Name: pulumi.String("Basic"),
				Tier: pulumi.String("Free"),
			},
			Tags: pulumi.StringMap{
				"archv2": pulumi.String(""),
				"tier":   pulumi.String("production"),
			},
			WindowsProfile: &containerservice.ManagedClusterWindowsProfileArgs{
				AdminPassword: pulumi.String("replacePassword1234$"),
				AdminUsername: pulumi.String("azureuser"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
