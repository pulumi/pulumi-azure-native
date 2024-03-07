package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcompute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcompute.NewLicense(ctx, "license", &hybridcompute.LicenseArgs{
			LicenseDetails: &hybridcompute.LicenseDetailsArgs{
				Edition:    pulumi.String("Datacenter"),
				Processors: pulumi.Int(6),
				State:      pulumi.String("Activated"),
				Target:     pulumi.String("Windows Server 2012"),
				Type:       pulumi.String("pCore"),
			},
			LicenseName:       pulumi.String("{licenseName}"),
			LicenseType:       pulumi.String("ESU"),
			Location:          pulumi.String("eastus2euap"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
