package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewManagedCluster(ctx, "managedCluster", &containerservice.ManagedClusterArgs{
			AgentPoolProfiles: containerservice.ManagedClusterAgentPoolProfileArray{
				&containerservice.ManagedClusterAgentPoolProfileArgs{
					Count:              pulumi.Int(3),
					EnableNodePublicIP: pulumi.Bool(true),
					Mode:               pulumi.String("System"),
					Name:               pulumi.String("nodepool1"),
					OsType:             pulumi.String("Linux"),
					Type:               pulumi.String("VirtualMachineScaleSets"),
					VmSize:             pulumi.String("Standard_DS2_v2"),
				},
			},
			DnsPrefix:         pulumi.String("dnsprefix1"),
			KubernetesVersion: pulumi.String(""),
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
			SecurityProfile: &containerservice.ManagedClusterSecurityProfileArgs{
				Defender: &containerservice.ManagedClusterSecurityProfileDefenderArgs{
					LogAnalyticsWorkspaceResourceId: pulumi.String("/subscriptions/SUB_ID/resourcegroups/RG_NAME/providers/microsoft.operationalinsights/workspaces/WORKSPACE_NAME"),
					SecurityMonitoring: &containerservice.ManagedClusterSecurityProfileDefenderSecurityMonitoringArgs{
						Enabled: pulumi.Bool(true),
					},
				},
				WorkloadIdentity: &containerservice.ManagedClusterSecurityProfileWorkloadIdentityArgs{
					Enabled: pulumi.Bool(true),
				},
			},
			Sku: &containerservice.ManagedClusterSKUArgs{
				Name: pulumi.String("Basic"),
				Tier: pulumi.String("Free"),
			},
			Tags: pulumi.StringMap{
				"archv2": pulumi.String(""),
				"tier":   pulumi.String("production"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
