package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewHypervHostController(ctx, "hypervHostController", &offazure.HypervHostControllerArgs{
			Fqdn:              pulumi.String("okkwk"),
			HostName:          pulumi.String("XQ-753h-765DG2"),
			ProvisioningState: pulumi.String("Created"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			RunAsAccountId:    pulumi.String("pjgyebznvfxnjzjqt"),
			SiteName:          pulumi.String("---7ed-q11Nx46-98"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
