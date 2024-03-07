package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewSite(ctx, "site", &hybridnetwork.SiteArgs{
			Location: pulumi.String("westUs2"),
			Properties: hybridnetwork.SitePropertiesFormatResponse{
				Nfvis: pulumi.Array{
					hybridnetwork.AzureCoreNFVIDetails{
						Location: "westUs2",
						Name:     "nfvi1",
						NfviType: "AzureCore",
					},
					hybridnetwork.AzureArcK8sClusterNFVIDetails{
						CustomLocationReference: hybridnetwork.ReferencedResource{
							Id: "/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation1",
						},
						Name:     "nfvi2",
						NfviType: "AzureArcKubernetes",
					},
					hybridnetwork.AzureOperatorNexusClusterNFVIDetails{
						CustomLocationReference: hybridnetwork.ReferencedResource{
							Id: "/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation2",
						},
						Name:     "nfvi3",
						NfviType: "AzureOperatorNexus",
					},
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			SiteName:          pulumi.String("testSite"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
