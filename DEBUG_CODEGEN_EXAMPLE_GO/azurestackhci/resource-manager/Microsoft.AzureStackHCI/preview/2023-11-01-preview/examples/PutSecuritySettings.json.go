package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewSecuritySetting(ctx, "securitySetting", &azurestackhci.SecuritySettingArgs{
			ClusterName:                     pulumi.String("myCluster"),
			ResourceGroupName:               pulumi.String("test-rg"),
			SecuredCoreComplianceAssignment: pulumi.String("Audit"),
			SecuritySettingsName:            pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
