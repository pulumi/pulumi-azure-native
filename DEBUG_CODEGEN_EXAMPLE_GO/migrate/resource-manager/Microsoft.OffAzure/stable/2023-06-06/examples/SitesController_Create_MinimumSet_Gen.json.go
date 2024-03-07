package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewSitesController(ctx, "sitesController", &offazure.SitesControllerArgs{
			Location:          pulumi.String("mh"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("KPV1Y68-G0V"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
