package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkcloud.NewAgentPool(ctx, "agentPool", &networkcloud.AgentPoolArgs{
			AdministratorConfiguration: &networkcloud.AdministratorConfigurationArgs{
				AdminUsername: pulumi.String("azure"),
				SshPublicKeys: networkcloud.SshPublicKeyArray{
					&networkcloud.SshPublicKeyArgs{
						KeyData: pulumi.String("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
					},
				},
			},
			AgentOptions: &networkcloud.AgentOptionsArgs{
				HugepagesCount: pulumi.Float64(96),
				HugepagesSize:  pulumi.String("1G"),
			},
			AgentPoolName: pulumi.String("agentPoolName"),
			AttachedNetworkConfiguration: &networkcloud.AttachedNetworkConfigurationArgs{
				L2Networks: networkcloud.L2NetworkAttachmentConfigurationArray{
					&networkcloud.L2NetworkAttachmentConfigurationArgs{
						NetworkId:  pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName"),
						PluginType: pulumi.String("DPDK"),
					},
				},
				L3Networks: networkcloud.L3NetworkAttachmentConfigurationArray{
					&networkcloud.L3NetworkAttachmentConfigurationArgs{
						IpamEnabled: pulumi.String("False"),
						NetworkId:   pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
						PluginType:  pulumi.String("SRIOV"),
					},
				},
				TrunkedNetworks: networkcloud.TrunkedNetworkAttachmentConfigurationArray{
					&networkcloud.TrunkedNetworkAttachmentConfigurationArgs{
						NetworkId:  pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName"),
						PluginType: pulumi.String("MACVLAN"),
					},
				},
			},
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("1"),
				pulumi.String("2"),
				pulumi.String("3"),
			},
			Count: pulumi.Float64(3),
			ExtendedLocation: &networkcloud.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
				Type: pulumi.String("CustomLocation"),
			},
			KubernetesClusterName: pulumi.String("kubernetesClusterName"),
			Labels: networkcloud.KubernetesLabelArray{
				&networkcloud.KubernetesLabelArgs{
					Key:   pulumi.String("kubernetes.label"),
					Value: pulumi.String("true"),
				},
			},
			Location:          pulumi.String("location"),
			Mode:              pulumi.String("System"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("myvalue1"),
				"key2": pulumi.String("myvalue2"),
			},
			Taints: networkcloud.KubernetesLabelArray{
				&networkcloud.KubernetesLabelArgs{
					Key:   pulumi.String("kubernetes.taint"),
					Value: pulumi.String("true"),
				},
			},
			UpgradeSettings: &networkcloud.AgentPoolUpgradeSettingsArgs{
				MaxSurge: pulumi.String("1"),
			},
			VmSkuName: pulumi.String("NC_M16_v1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
