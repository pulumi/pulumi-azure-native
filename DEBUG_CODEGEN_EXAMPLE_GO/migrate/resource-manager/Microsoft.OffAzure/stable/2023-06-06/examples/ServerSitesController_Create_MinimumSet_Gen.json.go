package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewServerSitesController(ctx, "serverSitesController", &offazure.ServerSitesControllerArgs{
			Location:          pulumi.String("iipybgplhzhpcfv"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("74-35-"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
