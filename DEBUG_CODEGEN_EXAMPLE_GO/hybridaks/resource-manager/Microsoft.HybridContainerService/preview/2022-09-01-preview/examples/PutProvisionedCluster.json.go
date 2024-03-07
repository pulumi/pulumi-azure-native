package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcontainerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcontainerservice.NewProvisionedCluster(ctx, "provisionedCluster", &hybridcontainerservice.ProvisionedClusterArgs{
			ExtendedLocation: &hybridcontainerservice.ProvisionedClustersExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
				Type: pulumi.String("CustomLocation"),
			},
			Location: pulumi.String("westus"),
			Properties: &hybridcontainerservice.ProvisionedClustersAllPropertiesArgs{
				AgentPoolProfiles: hybridcontainerservice.NamedAgentPoolProfileArray{
					&hybridcontainerservice.NamedAgentPoolProfileArgs{
						Count:  pulumi.Int(1),
						Name:   pulumi.String("default-nodepool-1"),
						OsType: pulumi.String("Linux"),
						VmSize: pulumi.String("Standard_A4_v2"),
					},
				},
				CloudProviderProfile: &hybridcontainerservice.CloudProviderProfileArgs{
					InfraNetworkProfile: &hybridcontainerservice.CloudProviderProfileInfraNetworkProfileArgs{
						VnetSubnetIds: pulumi.StringArray{
							pulumi.String("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/virtualNetworks/test-vnet-static"),
						},
					},
					InfraStorageProfile: &hybridcontainerservice.CloudProviderProfileInfraStorageProfileArgs{
						StorageSpaceIds: pulumi.StringArray{
							pulumi.String("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/storageSpaces/test-storage"),
						},
					},
				},
				ControlPlane: &hybridcontainerservice.ControlPlaneProfileArgs{
					Count: pulumi.Int(1),
					LinuxProfile: &hybridcontainerservice.LinuxProfilePropertiesArgs{
						Ssh: &hybridcontainerservice.LinuxProfilePropertiesSshArgs{
							PublicKeys: hybridcontainerservice.LinuxProfilePropertiesPublicKeysArray{
								&hybridcontainerservice.LinuxProfilePropertiesPublicKeysArgs{
									KeyData: pulumi.String("ssh-rsa AAAAB1NzaC1yc2EAAAADAQABAAACAQCY......"),
								},
							},
						},
					},
					OsType: pulumi.String("Linux"),
					VmSize: pulumi.String("Standard_A4_v2"),
				},
				KubernetesVersion: pulumi.String("v1.20.5"),
				LinuxProfile: &hybridcontainerservice.LinuxProfilePropertiesArgs{
					Ssh: &hybridcontainerservice.LinuxProfilePropertiesSshArgs{
						PublicKeys: hybridcontainerservice.LinuxProfilePropertiesPublicKeysArray{
							&hybridcontainerservice.LinuxProfilePropertiesPublicKeysArgs{
								KeyData: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCY......."),
							},
						},
					},
				},
				NetworkProfile: &hybridcontainerservice.NetworkProfileArgs{
					LoadBalancerProfile: &hybridcontainerservice.LoadBalancerProfileArgs{
						Count: pulumi.Int(1),
						LinuxProfile: &hybridcontainerservice.LinuxProfilePropertiesArgs{
							Ssh: &hybridcontainerservice.LinuxProfilePropertiesSshArgs{
								PublicKeys: hybridcontainerservice.LinuxProfilePropertiesPublicKeysArray{
									&hybridcontainerservice.LinuxProfilePropertiesPublicKeysArgs{
										KeyData: pulumi.String("ssh-rsa AAAAB2NzaC1yc2EAAAADAQABAAACAQCY......"),
									},
								},
							},
						},
						OsType: pulumi.String("Linux"),
						VmSize: pulumi.String("Standard_K8S3_v1"),
					},
					LoadBalancerSku: pulumi.String("unstacked-haproxy"),
					NetworkPolicy:   pulumi.String("calico"),
					PodCidr:         pulumi.String("10.244.0.0/16"),
				},
			},
			ResourceGroupName: pulumi.String("test-arcappliance-resgrp"),
			ResourceName:      pulumi.String("test-hybridakscluster"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
