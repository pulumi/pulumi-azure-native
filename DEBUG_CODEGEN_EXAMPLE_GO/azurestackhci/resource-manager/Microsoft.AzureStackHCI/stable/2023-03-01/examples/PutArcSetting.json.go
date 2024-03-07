package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewArcSetting(ctx, "arcSetting", &azurestackhci.ArcSettingArgs{
			ArcSettingName:    pulumi.String("default"),
			ClusterName:       pulumi.String("myCluster"),
			ResourceGroupName: pulumi.String("test-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
