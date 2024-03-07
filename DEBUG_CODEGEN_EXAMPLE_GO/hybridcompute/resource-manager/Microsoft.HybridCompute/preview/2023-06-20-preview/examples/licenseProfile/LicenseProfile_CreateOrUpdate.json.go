package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcompute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcompute.NewLicenseProfile(ctx, "licenseProfile", &hybridcompute.LicenseProfileArgs{
			AssignedLicense:    pulumi.String("{LicenseResourceId}"),
			LicenseProfileName: pulumi.String("default"),
			Location:           pulumi.String("eastus2euap"),
			MachineName:        pulumi.String("myMachine"),
			ResourceGroupName:  pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
