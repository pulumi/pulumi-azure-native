package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabric.NewNodeType(ctx, "nodeType", &servicefabric.NodeTypeArgs{
			Capacities: pulumi.StringMap{
				"ClientConnections": pulumi.String("65536"),
			},
			ClusterName:             pulumi.String("myCluster"),
			DataDiskSizeGB:          pulumi.Int(200),
			DataDiskType:            pulumi.String("Premium_LRS"),
			IsPrimary:               pulumi.Bool(false),
			IsStateless:             pulumi.Bool(true),
			MultiplePlacementGroups: pulumi.Bool(true),
			NodeTypeName:            pulumi.String("BE"),
			PlacementProperties: pulumi.StringMap{
				"HasSSD":       pulumi.String("true"),
				"NodeColor":    pulumi.String("green"),
				"SomeProperty": pulumi.String("5"),
			},
			ResourceGroupName: pulumi.String("resRg"),
			VmExtensions: []servicefabric.VMSSExtensionArgs{
				{
					AutoUpgradeMinorVersion: pulumi.Bool(true),
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
			VmInstanceCount:  -1,
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
			VmSize: pulumi.String("Standard_DS3"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
