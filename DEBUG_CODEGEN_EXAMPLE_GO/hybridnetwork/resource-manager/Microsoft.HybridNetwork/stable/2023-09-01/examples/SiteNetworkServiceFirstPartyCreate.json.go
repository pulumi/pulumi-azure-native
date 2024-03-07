package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewSiteNetworkService(ctx, "siteNetworkService", &hybridnetwork.SiteNetworkServiceArgs{
			Location: pulumi.String("westUs2"),
			Properties: &hybridnetwork.SiteNetworkServicePropertiesFormatArgs{
				DesiredStateConfigurationGroupValueReferences: hybridnetwork.ReferencedResourceMap{
					"MyVM_Configuration": &hybridnetwork.ReferencedResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/configurationgroupvalues/MyVM_Configuration1"),
					},
				},
				NetworkServiceDesignVersionResourceReference: hybridnetwork.SecretDeploymentResourceReference{
					Id:     "/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/TestPublisher/networkServiceDesignGroups/TestNetworkServiceDesignGroupName/networkServiceDesignVersions/1.0.0",
					IdType: "Secret",
				},
				SiteReference: &hybridnetwork.ReferencedResourceArgs{
					Id: pulumi.String("/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/sites/testSite"),
				},
			},
			ResourceGroupName:      pulumi.String("rg1"),
			SiteNetworkServiceName: pulumi.String("testSiteNetworkServiceName"),
			Sku: &hybridnetwork.SkuArgs{
				Name: pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
