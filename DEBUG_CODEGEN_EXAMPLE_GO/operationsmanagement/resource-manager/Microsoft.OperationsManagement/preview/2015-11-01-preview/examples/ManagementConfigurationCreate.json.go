package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationsmanagement.NewManagementConfiguration(ctx, "managementConfiguration", &operationsmanagement.ManagementConfigurationArgs{
			Location:                    pulumi.String("East US"),
			ManagementConfigurationName: pulumi.String("managementConfiguration1"),
			ResourceGroupName:           pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
