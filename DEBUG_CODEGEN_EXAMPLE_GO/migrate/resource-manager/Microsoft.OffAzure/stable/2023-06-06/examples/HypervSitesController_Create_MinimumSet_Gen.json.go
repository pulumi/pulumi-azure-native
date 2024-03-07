package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewHypervSitesController(ctx, "hypervSitesController", &offazure.HypervSitesControllerArgs{
			Location:          pulumi.String("sctymxdndonxgejdhi"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("Y-C-7---V49GV-058XE-6P5"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
