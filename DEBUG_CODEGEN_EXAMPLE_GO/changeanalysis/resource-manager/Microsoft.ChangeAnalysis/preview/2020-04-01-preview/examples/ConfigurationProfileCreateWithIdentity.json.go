package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/changeanalysis/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := changeanalysis.NewConfigurationProfile(ctx, "configurationProfile", &changeanalysis.ConfigurationProfileArgs{
			ProfileName: pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
