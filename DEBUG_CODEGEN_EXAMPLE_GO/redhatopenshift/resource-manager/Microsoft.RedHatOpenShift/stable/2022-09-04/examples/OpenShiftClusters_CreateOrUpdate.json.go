package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/redhatopenshift/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redhatopenshift.NewOpenShiftCluster(ctx, "openShiftCluster", &redhatopenshift.OpenShiftClusterArgs{
			ApiserverProfile: &redhatopenshift.APIServerProfileArgs{
				Visibility: pulumi.String("Public"),
			},
			ClusterProfile: &redhatopenshift.ClusterProfileArgs{
				Domain:               pulumi.String("cluster.location.aroapp.io"),
				FipsValidatedModules: pulumi.String("Enabled"),
				PullSecret:           pulumi.String("{\"auths\":{\"registry.connect.redhat.com\":{\"auth\":\"\"},\"registry.redhat.io\":{\"auth\":\"\"}}}"),
				ResourceGroupId:      pulumi.String("/subscriptions/subscriptionId/resourceGroups/clusterResourceGroup"),
			},
			ConsoleProfile: nil,
			IngressProfiles: []redhatopenshift.IngressProfileArgs{
				{
					Name:       pulumi.String("default"),
					Visibility: pulumi.String("Public"),
				},
			},
			Location: pulumi.String("location"),
			MasterProfile: &redhatopenshift.MasterProfileArgs{
				EncryptionAtHost: pulumi.String("Enabled"),
				SubnetId:         pulumi.String("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/master"),
				VmSize:           pulumi.String("Standard_D8s_v3"),
			},
			NetworkProfile: &redhatopenshift.NetworkProfileArgs{
				PodCidr:     pulumi.String("10.128.0.0/14"),
				ServiceCidr: pulumi.String("172.30.0.0/16"),
			},
			ResourceGroupName: pulumi.String("resourceGroup"),
			ResourceName:      pulumi.String("resourceName"),
			ServicePrincipalProfile: &redhatopenshift.ServicePrincipalProfileArgs{
				ClientId:     pulumi.String("clientId"),
				ClientSecret: pulumi.String("clientSecret"),
			},
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
			WorkerProfiles: []redhatopenshift.WorkerProfileArgs{
				{
					Count:      pulumi.Int(3),
					DiskSizeGB: pulumi.Int(128),
					Name:       pulumi.String("worker"),
					SubnetId:   pulumi.String("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
					VmSize:     pulumi.String("Standard_D2s_v3"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
