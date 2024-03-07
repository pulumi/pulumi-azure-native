package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewExtension(ctx, "extension", &azurestackhci.ExtensionArgs{
			ArcSettingName:         pulumi.String("default"),
			ClusterName:            pulumi.String("myCluster"),
			EnableAutomaticUpgrade: pulumi.Bool(false),
			ExtensionName:          pulumi.String("MicrosoftMonitoringAgent"),
			ProtectedSettings: pulumi.Any{
				WorkspaceKey: "xx",
			},
			Publisher:         pulumi.String("Microsoft.Compute"),
			ResourceGroupName: pulumi.String("test-rg"),
			Settings: pulumi.Any{
				WorkspaceId: "xx",
			},
			Type:               pulumi.String("MicrosoftMonitoringAgent"),
			TypeHandlerVersion: pulumi.String("1.10"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
