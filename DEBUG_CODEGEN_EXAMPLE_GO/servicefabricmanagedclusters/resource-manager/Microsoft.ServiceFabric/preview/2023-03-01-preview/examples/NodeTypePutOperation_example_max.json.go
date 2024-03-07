package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabric.NewNodeType(ctx, "nodeType", &servicefabric.NodeTypeArgs{
			AdditionalDataDisks: []servicefabric.VmssDataDiskArgs{
				{
					DiskLetter: pulumi.String("F"),
					DiskSizeGB: pulumi.Int(256),
					DiskType:   pulumi.String("StandardSSD_LRS"),
					Lun:        pulumi.Int(1),
				},
				{
					DiskLetter: pulumi.String("G"),
					DiskSizeGB: pulumi.Int(150),
					DiskType:   pulumi.String("Premium_LRS"),
					Lun:        pulumi.Int(2),
				},
			},
			Capacities: pulumi.StringMap{
				"ClientConnections": pulumi.String("65536"),
			},
			ClusterName:                 pulumi.String("myCluster"),
			DataDiskLetter:              pulumi.String("S"),
			DataDiskSizeGB:              pulumi.Int(200),
			DataDiskType:                pulumi.String("Premium_LRS"),
			EnableAcceleratedNetworking: pulumi.Bool(true),
			EnableEncryptionAtHost:      pulumi.Bool(true),
			EnableNodePublicIP:          pulumi.Bool(true),
			EnableOverProvisioning:      pulumi.Bool(false),
			EvictionPolicy:              pulumi.String("Deallocate"),
			FrontendConfigurations: []servicefabric.FrontendConfigurationArgs{
				{
					ApplicationGatewayBackendAddressPoolId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/applicationGateways/appgw-test/backendAddressPools/appgwBepoolTest"),
					LoadBalancerBackendAddressPoolId:       pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/backendAddressPools/LoadBalancerBEAddressPool"),
					LoadBalancerInboundNatPoolId:           pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/inboundNatPools/LoadBalancerNATPool"),
				},
			},
			IsPrimary:               pulumi.Bool(false),
			IsSpotVM:                pulumi.Bool(true),
			IsStateless:             pulumi.Bool(true),
			MultiplePlacementGroups: pulumi.Bool(true),
			NatGatewayId:            pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/natGateways/myNatGateway"),
			NodeTypeName:            pulumi.String("BE"),
			PlacementProperties: pulumi.StringMap{
				"HasSSD":       pulumi.String("true"),
				"NodeColor":    pulumi.String("green"),
				"SomeProperty": pulumi.String("5"),
			},
			ResourceGroupName:            pulumi.String("resRg"),
			SecureBootEnabled:            pulumi.Bool(true),
			SecurityType:                 pulumi.String("TrustedLaunch"),
			SpotRestoreTimeout:           pulumi.String("PT30M"),
			SubnetId:                     pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
			UseDefaultPublicLoadBalancer: pulumi.Bool(true),
			UseEphemeralOSDisk:           pulumi.Bool(true),
			VmExtensions: []servicefabric.VMSSExtensionArgs{
				{
					AutoUpgradeMinorVersion: pulumi.Bool(true),
					EnableAutomaticUpgrade:  pulumi.Bool(true),
					ForceUpdateTag:          pulumi.String("v.1.0"),
					Name:                    pulumi.String("Microsoft.Azure.Geneva.GenevaMonitoring"),
					Publisher:               pulumi.String("Microsoft.Azure.Geneva"),
					Settings:                nil,
					Type:                    pulumi.String("GenevaMonitoring"),
					TypeHandlerVersion:      pulumi.String("2.0"),
				},
			},
			VmImageOffer:     pulumi.String("WindowsServer"),
			VmImagePublisher: pulumi.String("MicrosoftWindowsServer"),
			VmImageSku:       pulumi.String("2016-Datacenter-Server-Core"),
			VmImageVersion:   pulumi.String("latest"),
			VmInstanceCount:  pulumi.Int(10),
			VmManagedIdentity: &servicefabric.VmManagedIdentityArgs{
				UserAssignedIdentities: pulumi.StringArray{
					pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity"),
					pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity2"),
				},
			},
			VmSecrets: []servicefabric.VaultSecretGroupArgs{
				{
					SourceVault: {
						Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.KeyVault/vaults/myVault"),
					},
					VaultCertificates: servicefabric.VaultCertificateArray{
						{
							CertificateStore: pulumi.String("My"),
							CertificateUrl:   pulumi.String("https://myVault.vault.azure.net:443/secrets/myCert/ef1a31d39e1f46bca33def54b6cda54c"),
						},
					},
				},
			},
			VmSetupActions: pulumi.StringArray{
				pulumi.String("EnableContainers"),
				pulumi.String("EnableHyperV"),
			},
			VmSize: pulumi.String("Standard_DS3"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
