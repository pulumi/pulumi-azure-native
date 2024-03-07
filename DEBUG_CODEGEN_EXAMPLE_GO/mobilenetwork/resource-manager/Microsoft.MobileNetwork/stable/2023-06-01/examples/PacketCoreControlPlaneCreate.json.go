package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewPacketCoreControlPlane(ctx, "packetCoreControlPlane", &mobilenetwork.PacketCoreControlPlaneArgs{
			ControlPlaneAccessInterface: &mobilenetwork.InterfacePropertiesArgs{
				Name: pulumi.String("N2"),
			},
			CoreNetworkTechnology: pulumi.String("5GC"),
			Installation: &mobilenetwork.InstallationArgs{
				DesiredState: pulumi.String("Installed"),
			},
			LocalDiagnosticsAccess: &mobilenetwork.LocalDiagnosticsAccessConfigurationArgs{
				AuthenticationType: pulumi.String("AAD"),
				HttpsServerCertificate: &mobilenetwork.HttpsServerCertificateArgs{
					CertificateUrl: pulumi.String("https://contosovault.vault.azure.net/certificates/ingress"),
				},
			},
			Location:                   pulumi.String("eastus"),
			PacketCoreControlPlaneName: pulumi.String("TestPacketCoreCP"),
			Platform: &mobilenetwork.PlatformConfigurationArgs{
				AzureStackEdgeDevice: &mobilenetwork.AzureStackEdgeDeviceResourceIdArgs{
					Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestAzureStackEdgeDevice"),
				},
				ConnectedCluster: &mobilenetwork.ConnectedClusterResourceIdArgs{
					Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/TestConnectedCluster"),
				},
				CustomLocation: &mobilenetwork.CustomLocationResourceIdArgs{
					Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/TestCustomLocation"),
				},
				Type: pulumi.String("AKS-HCI"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			Sites: mobilenetwork.SiteResourceIdArray{
				&mobilenetwork.SiteResourceIdArgs{
					Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite"),
				},
			},
			Sku:     pulumi.String("G0"),
			UeMtu:   pulumi.Int(1600),
			Version: pulumi.String("0.2.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
