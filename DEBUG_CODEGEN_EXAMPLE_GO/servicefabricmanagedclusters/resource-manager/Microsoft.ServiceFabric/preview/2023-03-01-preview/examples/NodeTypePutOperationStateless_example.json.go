package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabric.NewNodeType(ctx, "nodeType", &servicefabric.NodeTypeArgs{
			ClusterName:             pulumi.String("myCluster"),
			EnableEncryptionAtHost:  pulumi.Bool(true),
			IsPrimary:               pulumi.Bool(false),
			IsStateless:             pulumi.Bool(true),
			MultiplePlacementGroups: pulumi.Bool(true),
			NodeTypeName:            pulumi.String("BE"),
			ResourceGroupName:       pulumi.String("resRg"),
			UseTempDataDisk:         pulumi.Bool(true),
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
			VmInstanceCount:  pulumi.Int(10),
			VmSize:           pulumi.String("Standard_DS3"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
